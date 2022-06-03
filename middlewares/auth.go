package middlewares

import (
	"encoding/json"
	"fmt"
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid Token",
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

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdParam, _ := strconv.Atoi(c.Param("userId"))
		userIdStr := c.Request.Header.Get("user_id")
		userId, _ := strconv.Atoi(userIdStr)

		if userId != userIdParam {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Forbidden",
				"message": "You dont have any access",
			})
			return
		}
		c.Next()
	}
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

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		commentId, _ := strconv.Atoi(c.Param("commentId"))
		comment := models.Comment{}

		err := db.Select("user_id").First(&comment, commentId).Error
		if err != nil {
			msg := fmt.Sprintf("Comment with id %v doesnt exist", commentId)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": msg,
			})
			return
		}

		userID := c.Request.Header.Get("user_id")
		currUserIdInt, _ := strconv.Atoi(userID)

		if comment.UserID != uint(currUserIdInt) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You dont have any access",
			})
			return
		}
		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
		socialMedia := models.SocialMedia{}

		err := db.Select("user_id").First(&socialMedia, socialMediaId).Error
		if err != nil {
			msg := fmt.Sprintf("Social Media with id %v doesnt exist", socialMediaId)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": msg,
			})
			return
		}

		userID := c.Request.Header.Get("user_id")
		currUserIdInt, _ := strconv.Atoi(userID)

		if socialMedia.UserID != uint(currUserIdInt) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You dont have any access",
			})
			return
		}
		c.Next()
	}
}
