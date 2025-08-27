package mw

import (
	"calcAPI/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.FullPath() == "/token" || c.FullPath() == "/health" {
			c.Next()
			return
		}

		got := c.GetHeader("X-API-Key")
		valid, err := storage.IsKeyValid(got)
		if err != nil || !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired API key"})
			return
		}

		c.Next()
	}
}
