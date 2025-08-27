package main

import (
	mw "calcAPI/middleware"
	"calcAPI/routes"
	"calcAPI/storage"
	s "calcAPI/storage"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	err := s.EnsureCSV()
	if err != nil {
		fmt.Errorf("failed to create CSV file: %w", err)
	}

	go func() {
		t := time.NewTicker(5 * time.Minute)
		defer t.Stop()
		for range t.C {
			_ = storage.PurgeExpired()
		}
	}()

	r.Use(mw.APIKeyAuth(), mw.SecurityHeaders(), mw.CORS())

	r.GET("/add/:a/:b", routes.Add)
	r.GET("/sub/:a/:b", routes.Sub)
	r.GET("/multiply/:a/:b", routes.Multiply)
	r.GET("/divide/:a/:b", routes.Add)
	r.GET("/token", routes.Apigen)
	r.Run("localhost:8080")
}
