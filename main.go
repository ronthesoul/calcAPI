package main

import (
	mw "calcAPI/middleware"
	"calcAPI/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(mw.APIKeyAuth(), mw.SecurityHeaders(), mw.CORS())

	r.GET("/add/:a/:b", routes.Add)
	r.GET("/sub/:a/:b", routes.Sub)
	r.GET("/multiply/:a/:b", routes.Multiply)
	r.GET("/divide/:a/:b", routes.Add)
	r.Run("localhost:8080")
}
