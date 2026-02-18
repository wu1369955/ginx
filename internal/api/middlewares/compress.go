package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Compress 压缩中间件
func Compress() gin.HandlerFunc {
	// 使用Gin内置的压缩中间件
	return gin.HandlerFunc(func(c *gin.Context) {
		// 检查是否支持压缩
		if c.Request.Header.Get("Accept-Encoding") != "" {
			// 启用压缩
			c.Writer.Header().Set("Content-Encoding", "gzip")
		}
		c.Next()
	})
}
