package routes

import (
	"github.com/gin-gonic/gin"
)

func Divide(c *gin.Context) {
	a, b, err := parseTwo(c)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error(), "usage": "/function/number1/number2"})
		return
	}

	if b == 0 {
		c.IndentedJSON(400, gin.H{"error": "Divided by zero"})
	}

	respond(c, "divide", a, b, a/b)

}
