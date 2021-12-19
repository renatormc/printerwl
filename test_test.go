package main

import (
	"testing"

	"github.com/renatormc/rprinter/config"
	"github.com/renatormc/rprinter/helpers"
)

func TestMain(t *testing.T) {
	config.LoadConfig()
	helpers.DeleteOldFiles()

}
