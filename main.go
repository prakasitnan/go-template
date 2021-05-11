package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/prakasitnan/go-template/src"
	"time"
)

func main() {
	fmt.Println("Hello World....")
	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins: "*",
		Methods: "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))

	src.Route(r)
	r.Run(":8080")
}
