package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/config"
	"github.com/wu136995/ginx/internal/utils"
)

// Auth 认证中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "未提供认证信息",
			})
			c.Abort()
			return
		}

		// 检查token格式
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "认证格式错误",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 从配置中获取JWT密钥
		jwtSecret := config.AppConfig.JWT.Secret

		// 验证token
		claims, err := utils.ParseToken(tokenString, jwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("userID", claims.UserID)
		c.Set("userName", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
