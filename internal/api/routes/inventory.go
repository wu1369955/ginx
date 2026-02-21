package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/api/handlers"
)

// SetupInventoryRoutes 设置库存模块路由
func SetupInventoryRoutes(router *gin.Engine, inventoryHandler *handlers.InventoryHandler) {
	// 库存模块API路由组
	inventory := router.Group("/api/v1/inventory")
	{
		// 仓库管理
		warehouses := inventory.Group("/warehouses")
		{
			warehouses.GET("", inventoryHandler.GetWarehouseList)
			warehouses.GET("/:id", inventoryHandler.GetWarehouseDetail)
			warehouses.POST("", inventoryHandler.CreateWarehouse)
			warehouses.PUT("/:id", inventoryHandler.UpdateWarehouse)
			warehouses.DELETE("/:id", inventoryHandler.DeleteWarehouse)
		}

		// 物料管理
		items := inventory.Group("/items")
		{
			items.GET("", inventoryHandler.GetMaterialList)
			items.GET("/:id", inventoryHandler.GetMaterialDetail)
			items.POST("", inventoryHandler.CreateMaterial)
			items.PUT("/:id", inventoryHandler.UpdateMaterial)
			items.DELETE("/:id", inventoryHandler.DeleteMaterial)
		}

		// 库存交易管理
		transactions := inventory.Group("/transactions")
		{
			transactions.GET("", inventoryHandler.GetInventoryTransactionList)
			transactions.GET("/:id", inventoryHandler.GetInventoryTransactionDetail)
			transactions.POST("", inventoryHandler.CreateInventoryTransaction)
		}

		// 库存盘点管理
		counts := inventory.Group("/counts")
		{
			counts.GET("", inventoryHandler.GetInventoryCountList)
			counts.GET("/:id", inventoryHandler.GetInventoryCountDetail)
			counts.POST("", inventoryHandler.CreateInventoryCount)
			counts.PUT("/:id", inventoryHandler.UpdateInventoryCount)
			counts.DELETE("/:id", inventoryHandler.DeleteInventoryCount)
			counts.POST("/:id/submit", inventoryHandler.SubmitInventoryCount)
			counts.POST("/:id/approve", inventoryHandler.ApproveInventoryCount)
		}

		// 库存报表管理
		reports := inventory.Group("/reports")
		{
			reports.GET("/balance", inventoryHandler.GetInventoryStatusReport)
			reports.GET("/movement", inventoryHandler.GetInventoryMovementReport)
			reports.GET("/valuation", inventoryHandler.GetInventoryValuationReport)
			reports.GET("/analysis", inventoryHandler.GetMaterialAnalysisReport)
			reports.GET("/export", inventoryHandler.ExportInventoryReport)
		}
	}
}
