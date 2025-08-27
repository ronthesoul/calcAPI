package mw

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type clientLimiter struct {
	lim      *rate.Limiter
	lastSeen time.Time
}

type limiter struct {
	mu sync.Mutex
	m  map[string]*clientLimiter
	r  rate.Limit
	b  int
}

func newLimiter(r rate.Limit, b int) *limiter {
	return &limiter{
		m: make(map[string]*clientLimiter),
		r: r,
		b: b,
	}
}

func (l *limiter) get(ip string) *rate.Limiter {
	l.mu.Lock()
	defer l.mu.Unlock()

	cl, ok := l.m[ip]
	if !ok {
		cl = &clientLimiter{
			lim:      rate.NewLimiter(l.r, l.b),
			lastSeen: time.Now(),
		}
		l.m[ip] = cl
	} else {
		cl.lastSeen = time.Now()
	}
	return cl.lim
}

func RateLimitPerIP(reqPerSec float64, burst int) gin.HandlerFunc {
	l := newLimiter(rate.Limit(reqPerSec), burst)

	// Periodic TTL cleanup of idle entries
	go func() {
		t := time.NewTicker(10 * time.Minute)
		defer t.Stop()

		for range t.C {
			cutoff := time.Now().Add(-30 * time.Minute)
			l.mu.Lock()
			for k, cl := range l.m {
				if cl.lastSeen.Before(cutoff) {
					delete(l.m, k)
				}
			}
			l.mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		ip := c.ClientIP()
		if !l.get(ip).Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			return
		}
		c.Next()
	}
}
