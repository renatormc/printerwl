package helpers

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/renatormc/rprinter/config"
)

func DeleteOldFiles() {
	cf := config.GetConfig()
	entries, err := ioutil.ReadDir(cf.TempFolder)
	if err != nil {
		log.Println(cf.TempFolder)
	}
	for _, e := range entries {
		delta := time.Since(e.ModTime())
		if delta > time.Duration(60)*time.Second {
			os.Remove(filepath.Join(cf.TempFolder, e.Name()))
		}

	}
}
