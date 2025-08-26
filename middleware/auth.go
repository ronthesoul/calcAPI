package mw

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func APIKeyAuth() gin.HandlerFunc {
	required := os.Getenv("CALC_API_KEY")
	return func(c *gin.Context) {
		if required == "" {
			c.Next()
			return
		}
		got := c.GetHeader("X-API-Key")
		if got == "" || got != required {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid API key"})
			return
		}
		c.Next()
	}
}
