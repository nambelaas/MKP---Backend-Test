package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth, exists := c.Get("auth")
		if !exists {
			return
		}

		dataJwt, ok := auth.(*JwtData)
		if !ok {
			return
		}

		result := *dataJwt

		if result.Role != "Admin" {
			c.JSON(http.StatusForbidden, "hanya bisa diakses oleh admin")
			c.Abort()
			return
		}

		c.Next()
	}
}
