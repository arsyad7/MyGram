package helpers

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

const SECRET_KEY = "rahasia"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(SECRET_KEY))

	return signedToken
}

func VerifyToken(tokenStr string) (interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token invalid")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errors.New("token invalid")
	}

	return token.Claims.(jwt.MapClaims), nil
}
