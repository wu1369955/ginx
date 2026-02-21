package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Compress 压缩中间件
func Compress() gin.HandlerFunc {
	// 暂时返回一个空的中间件，避免 "incorrect header check" 错误
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Next()
	})
}
