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

	rand.Seed(time.Now().UnixNano())
	apikey := make([]byte, apilen)

	for i := 0; i < apilen; i++ {
		apikey[i] = chars[rand.Intn(len(chars))]
	}

	// save key with 5 minute TTL
	if err := s.AppendKey(string(apikey), 5*time.Minute); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"API-Key": string(apikey)})
}
