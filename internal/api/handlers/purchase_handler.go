package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/services"
)

// PurchaseHandler 采购处理器
type PurchaseHandler struct {
	purchaseService services.PurchaseService
}

// NewPurchaseHandler 创建采购处理器实例
func NewPurchaseHandler(purchaseService services.PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{
		purchaseService: purchaseService,
	}
}

// 供应商管理路由处理函数
// @Summary 获取供应商列表
// @Description 获取所有供应商的列表
// @Tags 采购-供应商管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/suppliers [get]
func (h *PurchaseHandler) GetSupplierList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	suppliers, err := h.purchaseService.GetSupplierList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get supplier list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    suppliers,
	})
}

// @Summary 获取供应商详情
// @Description 根据ID获取供应商详情
// @Tags 采购-供应商管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "供应商ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/suppliers/{id} [get]
func (h *PurchaseHandler) GetSupplierDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	supplier, err := h.purchaseService.GetSupplierDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get supplier detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    supplier,
	})
}

// @Summary 创建供应商
// @Description 创建新供应商
// @Tags 采购-供应商管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param supplier body map[string]interface{} true "供应商信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/suppliers [post]
func (h *PurchaseHandler) CreateSupplier(c *gin.Context) {
	// 解析请求体
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	supplier, err := h.purchaseService.CreateSupplier(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create supplier: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    supplier,
	})
}

// @Summary 更新供应商
// @Description 根据ID更新供应商信息
// @Tags 采购-供应商管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "供应商ID"
// @Param supplier body map[string]interface{} true "供应商信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/suppliers/{id} [put]
func (h *PurchaseHandler) UpdateSupplier(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 解析请求体
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	supplier, err := h.purchaseService.UpdateSupplier(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update supplier: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    supplier,
	})
}

// @Summary 删除供应商
// @Description 根据ID删除供应商
// @Tags 采购-供应商管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "供应商ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/suppliers/{id} [delete]
func (h *PurchaseHandler) DeleteSupplier(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.DeleteSupplier(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete supplier: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// @Summary 获取供应商联系人列表
// @Description 获取供应商的联系人列表
// @Tags 采购-供应商管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "供应商ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/suppliers/{id}/contacts [get]
func (h *PurchaseHandler) GetSupplierContacts(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	contacts, err := h.purchaseService.GetSupplierContacts(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get supplier contacts: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    contacts,
	})
}

// @Summary 添加供应商联系人
// @Description 为供应商添加联系人
// @Tags 采购-供应商管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "供应商ID"
// @Param contact body map[string]interface{} true "联系人信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/suppliers/{id}/contacts [post]
func (h *PurchaseHandler) AddSupplierContact(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 解析请求体
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	contact, err := h.purchaseService.AddSupplierContact(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to add supplier contact: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    contact,
	})
}

// 采购申请管理路由处理函数
// @Summary 获取采购申请列表
// @Description 获取所有采购申请的列表
// @Tags 采购-采购申请
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/requisitions [get]
func (h *PurchaseHandler) GetRequisitionList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	requisitions, err := h.purchaseService.GetRequisitionList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get requisition list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    requisitions,
	})
}

// @Summary 获取采购申请详情
// @Description 根据ID获取采购申请详情
// @Tags 采购-采购申请
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购申请ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/requisitions/{id} [get]
func (h *PurchaseHandler) GetRequisitionDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	requisition, err := h.purchaseService.GetRequisitionDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get requisition detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    requisition,
	})
}

// @Summary 创建采购申请
// @Description 创建新采购申请
// @Tags 采购-采购申请
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param requisition body map[string]interface{} true "采购申请信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/requisitions [post]
func (h *PurchaseHandler) CreateRequisition(c *gin.Context) {
	// 解析请求体
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	requisition, err := h.purchaseService.CreateRequisition(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create requisition: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    requisition,
	})
}

// @Summary 更新采购申请
// @Description 根据ID更新采购申请
// @Tags 采购-采购申请
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购申请ID"
// @Param requisition body map[string]interface{} true "采购申请信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/requisitions/{id} [put]
func (h *PurchaseHandler) UpdateRequisition(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 解析请求体
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	requisition, err := h.purchaseService.UpdateRequisition(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update requisition: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    requisition,
	})
}

