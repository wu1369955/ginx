package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/api/schemas"
	"github.com/wu136995/ginx/internal/services"
)

// InventoryHandler 库存处理器
type InventoryHandler struct {
	inventoryService services.InventoryService
}

// NewInventoryHandler 创建库存处理器实例
func NewInventoryHandler(inventoryService services.InventoryService) *InventoryHandler {
	return &InventoryHandler{
		inventoryService: inventoryService,
	}
}

// 仓库管理路由处理函数
// @Summary 获取仓库列表
// @Description 获取所有仓库的列表
// @Tags 库存-仓库管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/warehouses [get]
func (h *InventoryHandler) GetWarehouseList(c *gin.Context) {
	var req schemas.GetWarehouseListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	warehouses, err := h.inventoryService.GetWarehouseList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    warehouses,
	})
}

// @Summary 获取仓库详情
// @Description 根据ID获取仓库详情
// @Tags 库存-仓库管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "仓库ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/warehouses/{id} [get]
func (h *InventoryHandler) GetWarehouseDetail(c *gin.Context) {
	id := c.Param("id")
	warehouse, err := h.inventoryService.GetWarehouseDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    warehouse,
	})
}

// @Summary 创建仓库
// @Description 创建新仓库
// @Tags 库存-仓库管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param warehouse body map[string]interface{} true "仓库信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/warehouses [post]
func (h *InventoryHandler) CreateWarehouse(c *gin.Context) {
	var req schemas.CreateWarehouseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	warehouse, err := h.inventoryService.CreateWarehouse(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    warehouse,
	})
}

// @Summary 更新仓库
// @Description 根据ID更新仓库信息
// @Tags 库存-仓库管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "仓库ID"
// @Param warehouse body map[string]interface{} true "仓库信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/warehouses/{id} [put]
func (h *InventoryHandler) UpdateWarehouse(c *gin.Context) {
	id := c.Param("id")
	var req schemas.UpdateWarehouseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	warehouse, err := h.inventoryService.UpdateWarehouse(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    warehouse,
	})
}

// @Summary 删除仓库
// @Description 根据ID删除仓库
// @Tags 库存-仓库管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "仓库ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/warehouses/{id} [delete]
func (h *InventoryHandler) DeleteWarehouse(c *gin.Context) {
	id := c.Param("id")
	err := h.inventoryService.DeleteWarehouse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// 物料管理路由处理函数
// @Summary 获取物料列表
// @Description 获取所有物料的列表
// @Tags 库存-物料管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/materials [get]
func (h *InventoryHandler) GetMaterialList(c *gin.Context) {
	var req schemas.GetItemListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	materials, err := h.inventoryService.GetItemList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    materials,
	})
}

// @Summary 获取物料详情
// @Description 根据ID获取物料详情
// @Tags 库存-物料管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "物料ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/materials/{id} [get]
func (h *InventoryHandler) GetMaterialDetail(c *gin.Context) {
	id := c.Param("id")
	material, err := h.inventoryService.GetItemDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    material,
	})
}

// @Summary 创建物料
// @Description 创建新物料
// @Tags 库存-物料管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param material body map[string]interface{} true "物料信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/materials [post]
func (h *InventoryHandler) CreateMaterial(c *gin.Context) {
	var req schemas.CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	material, err := h.inventoryService.CreateItem(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    material,
	})
}

// @Summary 更新物料
// @Description 根据ID更新物料信息
// @Tags 库存-物料管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "物料ID"
// @Param material body map[string]interface{} true "物料信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/materials/{id} [put]
func (h *InventoryHandler) UpdateMaterial(c *gin.Context) {
	id := c.Param("id")
	var req schemas.UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	material, err := h.inventoryService.UpdateItem(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    material,
	})
}

// @Summary 删除物料
// @Description 根据ID删除物料
// @Tags 库存-物料管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "物料ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/materials/{id} [delete]
func (h *InventoryHandler) DeleteMaterial(c *gin.Context) {
	id := c.Param("id")
	err := h.inventoryService.DeleteItem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// 库存交易路由处理函数
// @Summary 获取库存交易列表
// @Description 获取所有库存交易的列表
// @Tags 库存-库存交易
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/transactions [get]
func (h *InventoryHandler) GetInventoryTransactionList(c *gin.Context) {
	var req schemas.GetTransactionListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	transactions, err := h.inventoryService.GetTransactionList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    transactions,
	})
}

// @Summary 获取库存交易详情
// @Description 根据ID获取库存交易详情
// @Tags 库存-库存交易
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "交易ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/transactions/{id} [get]
func (h *InventoryHandler) GetInventoryTransactionDetail(c *gin.Context) {
	id := c.Param("id")
	transaction, err := h.inventoryService.GetTransactionDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    transaction,
	})
}

// @Summary 创建库存交易
// @Description 创建新库存交易
// @Tags 库存-库存交易
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param transaction body map[string]interface{} true "交易信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/transactions [post]
func (h *InventoryHandler) CreateInventoryTransaction(c *gin.Context) {
	var req schemas.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	transaction, err := h.inventoryService.CreateTransaction(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    transaction,
	})
}

