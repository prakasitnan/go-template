package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prakasitnan/go-template/src/handler"
)

func main() {
	fmt.Println("Hello World....")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/v1")
	{
		v1.GET("/users", handler.GelAllUsers)
	}


	router.Run(":8080")
}
