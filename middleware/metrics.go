package mw

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HTTPReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "http_requests_total", Help: "Total HTTP requests"},
		[]string{"method", "route", "status"},
	)
	HTTPDur = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{Name: "http_request_duration_seconds", Help: "Request duration (s)"},
		[]string{"method", "route", "status"},
	)
)

func PrometheusMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		route := c.FullPath()
		if route == "" {
			route = c.Request.URL.Path
		}
		HTTPReqs.WithLabelValues(c.Request.Method, route, strconv.Itoa(c.Writer.Status())).Inc()
		HTTPDur.WithLabelValues(c.Request.Method, route, strconv.Itoa(c.Writer.Status())).Observe(time.Since(start).Seconds())
	}
}
