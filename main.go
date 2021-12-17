package main

import (
	"bytes"
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
	"github.com/renatormc/rprinter/server"
)

func SendPostRequest(url string, filename string) []byte {
	fieldname := "file"
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

	request.Header.Add("Content-Type", writer.FormDataContentType())
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

	return content
}

func main() {
	parser := argparse.NewParser("Remote printer", "This app can be used to use a printer installed in a remote server")
	printer := parser.String("p", "printer", &argparse.Options{Help: "Printer name", Default: "default"})

	printCmd := parser.NewCommand("print", "Print a pdf document")
	filePath := printCmd.String("f", "file", &argparse.Options{Help: "File path", Required: true})
	serveCmd := parser.NewCommand("serve", "Run server")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	config.LoadConfig()
	cf := config.GetConfig()

	switch {
	case printCmd.Happened():
		url := fmt.Sprintf("%s/print?printer=%s", cf.UrlHost, *printer)
		fmt.Println(url)
		c := SendPostRequest(url, *filePath)
		fmt.Println(string(c))
	case serveCmd.Happened():
		s := server.NewServer()
		s.Run()
	}
}
