package main

import (
	config "main/configs"
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.Connect()

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/dogs", handlers.GetDogs)
	router.GET("/dogs/:id", handlers.GetDog)
	router.POST("/dogs", handlers.CreateDog)
	router.PUT("/dogs/:id", handlers.UpdateDog)
	router.DELETE("/dogs/:id", handlers.DeleteDog)

	router.Run(":4500")
}
