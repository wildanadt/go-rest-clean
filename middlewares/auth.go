package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildanadt/go-rest-clean/utils/token"
)

func NewAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
