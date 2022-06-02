package controllers

import (
	"mygram/params"
	"mygram/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	photoService services.PhotoService
}

func NewPhotoController(srvcs *services.PhotoService) *PhotoController {
	return &PhotoController{
		photoService: *srvcs,
	}
}

func (p *PhotoController) CreatePhoto(c *gin.Context) {
	var req params.CreatePhoto
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
	res, errCreate := p.photoService.CreatePhoto(&req, id)
	if errCreate != nil {
		c.JSON(errCreate.Status, errCreate)
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

func (p *PhotoController) GetPhotos(c *gin.Context) {
	res, errGet := p.photoService.GetPhotos()
	if errGet != nil {
		c.JSON(errGet.Status, errGet)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (p *PhotoController) UpdatePhoto(c *gin.Context) {
	var req params.CreatePhoto
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err,
		})
		return
	}
	id, _ := strconv.Atoi(c.Param("photoId"))
	res, errUpdate := p.photoService.UpdatePhoto(&req, id)
	if errUpdate != nil {
		c.JSON(errUpdate.Status, errUpdate)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (p *PhotoController) DeletePhoto(c *gin.Context) {
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	res := p.photoService.DeletePhoto(photoId)
	c.JSON(res.Status, params.Response{
		Message: res.Message,
	})
}
