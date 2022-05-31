package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"mygram/helpers"
	"mygram/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) error {
	authorization := c.GetHeader("Authorization")

	errResp := errors.New("invalid token")
	if authorization == "" {
		return errResp
	}
	token := strings.Split(authorization, " ")
	if len(token) != 2 {
		return errResp
	}

	if token[0] != "Bearer" {
		return errResp
	}

	payload, err := helpers.VerifyToken(token[1])
	if err != nil {
		return errResp
	}

	var user models.User
	jsonBody, _ := json.Marshal(payload)
	_ = json.Unmarshal(jsonBody, &user)
	c.Request.Header.Add("user_id", fmt.Sprint(user.ID))

	return nil
}
