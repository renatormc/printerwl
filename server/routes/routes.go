package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatormc/rprinter/config"
	"github.com/renatormc/rprinter/controllers"
)

func authRequired(c *gin.Context) {
	cf := config.GetConfig()
	pass := c.GetHeader("Password")
	fmt.Println(cf.ServerConfig.Password, pass)
	if pass != string(cf.ServerConfig.Password) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "unauthorized",
		})
		return
	}
}

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	router.Use(authRequired)
	router.GET("/test", controllers.Test)
	router.POST("/print", controllers.Print)

	return router
}
