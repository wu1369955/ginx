package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 请求日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latency := endTime.Sub(startTime)

		// 请求方法
		method := c.Request.Method

		// 请求路由
		path := c.Request.URL.Path

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logMsg := fmt.Sprintf("[GIN] %v | %3d | %13v | %15s | %-7s %s",
			endTime.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)

		// 根据状态码设置日志颜色
		switch {
		case statusCode >= 500:
			fmt.Println("\033[31m" + logMsg + "\033[0m") // 红色
		case statusCode >= 400:
			fmt.Println("\033[33m" + logMsg + "\033[0m") // 黄色
		case statusCode >= 300:
			fmt.Println("\033[36m" + logMsg + "\033[0m") // 青色
		default:
			fmt.Println("\033[32m" + logMsg + "\033[0m") // 绿色
		}
	}
}
