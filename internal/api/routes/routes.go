package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/api/handlers"
	"github.com/wu136995/ginx/internal/api/middlewares"
)

// SetupRoutes 设置路由
func SetupRoutes(router *gin.Engine, userHandler *handlers.UserHandler, salesHandler *handlers.SalesHandler, inventoryHandler *handlers.InventoryHandler, purchaseHandler *handlers.PurchaseHandler, financeHandler *handlers.FinanceHandler, productionHandler *handlers.ProductionHandler, crmHandler *handlers.CRMHandler) {
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
		protected.GET("/user/info", userHandler.GetUserInfo)
	}

	// 销售模块路由
	SetupSalesRoutes(router, salesHandler)
	// 库存模块路由
	SetupInventoryRoutes(router, inventoryHandler)
	// 采购模块路由
	SetupPurchaseRoutes(router, purchaseHandler)
	// 财务模块路由
	SetupFinanceRoutes(router, financeHandler)
	// 生产模块路由
	SetupProductionRoutes(router, productionHandler)
	// // 人力资源模块路由
	// SetupHRRoutes(router, hrHandler)
	// CRM模块路由
	SetupCRMRoutes(router, crmHandler)
}
