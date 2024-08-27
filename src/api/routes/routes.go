package routes

import (
	"example.com/src/api/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func ConfigureRoutes(
	server *gin.Engine, controller controllers.IController,
) {
	log.Printf("Setting up routes")
	api := server.Group("/api/v1")
	{
		api.GET("health", controller.Health)
		api.GET("greeting", controller.Greeting)
		api.GET("bye", controller.Greeting)
	}
}
