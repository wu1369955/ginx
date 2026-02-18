package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/api/handlers"
	"github.com/wu136995/ginx/internal/api/middlewares"
)

// SetupRoutes 设置路由
func SetupRoutes(router *gin.Engine) {
	// 公共路由组
	public := router.Group("/")
	{
		// 健康检查
		public.GET("/health", handlers.HealthCheck)
	}

	// 需要认证的路由组
	protected := router.Group("/")
	protected.Use(middlewares.Auth())
	{
		// 用户信息
		protected.GET("/user/info", handlers.GetUserInfo)
	}
}
