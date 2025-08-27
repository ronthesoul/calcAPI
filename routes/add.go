package routes

import "github.com/gin-gonic/gin"

func Add(c *gin.Context) {
	a, b, err := parseTwo(c)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error(), "usage": "/add/number1/number2"})
		return
	}
	respond(c, "add", a, b, a+b)
}
