package controllers

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/renatormc/rprinter/config"
	"github.com/renatormc/rprinter/helpers"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Print(c *gin.Context) {
	var err error
	cf := config.GetConfig()
	printer := c.Query("printer")
	if !helpers.SliceContains(cf.ServerConfig.Printers, printer) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Printer %q not found", printer),
		})
		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	p := filepath.Join(cf.TempFolder, newFileName)
	if err := c.SaveUploadedFile(file, p); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	go helpers.DeleteOldFiles()

	var out string
	if runtime.GOOS == "windows" {
		out, err = helpers.CmdExecStrOutput("PDFtoPrinter", p, printer)
	} else {
		out, err = helpers.CmdExecStrOutput("lp", "-d", printer, p)
	}

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error on print document",
		})
		return
	}
	log.Println(out)

	// File saved successfully. Return proper result
	log.Println("A file was sent to printer")
	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been sent to printer.",
	})
}
