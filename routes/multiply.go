// routes/multiply.go
package routes

import "github.com/gin-gonic/gin"

func Multiply(c *gin.Context) {
	a, b, err := parseTwo(c)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error(), "usage": "/multiply/number1/number2"})
		return
	}
	respond(c, "multiply", a, b, a*b)
}
