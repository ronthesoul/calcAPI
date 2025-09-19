package main

import (
	l "calcAPI/logging"
	mw "calcAPI/middleware"
	"calcAPI/routes"
	"calcAPI/storage"
	"os"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	r := gin.New()

	zl, err := l.NewLogger()
	if err != nil {
		panic(err)
	}
	defer func() { _ = zl.Sync() }()

	if err = storage.EnsureCSV(); err != nil {
		zl.Error("failed to create CSV file", zap.Error(err))
		os.Exit(1)
	}

	// background cleanup
	go func() {
		t := time.NewTicker(1 * time.Hour)
		defer t.Stop()
		for range t.C {
			_ = storage.PurgeExpired()
		}
	}()

	// global middlewares (no auth)
	r.Use(gin.Recovery(), requestid.New(), mw.WithRequestLogger(zl), mw.AccessLog(zl), mw.PrometheusMetrics())

	// PUBLIC endpoints (no API key)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/healthz", routes.HealthCheck)
	r.GET("/readyz", routes.Readyz)

	//  PROTECTED API group
	api := r.Group("/")
	api.Use(mw.APIKeyAuth(), mw.SecurityHeaders(), mw.CORS(), mw.RateLimitPerIP(5, 15))
	api.GET("/add/:a/:b", routes.Add)
	api.GET("/sub/:a/:b", routes.Sub)
	api.GET("/multiply/:a/:b", routes.Multiply)
	api.GET("/divide/:a/:b", routes.Divide)
	api.GET("/token", routes.Apigen)

	if err := r.Run(":8080"); err != nil {
		zl.Fatal("server failed to start", zap.Error(err))
	}
}
