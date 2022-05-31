package controllers

import (
	"mygram/params"
	"mygram/services"
	"net/http"

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
