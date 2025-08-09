package main

import (
	"calcAPI/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/add/:a/:b", routes.Add)
	router.GET("/sub/:a/:b", routes.Sub)
	router.GET("/multiply/:a/:b", routes.Multiply)
	router.GET("/divide/:a/:b", routes.Add)
	router.Run("localhost:8080")
}
