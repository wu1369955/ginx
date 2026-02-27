package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/api/handlers"
)

// SetupSearchRoutes 设置搜索相关路由
func SetupSearchRoutes(router *gin.Engine, searchHandler *handlers.SearchHandler) {
	// 搜索相关路由组
	searchGroup := router.Group("/api/search")
	{
		// 关键词搜索
		searchGroup.POST("/keyword", searchHandler.SearchByKeyword)
		
		// 数量排行数据
		searchGroup.POST("/ranking", searchHandler.GetRankingData)
	}
}
