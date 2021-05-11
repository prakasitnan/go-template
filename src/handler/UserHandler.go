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

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	_id, _ := strconv.ParseInt(id, 10,64)
	user ,_ := repository.GetUserById(_id)

	if id == "" {
		c.JSON(200, model.UserResponseError{200, "Not Found", ""})
	} else {
		c.JSON(200, model.UserResponse{200, "Successfully", user})
	}
}

func CreateUser(c *gin.Context) {
	var user model.User
	bindErr := c.BindJSON(&user)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
		return
	}

	id, err := repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.UserResponseError{
			http.StatusInternalServerError,
			"Something wrong on server.",
			""})
	} else {
		user.Id = id
		c.JSON(http.StatusCreated, model.UserResponse{200, "Successfully", user})
	}

}

func UpdateUser(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.ParseInt(paramId, 10,64)

	var user model.User
	bindErr := c.BindJSON(&user)
	if bindErr != nil {
		c.JSON(http.StatusInternalServerError, bindErr.Error())
	}

	user, err := repository.UpdateUser(id, user)
	if err != nil{
		c.JSON(500, model.UserResponseError{500, "Something went wrong", ""})
	} else {
		c.JSON(200, model.UserResponse{200, "successfully", user})
	}
}

func DeleteUser(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.ParseInt(paramId, 10, 64)

	err := repository.DeleteUser(id)
	if err != nil {
		c.JSON(500, model.UserResponseError{500, "Something went wrong", ""})
	} else {
		c.JSON(200, model.UserResponseError{200, "Successfully", ""})
	}
}
