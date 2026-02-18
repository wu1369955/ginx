package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Session 会话中间件
func Session() gin.HandlerFunc {
	// 暂时返回一个空的中间件，因为无法下载sessions包
	// 待网络恢复后再实现完整的会话中间件
	return func(c *gin.Context) {
		c.Next()
	}
}
