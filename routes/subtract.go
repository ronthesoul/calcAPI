// routes/sub.go
package routes

import "github.com/gin-gonic/gin"

func Sub(c *gin.Context) {
	a, b, err := parseTwo(c)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error(), "usage": "/sub/number1/number2"})
		return
	}
	respond(c, "sub", a, b, a-b)
}
