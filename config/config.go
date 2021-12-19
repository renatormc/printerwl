package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

var config Config

type ClientConfig struct {
	DefaultPrinter string `json:"default_printer"`
	UrlHost        string `json:"url_host"`
	Password       string `json:"password"`
}

type ServerConfig struct {
	ServerPort string   `json:"server_port"`
	TLSEnabled bool     `json:"tsl_enabled"`
	Printers   []string `json:"printers"`
	Password   string   `json:"password"`
}

type Config struct {
	ClientConfig ClientConfig `json:"client"`
	ServerConfig ServerConfig `json:"server"`
	TempFolder   string
	AppFolder    string
}

func LoadConfig() {

	appDir := os.Getenv("APP_FOLDER")
	_, err := os.Stat(appDir)
	if err != nil {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		config.AppFolder = filepath.Dir(ex)
	} else {
		config.AppFolder = appDir
	}

	jsonFile, err := os.Open(filepath.Join(config.AppFolder, "rprinter-settings.json"))
	if err != nil {
		log.Fatal("It was not possible to read settings file")
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("It was not possible to read settings file")
	}
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatal("It was not possible to read settings file")
	}

	config.TempFolder = path.Join(config.AppFolder, "temp")
	os.MkdirAll(config.TempFolder, os.ModePerm)

	file, err := os.OpenFile(filepath.Join(config.AppFolder, "log.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

}

func GetConfig() *Config {
	return &config
}
