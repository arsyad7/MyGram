package controllers

import (
	"mygram/params"
	"mygram/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	socialMediaService services.SocialMediaService
}

func NewSocialMediaController(srvcs *services.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{socialMediaService: *srvcs}
}

func (s *SocialMediaController) CreateSocialMedia(c *gin.Context) {
	var req params.CreateSocialMedia
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
	res, errCreate := s.socialMediaService.CreateSocialMedia(&req, id)
	if errCreate != nil {
		c.JSON(errCreate.Status, errCreate)
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

func (s *SocialMediaController) GetSocialMedias(c *gin.Context) {
	res, errGet := s.socialMediaService.GetSocialMedias()
	if errGet != nil {
		c.JSON(errGet.Status, errGet)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (s *SocialMediaController) UpdateSocialMedia(c *gin.Context) {
	var req params.CreateSocialMedia
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err,
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("socialMediaId"))
	res, errUpdate := s.socialMediaService.UpdateSocialMedia(&req, id)
	if errUpdate != nil {
		c.JSON(errUpdate.Status, errUpdate)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (s *SocialMediaController) DeleteSocialMedia(c *gin.Context) {
	socmedId, _ := strconv.Atoi(c.Param("socialMediaId"))
	res := s.socialMediaService.DeleteSocialMedia(socmedId)
	c.JSON(res.Status, params.Response{
		Message: res.Message,
	})
}
