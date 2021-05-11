package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prakasitnan/go-template/src/model"
	"github.com/prakasitnan/go-template/src/repository"
	"net/http"
	"strconv"
)


func  GelAllUsers(c *gin.Context) {
	users, err  := repository.GetAllUsers()
	if err != nil {
		c.JSON(500, "Something went wrong")
	}

	c.JSON(200, model.UserListResponse{200, "Successfully", users})
}

func CreateUser(c *gin.Context) {
	var user model.User
	bindErr := c.BindJSON(&user)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
		return
	}

	_, err := repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprint("Something wrong on our server"))
	} else {
		c.JSON(http.StatusCreated, model.UserResponse{200, "Successfully", user})
	}

}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	_id, _ := strconv.ParseInt(id, 10,64)
	user ,_ := repository.GetUserById(_id)

	if id == "" {
		c.JSON(404, "")
	} else {
		c.JSON(200, model.UserResponse{200, "Successfully", user})
	}
}