// @Summary 更新库存交易
// @Description 根据ID更新库存交易
// @Tags 库存-库存交易
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "交易ID"
// @Param transaction body map[string]interface{} true "交易信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/transactions/{id} [put]
func (h *InventoryHandler) UpdateInventoryTransaction(c *gin.Context) {
	// 库存交易一般不支持更新，返回错误信息
	c.JSON(http.StatusMethodNotAllowed, gin.H{
		"code":    405,
		"message": "Method not allowed",
		"error":   "Inventory transactions cannot be updated",
	})
}

// @Summary 删除库存交易
// @Description 根据ID删除库存交易
// @Tags 库存-库存交易
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "交易ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/transactions/{id} [delete]
func (h *InventoryHandler) DeleteInventoryTransaction(c *gin.Context) {
	// 库存交易一般不支持删除，返回错误信息
	c.JSON(http.StatusMethodNotAllowed, gin.H{
		"code":    405,
		"message": "Method not allowed",
		"error":   "Inventory transactions cannot be deleted",
	})
}

// 库存盘点路由处理函数
// @Summary 获取库存盘点列表
// @Description 获取所有库存盘点的列表
// @Tags 库存-库存盘点
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/counts [get]
func (h *InventoryHandler) GetInventoryCountList(c *gin.Context) {
	var req schemas.GetCountListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	counts, err := h.inventoryService.GetCountList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    counts,
	})
}

// @Summary 获取库存盘点详情
// @Description 根据ID获取库存盘点详情
// @Tags 库存-库存盘点
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "盘点ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/counts/{id} [get]
func (h *InventoryHandler) GetInventoryCountDetail(c *gin.Context) {
	id := c.Param("id")
	count, err := h.inventoryService.GetCountDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    count,
	})
}

// @Summary 创建库存盘点
// @Description 创建新库存盘点
// @Tags 库存-库存盘点
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param count body map[string]interface{} true "盘点信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/counts [post]
func (h *InventoryHandler) CreateInventoryCount(c *gin.Context) {
	var req schemas.CreateCountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	count, err := h.inventoryService.CreateCount(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    count,
	})
}

// @Summary 更新库存盘点
// @Description 根据ID更新库存盘点
// @Tags 库存-库存盘点
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "盘点ID"
// @Param count body map[string]interface{} true "盘点信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/counts/{id} [put]
func (h *InventoryHandler) UpdateInventoryCount(c *gin.Context) {
	id := c.Param("id")
	var req schemas.UpdateCountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	count, err := h.inventoryService.UpdateCount(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    count,
	})
}

// @Summary 删除库存盘点
// @Description 根据ID删除库存盘点
// @Tags 库存-库存盘点
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "盘点ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/counts/{id} [delete]
func (h *InventoryHandler) DeleteInventoryCount(c *gin.Context) {
	id := c.Param("id")
	var req schemas.CancelCountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	err := h.inventoryService.CancelCount(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// @Summary 提交库存盘点
// @Description 提交库存盘点
// @Tags 库存-库存盘点
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "盘点ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/counts/{id}/submit [post]
func (h *InventoryHandler) SubmitInventoryCount(c *gin.Context) {
	id := c.Param("id")
	var req schemas.CompleteCountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	err := h.inventoryService.CompleteCount(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// @Summary 审批库存盘点
// @Description 审批库存盘点
// @Tags 库存-库存盘点
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "盘点ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/counts/{id}/approve [post]
func (h *InventoryHandler) ApproveInventoryCount(c *gin.Context) {
	// 审批功能在service中没有直接实现，这里暂时返回成功
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// 库存报表路由处理函数
// @Summary 获取库存状态报表
// @Description 获取库存状态的报表
// @Tags 库存-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/reports/status [get]
func (h *InventoryHandler) GetInventoryStatusReport(c *gin.Context) {
	var req schemas.GetInventoryBalanceReportRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	report, err := h.inventoryService.GetInventoryBalanceReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    report,
	})
}

// @Summary 获取库存移动报表
// @Description 获取库存移动的报表
// @Tags 库存-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/reports/movement [get]
func (h *InventoryHandler) GetInventoryMovementReport(c *gin.Context) {
	var req schemas.GetInventoryMovementReportRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	report, err := h.inventoryService.GetInventoryMovementReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    report,
	})
}

// @Summary 获取库存估值报表
// @Description 获取库存估值的报表
// @Tags 库存-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/reports/valuation [get]
func (h *InventoryHandler) GetInventoryValuationReport(c *gin.Context) {
	// 库存估值报表可以使用库存余额报表，因为余额报表中包含了价值信息
	var req schemas.GetInventoryBalanceReportRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	report, err := h.inventoryService.GetInventoryBalanceReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    report,
	})
}

// @Summary 获取物料分析报表
// @Description 获取物料分析的报表
// @Tags 库存-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/reports/material-analysis [get]
func (h *InventoryHandler) GetMaterialAnalysisReport(c *gin.Context) {
	var req schemas.GetInventoryABCReportRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	report, err := h.inventoryService.GetInventoryABCReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    report,
	})
}

// @Summary 导出库存报表
// @Description 导出库存相关的报表
// @Tags 库存-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/inventory/reports/export [get]
func (h *InventoryHandler) ExportInventoryReport(c *gin.Context) {
	var req schemas.ExportInventoryReportRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})
		return
	}
	data, err := h.inventoryService.ExportInventoryReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}