// @Summary 删除采购申请
// @Description 根据ID删除采购申请
// @Tags 采购-采购申请
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购申请ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/requisitions/{id} [delete]
func (h *PurchaseHandler) DeleteRequisition(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.DeleteRequisition(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete requisition: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// @Summary 提交采购申请
// @Description 提交采购申请
// @Tags 采购-采购申请
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购申请ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/requisitions/{id}/submit [post]
func (h *PurchaseHandler) SubmitRequisition(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.SubmitRequisition(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to submit requisition: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// @Summary 审批采购申请
// @Description 审批采购申请
// @Tags 采购-采购申请
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购申请ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/requisitions/{id}/approve [post]
func (h *PurchaseHandler) ApproveRequisition(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.ApproveRequisition(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to approve requisition: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// @Summary 拒绝采购申请
// @Description 拒绝采购申请
// @Tags 采购-采购申请
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购申请ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/requisitions/{id}/reject [post]
func (h *PurchaseHandler) RejectRequisition(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.RejectRequisition(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to reject requisition: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// 采购订单管理路由处理函数
// @Summary 获取采购订单列表
// @Description 获取所有采购订单的列表
// @Tags 采购-采购订单
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/orders [get]
func (h *PurchaseHandler) GetPurchaseOrderList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	orders, err := h.purchaseService.GetPurchaseOrderList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get purchase order list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    orders,
	})
}

// @Summary 获取采购订单详情
// @Description 根据ID获取采购订单详情
// @Tags 采购-采购订单
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购订单ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/orders/{id} [get]
func (h *PurchaseHandler) GetPurchaseOrderDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	order, err := h.purchaseService.GetPurchaseOrderDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get purchase order detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    order,
	})
}

// @Summary 创建采购订单
// @Description 创建新采购订单
// @Tags 采购-采购订单
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body map[string]interface{} true "采购订单信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/orders [post]
func (h *PurchaseHandler) CreatePurchaseOrder(c *gin.Context) {
	// 解析请求体
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	order, err := h.purchaseService.CreatePurchaseOrder(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create purchase order: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    order,
	})
}

// @Summary 更新采购订单
// @Description 根据ID更新采购订单
// @Tags 采购-采购订单
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购订单ID"
// @Param order body map[string]interface{} true "采购订单信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/orders/{id} [put]
func (h *PurchaseHandler) UpdatePurchaseOrder(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 解析请求体
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	order, err := h.purchaseService.UpdatePurchaseOrder(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update purchase order: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    order,
	})
}

// @Summary 删除采购订单
// @Description 根据ID删除采购订单
// @Tags 采购-采购订单
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购订单ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/orders/{id} [delete]
func (h *PurchaseHandler) DeletePurchaseOrder(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.DeletePurchaseOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete purchase order: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// @Summary 提交采购订单
// @Description 提交采购订单
// @Tags 采购-采购订单
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购订单ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/orders/{id}/submit [post]
func (h *PurchaseHandler) SubmitPurchaseOrder(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.SubmitPurchaseOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to submit purchase order: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// @Summary 审批采购订单
// @Description 审批采购订单
// @Tags 采购-采购订单
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购订单ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/orders/{id}/approve [post]
func (h *PurchaseHandler) ApprovePurchaseOrder(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.ApprovePurchaseOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to approve purchase order: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// @Summary 拒绝采购订单
// @Description 拒绝采购订单
// @Tags 采购-采购订单
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购订单ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/orders/{id}/reject [post]
func (h *PurchaseHandler) RejectPurchaseOrder(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.RejectPurchaseOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to reject purchase order: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// @Summary 关闭采购订单
// @Description 关闭采购订单
// @Tags 采购-采购订单
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购订单ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/orders/{id}/close [post]
func (h *PurchaseHandler) ClosePurchaseOrder(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.ClosePurchaseOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to close purchase order: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// 采购收货管理路由处理函数
// @Summary 获取采购收货列表
// @Description 获取所有采购收货的列表
// @Tags 采购-采购收货
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/receipts [get]
func (h *PurchaseHandler) GetPurchaseReceiptList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	receipts, err := h.purchaseService.GetPurchaseReceiptList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get purchase receipt list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    receipts,
	})
}

// @Summary 获取采购收货详情
// @Description 根据ID获取采购收货详情
// @Tags 采购-采购收货
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购收货ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/receipts/{id} [get]
func (h *PurchaseHandler) GetPurchaseReceiptDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	receipt, err := h.purchaseService.GetPurchaseReceiptDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get purchase receipt detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    receipt,
	})
}

