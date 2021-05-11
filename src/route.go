package src

import (
	"github.com/gin-gonic/gin"
	"github.com/prakasitnan/go-template/src/handler"
)

func Route(router *gin.Engine) *gin.Engine {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/v1")
	{
		// USER HANDLER
		v1.GET("/users", handler.GelAllUsers)
		v1.GET("/user/:id", handler.GetUserById)
		v1.POST("/user", handler.CreateUser)
		v1.PUT("/user/:id", handler.UpdateUser)
		v1.DELETE("/user/:id", handler.DeleteUser)
	}

	return router
}