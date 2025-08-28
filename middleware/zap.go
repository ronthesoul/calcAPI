package mw

import (
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const RequestIDKey = "X-Request-Id"

func WithRequestLogger(zl *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := requestid.Get(c)
		l := zl
		if rid != "" {
			l = zl.With(zap.String("request_id", rid))
		}
		c.Set("Logger", l)
		c.Next()
	}

}

func AccessLog(zl *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		route := c.FullPath()
		if route == "" {
			route = c.Request.URL.Path
		}

		zl.Info("http_request",
			zap.String("method", c.Request.Method),
			zap.String("route", route),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
			zap.String("client_ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.String("request_id", requestid.Get(c)),
		)
	}
}