// @Summary 创建采购收货
// @Description 创建新采购收货
// @Tags 采购-采购收货
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param receipt body map[string]interface{} true "采购收货信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/receipts [post]
func (h *PurchaseHandler) CreatePurchaseReceipt(c *gin.Context) {
	// 解析请求体
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	receipt, err := h.purchaseService.CreatePurchaseReceipt(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create purchase receipt: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    receipt,
	})
}

// @Summary 更新采购收货
// @Description 根据ID更新采购收货
// @Tags 采购-采购收货
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购收货ID"
// @Param receipt body map[string]interface{} true "采购收货信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/receipts/{id} [put]
func (h *PurchaseHandler) UpdatePurchaseReceipt(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 解析请求体
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	receipt, err := h.purchaseService.UpdatePurchaseReceipt(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update purchase receipt: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    receipt,
	})
}

// @Summary 删除采购收货
// @Description 根据ID删除采购收货
// @Tags 采购-采购收货
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购收货ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/receipts/{id} [delete]
func (h *PurchaseHandler) DeletePurchaseReceipt(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.DeletePurchaseReceipt(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete purchase receipt: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// @Summary 完成采购收货
// @Description 完成采购收货
// @Tags 采购-采购收货
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购收货ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/receipts/{id}/complete [post]
func (h *PurchaseHandler) CompletePurchaseReceipt(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.CompletePurchaseReceipt(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to complete purchase receipt: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// 采购发票管理路由处理函数
// @Summary 获取采购发票列表
// @Description 获取所有采购发票的列表
// @Tags 采购-采购发票
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/invoices [get]
func (h *PurchaseHandler) GetPurchaseInvoiceList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	invoices, err := h.purchaseService.GetPurchaseInvoiceList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get purchase invoice list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    invoices,
	})
}

// @Summary 获取采购发票详情
// @Description 根据ID获取采购发票详情
// @Tags 采购-采购发票
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购发票ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/invoices/{id} [get]
func (h *PurchaseHandler) GetPurchaseInvoiceDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	invoice, err := h.purchaseService.GetPurchaseInvoiceDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get purchase invoice detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    invoice,
	})
}

// @Summary 创建采购发票
// @Description 创建新采购发票
// @Tags 采购-采购发票
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param invoice body map[string]interface{} true "采购发票信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/invoices [post]
func (h *PurchaseHandler) CreatePurchaseInvoice(c *gin.Context) {
	// 解析请求体
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	invoice, err := h.purchaseService.CreatePurchaseInvoice(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create purchase invoice: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    invoice,
	})
}

// @Summary 更新采购发票
// @Description 根据ID更新采购发票
// @Tags 采购-采购发票
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购发票ID"
// @Param invoice body map[string]interface{} true "采购发票信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/invoices/{id} [put]
func (h *PurchaseHandler) UpdatePurchaseInvoice(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 解析请求体
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	invoice, err := h.purchaseService.UpdatePurchaseInvoice(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update purchase invoice: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    invoice,
	})
}

