package middlewares

import (
	"github.com/gin-gonic/gin"
)

// RequestID 请求ID中间件
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 暂时返回一个空的中间件，因为无法下载uuid包
		// 待网络恢复后再实现完整的请求ID中间件
		c.Next()
	}
}
