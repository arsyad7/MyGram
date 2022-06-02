package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"mygram/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		errResp := errors.New("invalid token")
		if authorization == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": errResp,
			})
			return
		}
		token := strings.Split(authorization, " ")
		if len(token) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid Token",
			})
			return
		}

		if token[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid Token",
			})
			return
		}

		payload, err := helpers.VerifyToken(token[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid Token",
			})
			return
		}

		var user models.User
		jsonBody, _ := json.Marshal(payload)
		_ = json.Unmarshal(jsonBody, &user)
		c.Request.Header.Add("user_id", fmt.Sprint(user.ID))

		c.Next()
	}
}

func AuthorizationPhoto(c *gin.Context, u services.PhotoService) bool {
	userID := c.Request.Header.Get("user_id")
	photoId := c.Param("photoId")

	id, _ := strconv.Atoi(photoId)
	currUserIdInt, _ := strconv.Atoi(userID)
	photo, err := u.GetPhotoById(id)
	if err != nil {
		return false
	}

	if photo.UserID != uint(currUserIdInt) {
		return false
	}

	return true
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		photoId, _ := strconv.Atoi(c.Param("photoId"))
		photo := models.Photo{}

		err := db.Select("user_id").First(&photo, uint(photoId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "photo doesnt exist",
			})
			return
		}

		userID := c.Request.Header.Get("user_id")
		currUserIdInt, _ := strconv.Atoi(userID)
		if photo.UserID != uint(currUserIdInt) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You dont have any access",
			})
			return
		}
		c.Next()
	}
}
