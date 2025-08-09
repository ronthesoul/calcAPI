package routes

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Multiply(c *gin.Context) {
	aStr := c.Param("a")
	bStr := c.Param("b")

	if strings.TrimSpace(aStr) == "" || strings.TrimSpace(bStr) == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Usage": "/multiply/number1/number2",
		})
		return
	}

	a, err := strconv.Atoi(aStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Usage": "/multiply/number1/number2",
			"error": err.Error(),
		})
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Usage": "/multiply/number1/number2",
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"Result": a * b,
	})
}
