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

func init() { // <- register once
	prometheus.MustRegister(HTTPReqs, HTTPDur)
}

func PrometheusMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		route := c.FullPath()
		if route == "" {
			route = "unknown"
		} // avoid path-based cardinality
		status := strconv.Itoa(c.Writer.Status())
		method := c.Request.Method

		HTTPReqs.WithLabelValues(method, route, status).Inc()
		HTTPDur.WithLabelValues(method, route, status).Observe(time.Since(start).Seconds())
	}
}
