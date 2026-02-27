package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/services"
)

// SearchHandler 搜索处理器
type SearchHandler struct {
	searchService services.SearchService
}

// NewSearchHandler 创建搜索处理器实例
func NewSearchHandler(searchService services.SearchService) *SearchHandler {
	return &SearchHandler{
		searchService: searchService,
	}
}

// SearchByKeyword 关键词搜索
// @Summary 关键词搜索
// @Description 使用关键词搜索数据（支持客户、产品、订单等类型）
// @Tags 搜索
// @Accept json
// @Produce json
// @Param request body map[string]interface{} true "搜索请求参数"
// @Success 200 {object} map[string]interface{} "搜索结果"
// @Failure 400 {object} map[string]interface{} "请求参数错误"
// @Failure 500 {object} map[string]interface{} "服务器内部错误"
// @Router /api/search/keyword [post]
func (h *SearchHandler) SearchByKeyword(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	result, err := h.searchService.SearchByKeyword(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetRankingData 数量排行数据
// @Summary 数量排行数据
// @Description 获取数量排行数据（支持客户订单数量、产品销量、客户销售额等类型）
// @Tags 搜索
// @Accept json
// @Produce json
// @Param request body map[string]interface{} true "排行请求参数"
// @Success 200 {object} map[string]interface{} "排行结果"
// @Failure 400 {object} map[string]interface{} "请求参数错误"
// @Failure 500 {object} map[string]interface{} "服务器内部错误"
// @Router /api/search/ranking [post]
func (h *SearchHandler) GetRankingData(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	result, err := h.searchService.GetRankingData(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
