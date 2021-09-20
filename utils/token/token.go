package token

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/wildanadt/go-rest-clean/config/env"
)

func GenerateToken(user_id uint) (string, error) {
	token_lifespan, err := strconv.Atoi(env.Config.JWTConfig.TokenLifespan)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	_token, err := at.SignedString([]byte(env.Config.JWTConfig.APISecret))
	if err != nil {
		return "", err
	}
	return _token, nil

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(env.Config.JWTConfig.APISecret), nil
	})

	if err != nil {
		return err
	}

	return nil
}

func ExtractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
