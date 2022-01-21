package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/akamensky/argparse"
	"github.com/renatormc/rprinter/config"
	"github.com/renatormc/rprinter/helpers"
	"github.com/renatormc/rprinter/server"
)

type ServerResponse struct {
	Message string
}

func SendPostRequest(url string, filename string, printer string) string {
	fieldname := "file"
	cf := config.GetConfig()
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fieldname, filepath.Base(file.Name()))

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(part, file)
	writer.Close()
	request, err := http.NewRequest("POST", url, body)

	if err != nil {
		log.Fatal(err)
	}
	q := request.URL.Query()
	q.Add("printer", printer)
	request.URL.RawQuery = q.Encode()

	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Password", cf.ClientConfig.Password)
	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var res ServerResponse
	err = json.Unmarshal(content, &res)
	if err != nil {
		log.Print("It was not possible parse serve response")
	}

	return res.Message
}

func main() {
	parser := argparse.NewParser("Remote printer", "This app can be used to use a printer installed in a remote server")
	printer := parser.String("p", "printer", &argparse.Options{Help: "Printer name"})

	printCmd := parser.NewCommand("print", "Print a pdf document")
	filePath := printCmd.String("f", "file", &argparse.Options{Help: "File path", Required: true})
	remotePrint := printCmd.Flag("r", "remote", &argparse.Options{Help: "Print in remote server"})

	serveCmd := parser.NewCommand("serve", "Run server")

	testConn := parser.NewCommand("test-conn", "Test connection with the server")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	config.LoadConfig()
	cf := config.GetConfig()

	switch {
	case printCmd.Happened():
		p := *printer
		if p == "" {
			p = cf.ClientConfig.DefaultPrinter
		}
		if *remotePrint {
			fmt.Printf("Printing remote on printer %q\n", p)
			url := fmt.Sprintf("%s/print", cf.ClientConfig.UrlHost)
			message := SendPostRequest(url, *filePath, p)
			fmt.Println(message)
		} else {
			fmt.Printf("Printing local on printer %q\n", p)
			out, err := helpers.PrintPdf(*filePath, p)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(out)
		}

	case serveCmd.Happened():
		file, err := os.OpenFile(filepath.Join(cf.AppFolder, "log.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}

		log.SetOutput(file)
		s := server.NewServer()
		s.Run()

	case testConn.Happened():
		url := fmt.Sprintf("%s/test-conn", cf.ClientConfig.UrlHost)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Print(resp.Body)
	}
}
