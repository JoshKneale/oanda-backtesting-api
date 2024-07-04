package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" || !validateToken(token) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func validateToken(token string) bool {
	return token == "valid-token" // TODO
}
