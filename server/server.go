package server

import (
	"fmt"
	"log"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/renatormc/rprinter/config"
	"github.com/renatormc/rprinter/server/routes"
)

type Server struct {
	server *gin.Engine
}

func NewServer() Server {
	server := Server{
		server: gin.Default(),
	}
	return server
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	config := config.GetConfig()
	log.Printf("Server running at port: %s", config.Port)
	if config.TLSEnabled {
		log.Fatal(router.RunTLS(fmt.Sprintf(":%s", config.Port), path.Join(config.AppFolder, "local/server.crt"), path.Join(config.AppFolder, "local/server.key")))
	} else {
		log.Fatal(router.Run(fmt.Sprintf(":%s", config.Port)))
	}

}
