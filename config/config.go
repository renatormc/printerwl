package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

var config Config

type Config struct {
	Port           string
	Password       []byte
	AppFolder      string
	TLSEnabled     bool
	TempFolder     string
	DefaultPrinter string
	UrlHost        string
	AcroRd32Path   string
}

func LoadConfig() {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	config.AppFolder = filepath.Dir(ex)

	err = godotenv.Load(filepath.Join(config.AppFolder, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	result, err := strconv.ParseBool(os.Getenv("TLS_ENABLED"))
	if err != nil {
		panic(err)
	}
	config.TLSEnabled = result
	config.Password = []byte(os.Getenv("PASSWORD"))
	config.Port = os.Getenv("PORT")
	config.DefaultPrinter = os.Getenv("DEFAULT_PRINTER")
	config.UrlHost = os.Getenv("URL_HOST")
	config.AcroRd32Path = os.Getenv("ACRORD32_PATH")

	config.AppFolder = filepath.ToSlash(config.AppFolder)
	config.TempFolder = path.Join(config.AppFolder, "temp")
	os.MkdirAll(config.TempFolder, os.ModePerm)
}

func GetConfig() *Config {
	return &config
}
