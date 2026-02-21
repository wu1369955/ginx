package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/api/handlers"
)

// SetupProductionRoutes 设置生产模块路由
func SetupProductionRoutes(router *gin.Engine, productionHandler *handlers.ProductionHandler) {
	// 生产模块API路由组
	production := router.Group("/api/v1/production")
	{
		// 生产订单管理
		orders := production.Group("/orders")
		{
			orders.GET("", productionHandler.GetProductionOrderList)
			orders.GET("/:id", productionHandler.GetProductionOrderDetail)
			orders.POST("", productionHandler.CreateProductionOrder)
			orders.PUT("/:id", productionHandler.UpdateProductionOrder)
			orders.DELETE("/:id", productionHandler.DeleteProductionOrder)
		}

		// 生产工单管理
		tickets := production.Group("/workorders")
		{
			tickets.GET("", productionHandler.GetWorkOrderList)
			tickets.GET("/:id", productionHandler.GetWorkOrderDetail)
			tickets.POST("", productionHandler.CreateWorkOrder)
			tickets.PUT("/:id", productionHandler.UpdateWorkOrder)
			tickets.DELETE("/:id", productionHandler.DeleteWorkOrder)
		}

		// 工艺路线管理
		routings := production.Group("/routings")
		{
			routings.GET("", productionHandler.GetRoutingList)
			routings.GET("/:id", productionHandler.GetRoutingDetail)
			routings.POST("", productionHandler.CreateRouting)
			routings.PUT("/:id", productionHandler.UpdateRouting)
			routings.DELETE("/:id", productionHandler.DeleteRouting)
		}

		// 工作中心管理
		workcenters := production.Group("/workcenters")
		{
			workcenters.GET("", productionHandler.GetWorkCenterList)
			workcenters.GET("/:id", productionHandler.GetWorkCenterDetail)
			workcenters.POST("", productionHandler.CreateWorkCenter)
			workcenters.PUT("/:id", productionHandler.UpdateWorkCenter)
			workcenters.DELETE("/:id", productionHandler.DeleteWorkCenter)
		}

		// 物料需求计划管理
		mrp := production.Group("/mrp")
		{
			mrp.POST("/run", productionHandler.RunMRP)
			mrp.GET("/results", productionHandler.GetMRPList)
		}

		// 生产报表管理
		reports := production.Group("/reports")
		{
			reports.GET("/status", productionHandler.GetProductionOrderReport)
			reports.GET("/plan", productionHandler.GetWorkOrderReport)
		}
	}
}
