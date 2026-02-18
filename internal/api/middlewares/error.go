package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/errors"
)

// ErrorHandler 错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 处理请求
		c.Next()

		// 检查是否有错误
		if len(c.Errors) > 0 {
			// 获取最后一个错误
			err := c.Errors.Last().Err

			// 检查是否是AppError类型
			if appErr, ok := err.(*errors.AppError); ok {
				// 返回AppError
				c.JSON(appErr.Code, gin.H{
					"code":    appErr.Code,
					"message": appErr.Message,
				})
				return
			}

			// 其他错误返回500
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "内部服务器错误",
			})
		}
	}
}
