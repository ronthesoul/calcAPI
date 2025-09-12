package routes

import (
	s "calcAPI/storage"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func Apigen(c *gin.Context) {
	apilen := 40
	const chars = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789"

	apikey := make([]byte, apilen)
	for i := 0; i < apilen; i++ {
		apikey[i] = chars[rand.Intn(len(chars))]
	}

	if err := s.AppendKey(string(apikey), 1*time.Hour); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"API-Key": string(apikey)})
}
