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

type Config struct {
	ServerPort     string   `json:"server_port"`
	Password       []byte   `json:"password"`
	TLSEnabled     bool     `json:"tsl_enabled"`
	Printers       []string `json:"printers"`
	DefaultPrinter string   `json:"default_printer"`
	UrlHost        string   `json:"url_host"`
	TempFolder     string
	AppFolder      string
}

func LoadConfig() {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	config.AppFolder = filepath.Dir(ex)

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
}

func GetConfig() *Config {
	return &config
}
