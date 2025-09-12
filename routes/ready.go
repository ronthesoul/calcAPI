package routes

import (
	"net/http"

	"calcAPI/storage"

	"github.com/gin-gonic/gin"
)

func Readyz(c *gin.Context) {
	if err := storage.SelfTest(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"ready": false,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ready": true})
}
