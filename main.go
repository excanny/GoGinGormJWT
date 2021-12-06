package main

import (
//	"net/http"
	"os"
	"ginjwt2/controllers"
	//"ginjwt2/services"
	"ginjwt2/models"
	"ginjwt2/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()

	server := gin.Default()
	
	server.POST("/register", controllers.Register)
	server.POST("/login", controllers.Login)

	protected := server.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)


	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)

}