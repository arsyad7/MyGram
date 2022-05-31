package controllers

import (
	"mygram/middlewares"
	"mygram/params"
	"mygram/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(srvcs *services.UserService) *UserController {
	return &UserController{
		userService: *srvcs,
	}
}

func (u *UserController) Register(c *gin.Context) {
	var req params.CreateUser

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err,
		})
		return
	}

	res, errRegister := u.userService.Register(&req)
	if errRegister != nil {
		c.JSON(errRegister.Status, errRegister)
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

func (u *UserController) Login(c *gin.Context) {
	var req params.CreateUser

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err,
		})
		return
	}

	res, errLogin := u.userService.Login(&req)
	if errLogin != nil {
		c.JSON(errLogin.Status, errLogin)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (u *UserController) UpdateUser(c *gin.Context) {
	auth := middlewares.Authentication(c)
	if auth != nil {
		c.AbortWithStatusJSON(401, params.Response{
			Status:  401,
			Message: auth.Error(),
		})
		return
	}

	var req params.CreateUser
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err,
		})
		return
	}

	id, _ := strconv.Atoi(c.Request.Header.Get("user_id"))
	res, errUpdate := u.userService.UpdateUser(&req, uint(id))
	if errUpdate != nil {
		c.JSON(errUpdate.Status, errUpdate)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (u *UserController) DeleteUser(c *gin.Context) {
	auth := middlewares.Authentication(c)
	if auth != nil {
		c.AbortWithStatusJSON(401, params.Response{
			Status:  401,
			Message: auth.Error(),
		})
		return
	}

	id, _ := strconv.Atoi(c.Request.Header.Get("user_id"))
	res, errDelete := u.userService.DeleteUser(uint(id))
	if errDelete != nil {
		c.JSON(errDelete.Status, errDelete)
	} else {
		c.JSON(http.StatusOK, res)
	}
}
