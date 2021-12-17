package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/renatormc/rprinter/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	router.GET("/test", controllers.Test)
	router.POST("/print", controllers.Print)

	return router
}
