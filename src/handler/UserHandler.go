package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prakasitnan/go-template/src/repository"
)


func  GelAllUsers(c *gin.Context) {
	users, err  := repository.GetAllUsers()
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, users)
}
