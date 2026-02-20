package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/services"
)

// FinanceHandler 财务处理器
type FinanceHandler struct {
	financeService services.FinanceService
}

// NewFinanceHandler 创建财务处理器实例
func NewFinanceHandler(financeService services.FinanceService) *FinanceHandler {
	return &FinanceHandler{
		financeService: financeService,
	}
}

// 账户管理路由处理函数
// @Summary 获取账户列表
// @Description 获取所有账户的列表
// @Tags 财务-账户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/accounts [get]
func (h *FinanceHandler) GetAccountList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	accounts, err := h.financeService.GetAccountList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get account list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    accounts,
	})
}

// @Summary 获取账户详情
// @Description 根据ID获取账户详情
// @Tags 财务-账户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "账户ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/accounts/{id} [get]
func (h *FinanceHandler) GetAccountDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	account, err := h.financeService.GetAccountDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get account detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    account,
	})
}

// @Summary 创建账户
// @Description 创建新账户
// @Tags 财务-账户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param account body map[string]interface{} true "账户信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/accounts [post]
func (h *FinanceHandler) CreateAccount(c *gin.Context) {
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
	account, err := h.financeService.CreateAccount(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create account: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    account,
	})
}

// @Summary 更新账户
// @Description 根据ID更新账户信息
// @Tags 财务-账户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "账户ID"
// @Param account body map[string]interface{} true "账户信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/accounts/{id} [put]
func (h *FinanceHandler) UpdateAccount(c *gin.Context) {
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
	account, err := h.financeService.UpdateAccount(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update account: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    account,
	})
}

// @Summary 删除账户
// @Description 根据ID删除账户
// @Tags 财务-账户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "账户ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/accounts/{id} [delete]
func (h *FinanceHandler) DeleteAccount(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.financeService.DeleteAccount(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete account: " + err.Error(),
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

// 凭证管理路由处理函数
// @Summary 获取凭证列表
// @Description 获取所有凭证的列表
// @Tags 财务-凭证管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/vouchers [get]
func (h *FinanceHandler) GetVoucherList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	vouchers, err := h.financeService.GetVoucherList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get voucher list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    vouchers,
	})
}

// @Summary 获取凭证详情
// @Description 根据ID获取凭证详情
// @Tags 财务-凭证管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "凭证ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/vouchers/{id} [get]
func (h *FinanceHandler) GetVoucherDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	voucher, err := h.financeService.GetVoucherDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get voucher detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    voucher,
	})
}

// @Summary 创建凭证
// @Description 创建新凭证
// @Tags 财务-凭证管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param voucher body map[string]interface{} true "凭证信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/vouchers [post]
func (h *FinanceHandler) CreateVoucher(c *gin.Context) {
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
	voucher, err := h.financeService.CreateVoucher(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create voucher: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    voucher,
	})
}

// @Summary 更新凭证
// @Description 根据ID更新凭证
// @Tags 财务-凭证管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "凭证ID"
// @Param voucher body map[string]interface{} true "凭证信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/vouchers/{id} [put]
func (h *FinanceHandler) UpdateVoucher(c *gin.Context) {
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
	voucher, err := h.financeService.UpdateVoucher(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update voucher: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    voucher,
	})
}

// @Summary 删除凭证
// @Description 根据ID删除凭证
// @Tags 财务-凭证管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "凭证ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/vouchers/{id} [delete]
func (h *FinanceHandler) DeleteVoucher(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.financeService.DeleteVoucher(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete voucher: " + err.Error(),
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

// @Summary 提交凭证
// @Description 提交凭证
// @Tags 财务-凭证管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "凭证ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/vouchers/{id}/submit [post]
func (h *FinanceHandler) SubmitVoucher(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.financeService.SubmitVoucher(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to submit voucher: " + err.Error(),
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

// @Summary 审批凭证
// @Description 审批凭证
// @Tags 财务-凭证管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "凭证ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/vouchers/{id}/approve [post]
func (h *FinanceHandler) ApproveVoucher(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.financeService.ApproveVoucher(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to approve voucher: " + err.Error(),
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

// @Summary 过账凭证
// @Description 过账凭证
// @Tags 财务-凭证管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "凭证ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/vouchers/{id}/post [post]
func (h *FinanceHandler) PostVoucher(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.financeService.SubmitVoucher(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to submit voucher: " + err.Error(),
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

// 预算管理路由处理函数
// @Summary 获取预算列表
// @Description 获取所有预算的列表
// @Tags 财务-预算管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/budgets [get]
func (h *FinanceHandler) GetBudgetList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	budgets, err := h.financeService.GetBudgetList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get budget list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    budgets,
	})
}

// @Summary 获取预算详情
// @Description 根据ID获取预算详情
// @Tags 财务-预算管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "预算ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/budgets/{id} [get]
func (h *FinanceHandler) GetBudgetDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	budget, err := h.financeService.GetBudgetDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get budget detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    budget,
	})
}

// @Summary 创建预算
// @Description 创建新预算
// @Tags 财务-预算管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param budget body map[string]interface{} true "预算信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/budgets [post]
func (h *FinanceHandler) CreateBudget(c *gin.Context) {
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
	budget, err := h.financeService.CreateBudget(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create budget: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    budget,
	})
}

// @Summary 更新预算
// @Description 根据ID更新预算
// @Tags 财务-预算管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "预算ID"
// @Param budget body map[string]interface{} true "预算信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/budgets/{id} [put]
func (h *FinanceHandler) UpdateBudget(c *gin.Context) {
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
	budget, err := h.financeService.UpdateBudget(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update budget: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    budget,
	})
}

