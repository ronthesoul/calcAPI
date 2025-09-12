package main

import (
	l "calcAPI/logging"
	mw "calcAPI/middleware"
	"calcAPI/routes"
	"calcAPI/storage"
	s "calcAPI/storage"
	"os"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	prometheus.MustRegister(mw.HTTPReqs, mw.HTTPDur)
	r := gin.New()
	zl, err := l.NewLogger()
	if err != nil {
		panic(err)
	}
	defer func() { _ = zl.Sync() }()

	if err = s.EnsureCSV(); err != nil {
		zl.Error("failed to create CSV file", zap.Error(err))
		os.Exit(1)
	}

	go func() {
		t := time.NewTicker(2 * time.Hour)
		defer t.Stop()
		for range t.C {
			_ = storage.PurgeExpired()
		}
	}()

	r.Use(mw.APIKeyAuth(), mw.SecurityHeaders(), mw.CORS(), mw.RateLimitPerIP(5, 15))
	r.Use(gin.Recovery(), requestid.New(), mw.WithRequestLogger(zl), mw.AccessLog(zl), mw.PrometheusMetrics())

	r.GET("/add/:a/:b", routes.Add)
	r.GET("/sub/:a/:b", routes.Sub)
	r.GET("/multiply/:a/:b", routes.Multiply)
	r.GET("/divide/:a/:b", routes.Divide)
	r.GET("/token", routes.Apigen)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/healthz", routes.HealthCheck)
	r.GET("/readyz", routes.Readyz)

	if err := r.Run(":8080"); err != nil {
		zl.Fatal("server failed to start", zap.Error(err))
	}
}
