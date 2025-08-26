package routes

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func parseTwo(c *gin.Context) (float64, float64, error) {

	aStr := strings.TrimSpace(c.Param("a"))
	bStr := strings.TrimSpace(c.Param("b"))

	if aStr == "" || bStr == "" {
		return 0, 0, errors.New("usage: /<op>/number1/number2")
	}
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		return 0, 0, err
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		return 0, 0, err
	}

	return a, b, nil
}

func respond(c *gin.Context, op string, a, b, res float64) {
	c.IndentedJSON(200, gin.H{"op": op, "a": a, "b": b, "result": res})
}
