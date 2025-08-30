package main

import (
	l "calcAPI/logging"
	mw "calcAPI/middleware"
	"calcAPI/routes"
	"calcAPI/storage"
	s "calcAPI/storage"
	"fmt"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	zl, err := l.NewLogger()
	if err != nil {
		panic(err)
	}
	defer zl.Sync()

	err = s.EnsureCSV()
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

	r.Use(mw.APIKeyAuth(), mw.SecurityHeaders(), mw.CORS(), mw.RateLimitPerIP(5, 15))
	r.Use(gin.Recovery(), requestid.New(), mw.WithRequestLogger(zl), mw.AccessLog(zl))

	r.GET("/add/:a/:b", routes.Add)
	r.GET("/sub/:a/:b", routes.Sub)
	r.GET("/multiply/:a/:b", routes.Multiply)
	r.GET("/divide/:a/:b", routes.Add)
	r.GET("/token", routes.Apigen)
	r.GET("/health", routes.HealthCheck)
	r.Run(":8080")
}

/* Test */
