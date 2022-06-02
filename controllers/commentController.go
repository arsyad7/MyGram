package controllers

import (
	"fmt"
	"mygram/params"
	"mygram/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService services.CommentService
}

func NewCommentController(srvcs *services.CommentService) *CommentController {
	return &CommentController{commentService: *srvcs}
}

func (cm *CommentController) CreateComment(c *gin.Context) {
	var req params.CreateComment

	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err.Error(),
		})
		return
	}

	id, _ := strconv.Atoi(c.Request.Header.Get("user_id"))
	res, errCreate := cm.commentService.CreateComment(&req, id)
	if errCreate != nil {
		c.JSON(errCreate.Status, errCreate)
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

func (cm *CommentController) GetComents(c *gin.Context) {
	res, errGet := cm.commentService.GetComments()
	if errGet != nil {
		c.JSON(errGet.Status, errGet)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (cm *CommentController) UpdateComment(c *gin.Context) {
	var req params.CreateComment
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err,
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("commentId"))
	res, errUpdate := cm.commentService.UpdateComment(&req, id)
	if errUpdate != nil {
		c.JSON(errUpdate.Status, errUpdate)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (cm *CommentController) DeleteComment(c *gin.Context) {
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	res := cm.commentService.DeleteComment(commentId)
	c.JSON(res.Status, params.Response{
		Message: res.Message,
	})
}