// @Summary 删除采购发票
// @Description 根据ID删除采购发票
// @Tags 采购-采购发票
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购发票ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/invoices/{id} [delete]
func (h *PurchaseHandler) DeletePurchaseInvoice(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.DeletePurchaseInvoice(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete purchase invoice: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// @Summary 验证采购发票
// @Description 验证采购发票
// @Tags 采购-采购发票
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购发票ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/invoices/{id}/verify [post]
func (h *PurchaseHandler) VerifyPurchaseInvoice(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.VerifyPurchaseInvoice(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to verify purchase invoice: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// @Summary 支付采购发票
// @Description 支付采购发票
// @Tags 采购-采购发票
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购发票ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/invoices/{id}/pay [post]
func (h *PurchaseHandler) PayPurchaseInvoice(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.PayPurchaseInvoice(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to pay purchase invoice: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// 采购退货管理路由处理函数
// @Summary 获取采购退货列表
// @Description 获取所有采购退货的列表
// @Tags 采购-采购退货
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/returns [get]
func (h *PurchaseHandler) GetPurchaseReturnList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	returns, err := h.purchaseService.GetPurchaseReturnList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get purchase return list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    returns,
	})
}

// @Summary 获取采购退货详情
// @Description 根据ID获取采购退货详情
// @Tags 采购-采购退货
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购退货ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/returns/{id} [get]
func (h *PurchaseHandler) GetPurchaseReturnDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	return_, err := h.purchaseService.GetPurchaseReturnDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get purchase return detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    return_,
	})
}

// @Summary 创建采购退货
// @Description 创建新采购退货
// @Tags 采购-采购退货
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param return body map[string]interface{} true "采购退货信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/returns [post]
func (h *PurchaseHandler) CreatePurchaseReturn(c *gin.Context) {
	// 解析请求体
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	return_, err := h.purchaseService.CreatePurchaseReturn(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create purchase return: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    return_,
	})
}

// @Summary 更新采购退货
// @Description 根据ID更新采购退货
// @Tags 采购-采购退货
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购退货ID"
// @Param return body map[string]interface{} true "采购退货信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/returns/{id} [put]
func (h *PurchaseHandler) UpdatePurchaseReturn(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 解析请求体
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	return_, err := h.purchaseService.UpdatePurchaseReturn(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update purchase return: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    return_,
	})
}

// @Summary 删除采购退货
// @Description 根据ID删除采购退货
// @Tags 采购-采购退货
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购退货ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/returns/{id} [delete]
func (h *PurchaseHandler) DeletePurchaseReturn(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.DeletePurchaseReturn(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete purchase return: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// @Summary 完成采购退货
// @Description 完成采购退货
// @Tags 采购-采购退货
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "采购退货ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/returns/{id}/complete [post]
func (h *PurchaseHandler) CompletePurchaseReturn(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.purchaseService.CompletePurchaseReturn(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to complete purchase return: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

// 采购报表管理路由处理函数
// @Summary 获取采购汇总报表
// @Description 获取采购汇总的报表
// @Tags 采购-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/reports/summary [get]
func (h *PurchaseHandler) GetPurchaseSummaryReport(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	report, err := h.purchaseService.GetPurchaseSummaryReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get purchase summary report: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    report,
	})
}

// @Summary 获取采购明细报表
// @Description 获取采购明细的报表
// @Tags 采购-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/reports/detail [get]
func (h *PurchaseHandler) GetPurchaseDetailReport(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	report, err := h.purchaseService.GetPurchaseDetailReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get purchase detail report: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    report,
	})
}

// @Summary 获取供应商分析报表
// @Description 获取供应商分析的报表
// @Tags 采购-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/reports/supplier-analysis [get]
func (h *PurchaseHandler) GetSupplierAnalysisReport(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	report, err := h.purchaseService.GetSupplierAnalysisReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get supplier analysis report: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    report,
	})
}

// @Summary 获取价格分析报表
// @Description 获取价格分析的报表
// @Tags 采购-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/reports/price-analysis [get]
func (h *PurchaseHandler) GetPriceAnalysisReport(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	report, err := h.purchaseService.GetPriceAnalysisReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get price analysis report: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    report,
	})
}

// @Summary 获取采购预测报表
// @Description 获取采购预测的报表
// @Tags 采购-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/reports/forecast [get]
func (h *PurchaseHandler) GetPurchaseForecastReport(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	report, err := h.purchaseService.GetPurchaseForecastReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get purchase forecast report: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    report,
	})
}

// @Summary 导出采购报表
// @Description 导出采购相关的报表
// @Tags 采购-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/purchase/reports/export [get]
func (h *PurchaseHandler) ExportPurchaseReport(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	data, err := h.purchaseService.ExportPurchaseReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to export purchase report: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    string(data),
	})
}
