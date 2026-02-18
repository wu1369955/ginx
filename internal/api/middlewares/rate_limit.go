package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter 简单的速率限制器
type RateLimiter struct {
	ips         map[string][]time.Time
	mu          sync.Mutex
	limit       int           // 限制时间内允许的请求数
	window      time.Duration // 限制时间窗口
	cleanupTime time.Duration // 清理过期记录的时间间隔
}

// NewRateLimiter 创建一个新的速率限制器
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	limiter := &RateLimiter{
		ips:         make(map[string][]time.Time),
		limit:       limit,
		window:      window,
		cleanupTime: time.Minute * 5,
	}

	// 启动定期清理过期记录的协程
	go limiter.cleanup()

	return limiter
}

// cleanup 定期清理过期的记录
func (rl *RateLimiter) cleanup() {
	for {
		time.Sleep(rl.cleanupTime)

		rl.mu.Lock()
		for ip, times := range rl.ips {
			var validTimes []time.Time
			now := time.Now()

			for _, t := range times {
				if now.Sub(t) < rl.window {
					validTimes = append(validTimes, t)
				}
			}

			if len(validTimes) == 0 {
				delete(rl.ips, ip)
			} else {
				rl.ips[ip] = validTimes
			}
		}
		rl.mu.Unlock()
	}
}

// Allow 检查是否允许请求
func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	// 获取该IP的请求时间记录
	times := rl.ips[ip]

	// 清理过期的记录
	var validTimes []time.Time
	for _, t := range times {
		if now.Sub(t) < rl.window {
			validTimes = append(validTimes, t)
		}
	}

	// 检查是否超过限制
	if len(validTimes) >= rl.limit {
		rl.ips[ip] = validTimes
		return false
	}

	// 添加新的请求时间记录
	validTimes = append(validTimes, now)
	rl.ips[ip] = validTimes

	return true
}

// RateLimit 限流中间件
func RateLimit() gin.HandlerFunc {
	// 创建一个速率限制器，每分钟允许60个请求
	limiter := NewRateLimiter(60, time.Minute)

	return func(c *gin.Context) {
		// 获取客户端IP
		clientIP := c.ClientIP()

		// 检查是否允许请求
		if !limiter.Allow(clientIP) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code":    http.StatusTooManyRequests,
				"message": "请求过于频繁，请稍后再试",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
