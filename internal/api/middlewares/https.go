package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPSRedirect HTTPS重定向中间件
func HTTPSRedirect() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查请求是否为HTTP
		if c.Request.URL.Scheme == "http" || c.Request.Header.Get("X-Forwarded-Proto") == "http" {
			// 构建HTTPS URL
			httpsURL := "https://" + c.Request.Host + c.Request.URL.Path
			if c.Request.URL.RawQuery != "" {
				httpsURL += "?" + c.Request.URL.RawQuery
			}
			// 重定向到HTTPS
			c.Redirect(http.StatusMovedPermanently, httpsURL)
			c.Abort()
			return
		}
		c.Next()
	}
}