// @Summary 删除预算
// @Description 根据ID删除预算
// @Tags 财务-预算管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "预算ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/budgets/{id} [delete]
func (h *FinanceHandler) DeleteBudget(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.financeService.DeleteBudget(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete budget: " + err.Error(),
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

// @Summary 提交预算
// @Description 提交预算
// @Tags 财务-预算管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "预算ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/budgets/{id}/submit [post]
func (h *FinanceHandler) SubmitBudget(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.financeService.SubmitBudget(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to submit budget: " + err.Error(),
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

// @Summary 审批预算
// @Description 审批预算
// @Tags 财务-预算管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "预算ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/budgets/{id}/approve [post]
func (h *FinanceHandler) ApproveBudget(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.financeService.ApproveBudget(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to approve budget: " + err.Error(),
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

// @Summary 调整预算
// @Description 调整预算
// @Tags 财务-预算管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "预算ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/budgets/{id}/adjust [post]
func (h *FinanceHandler) AdjustBudget(c *gin.Context) {
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
	budget, err := h.financeService.AdjustBudget(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to adjust budget: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    budget,
	})
}

// 固定资产管理路由处理函数
// @Summary 获取固定资产列表
// @Description 获取所有固定资产的列表
// @Tags 财务-固定资产管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/fixed-assets [get]
func (h *FinanceHandler) GetFixedAssetList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	assets, err := h.financeService.GetAssetList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get asset list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    assets,
	})
}

// @Summary 获取固定资产详情
// @Description 根据ID获取固定资产详情
// @Tags 财务-固定资产管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "固定资产ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/fixed-assets/{id} [get]
func (h *FinanceHandler) GetFixedAssetDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	asset, err := h.financeService.GetAssetDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get asset detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    asset,
	})
}

// @Summary 创建固定资产
// @Description 创建新固定资产
// @Tags 财务-固定资产管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param asset body map[string]interface{} true "固定资产信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/fixed-assets [post]
func (h *FinanceHandler) CreateFixedAsset(c *gin.Context) {
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
	asset, err := h.financeService.CreateAsset(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create asset: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    asset,
	})
}

// @Summary 更新固定资产
// @Description 根据ID更新固定资产
// @Tags 财务-固定资产管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "固定资产ID"
// @Param asset body map[string]interface{} true "固定资产信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/fixed-assets/{id} [put]
func (h *FinanceHandler) UpdateFixedAsset(c *gin.Context) {
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
	asset, err := h.financeService.UpdateAsset(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update asset: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    asset,
	})
}

// @Summary 删除固定资产
// @Description 根据ID删除固定资产
// @Tags 财务-固定资产管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "固定资产ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/fixed-assets/{id} [delete]
func (h *FinanceHandler) DeleteFixedAsset(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.financeService.DeleteAsset(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete asset: " + err.Error(),
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

// @Summary 折旧固定资产
// @Description 折旧固定资产
// @Tags 财务-固定资产管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "固定资产ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/fixed-assets/{id}/depreciate [post]
func (h *FinanceHandler) DepreciateFixedAsset(c *gin.Context) {
	// 调用service方法
	_, err := h.financeService.CalculateDepreciation(map[string]interface{}{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to calculate depreciation: " + err.Error(),
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

// @Summary 处置固定资产
// @Description 处置固定资产
// @Tags 财务-固定资产管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "固定资产ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/fixed-assets/{id}/dispose [post]
func (h *FinanceHandler) DisposeFixedAsset(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.financeService.DisposeAsset(id, map[string]interface{}{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to dispose asset: " + err.Error(),
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

// 财务报表路由处理函数
// @Summary 获取资产负债表
// @Description 获取资产负债表
// @Tags 财务-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/reports/balance-sheet [get]
func (h *FinanceHandler) GetBalanceSheet(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	report, err := h.financeService.GetBalanceSheet(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get balance sheet: " + err.Error(),
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

// @Summary 获取利润表
// @Description 获取利润表
// @Tags 财务-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/reports/income-statement [get]
func (h *FinanceHandler) GetIncomeStatement(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	report, err := h.financeService.GetIncomeStatement(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get income statement: " + err.Error(),
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

// @Summary 获取现金流量表
// @Description 获取现金流量表
// @Tags 财务-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/reports/cash-flow [get]
func (h *FinanceHandler) GetCashFlowStatement(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	report, err := h.financeService.GetCashFlowStatement(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get cash flow statement: " + err.Error(),
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

// @Summary 获取试算平衡表
// @Description 获取试算平衡表
// @Tags 财务-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/reports/trial-balance [get]
func (h *FinanceHandler) GetTrialBalance(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	report, err := h.financeService.GetTrialBalance(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get trial balance: " + err.Error(),
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

// @Summary 导出财务报表
// @Description 导出财务相关的报表
// @Tags 财务-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/finance/reports/export [get]
func (h *FinanceHandler) ExportFinancialReport(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	data, err := h.financeService.ExportFinanceReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to export finance report: " + err.Error(),
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
