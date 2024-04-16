package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nabinkatwal7/x-backend-go/controller"
)

func ServeApplication() {
	router := gin.Default()

	router.GET("/", func (c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Server responded with status code 200",
		})
	})

	router.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)

	router.Run(":8080")
	fmt.Println("Server started on port 8080")
}