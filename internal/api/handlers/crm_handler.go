package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/api/schemas"
	"github.com/wu136995/ginx/internal/services"
)

// CRMHandler CRM处理器
type CRMHandler struct {
	crmService services.CRMService
}

// NewCRMHandler 创建CRM处理器实例
func NewCRMHandler(crmService services.CRMService) *CRMHandler {
	return &CRMHandler{
		crmService: crmService,
	}
}

// 客户管理路由处理函数
// @Summary 获取客户列表
// @Description 获取所有客户的列表
// @Tags CRM-客户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/customers [get]
func (h *CRMHandler) GetCustomerList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	customers, err := h.crmService.GetCustomerList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get customer list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    customers,
	})
}

// @Summary 获取客户详情
// @Description 根据ID获取客户详情
// @Tags CRM-客户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "客户ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/customers/{id} [get]
func (h *CRMHandler) GetCustomerDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	customer, err := h.crmService.GetCustomerDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get customer detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    customer,
	})
}

// @Summary 创建客户
// @Description 创建新客户
// @Tags CRM-客户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param customer body schemas.CustomerCreateRequest true "客户信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/customers [post]
func (h *CRMHandler) CreateCustomer(c *gin.Context) {
	// 解析请求体
	var req schemas.CustomerCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	customer, err := h.crmService.CreateCustomer(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create customer: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    customer,
	})
}

// @Summary 更新客户
// @Description 根据ID更新客户信息
// @Tags CRM-客户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "客户ID"
// @Param customer body schemas.CustomerUpdateRequest true "客户信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/customers/{id} [put]
func (h *CRMHandler) UpdateCustomer(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 解析请求体
	var req schemas.CustomerUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	customer, err := h.crmService.UpdateCustomer(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update customer: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    customer,
	})
}

// @Summary 删除客户
// @Description 根据ID删除客户
// @Tags CRM-客户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "客户ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/customers/{id} [delete]
func (h *CRMHandler) DeleteCustomer(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.crmService.DeleteCustomer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete customer: " + err.Error(),
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

// 销售线索管理路由处理函数
// @Summary 获取销售线索列表
// @Description 获取所有销售线索的列表
// @Tags CRM-销售线索管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/leads [get]
func (h *CRMHandler) GetLeadList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	leads, err := h.crmService.GetLeadList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get lead list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    leads,
	})
}

// @Summary 获取销售线索详情
// @Description 根据ID获取销售线索详情
// @Tags CRM-销售线索管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "销售线索ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/leads/{id} [get]
func (h *CRMHandler) GetLeadDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	lead, err := h.crmService.GetLeadDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get lead detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    lead,
	})
}

// @Summary 创建销售线索
// @Description 创建新销售线索
// @Tags CRM-销售线索管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param lead body schemas.LeadCreateRequest true "销售线索信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/leads [post]
func (h *CRMHandler) CreateLead(c *gin.Context) {
	// 解析请求体
	var req schemas.LeadCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	lead, err := h.crmService.CreateLead(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create lead: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    lead,
	})
}

// @Summary 更新销售线索
// @Description 根据ID更新销售线索
// @Tags CRM-销售线索管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "销售线索ID"
// @Param lead body schemas.LeadUpdateRequest true "销售线索信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/leads/{id} [put]
func (h *CRMHandler) UpdateLead(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 解析请求体
	var req schemas.LeadUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	lead, err := h.crmService.UpdateLead(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update lead: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    lead,
	})
}

// @Summary 删除销售线索
// @Description 根据ID删除销售线索
// @Tags CRM-销售线索管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "销售线索ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/leads/{id} [delete]
func (h *CRMHandler) DeleteLead(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.crmService.DeleteLead(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete lead: " + err.Error(),
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

// @Summary 筛选销售线索
// @Description 筛选销售线索
// @Tags CRM-销售线索管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "销售线索ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/leads/{id}/qualify [post]
func (h *CRMHandler) QualifyLead(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.crmService.QualifyLead(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to qualify lead: " + err.Error(),
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

// @Summary 转换销售线索
// @Description 转换销售线索为客户
// @Tags CRM-销售线索管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "销售线索ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/leads/{id}/convert [post]
func (h *CRMHandler) ConvertLead(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	customer, err := h.crmService.ConvertLead(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to convert lead: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    customer,
	})
}

// 销售机会管理路由处理函数
// @Summary 获取销售机会列表
// @Description 获取所有销售机会的列表
// @Tags CRM-销售机会管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/opportunities [get]
func (h *CRMHandler) GetOpportunityList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	opportunities, err := h.crmService.GetOpportunityList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get opportunity list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    opportunities,
	})
}

// @Summary 获取销售机会详情
// @Description 根据ID获取销售机会详情
// @Tags CRM-销售机会管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "销售机会ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/opportunities/{id} [get]
func (h *CRMHandler) GetOpportunityDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	opportunity, err := h.crmService.GetOpportunityDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get opportunity detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    opportunity,
	})
}

// @Summary 创建销售机会
// @Description 创建新销售机会
// @Tags CRM-销售机会管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param opportunity body schemas.OpportunityCreateRequest true "销售机会信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/opportunities [post]
func (h *CRMHandler) CreateOpportunity(c *gin.Context) {
	// 解析请求体
	var req schemas.OpportunityCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	opportunity, err := h.crmService.CreateOpportunity(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create opportunity: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    opportunity,
	})
}

// @Summary 更新销售机会
// @Description 根据ID更新销售机会
// @Tags CRM-销售机会管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "销售机会ID"
// @Param opportunity body schemas.OpportunityUpdateRequest true "销售机会信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/opportunities/{id} [put]
func (h *CRMHandler) UpdateOpportunity(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 解析请求体
	var req schemas.OpportunityUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	opportunity, err := h.crmService.UpdateOpportunity(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update opportunity: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    opportunity,
	})
}

// @Summary 删除销售机会
// @Description 根据ID删除销售机会
// @Tags CRM-销售机会管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "销售机会ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/opportunities/{id} [delete]
func (h *CRMHandler) DeleteOpportunity(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.crmService.DeleteOpportunity(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete opportunity: " + err.Error(),
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

// @Summary 关闭销售机会
// @Description 关闭销售机会
// @Tags CRM-销售机会管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "销售机会ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/opportunities/{id}/close [post]
func (h *CRMHandler) CloseOpportunity(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.crmService.CloseOpportunity(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to close opportunity: " + err.Error(),
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

// 营销活动管理路由处理函数
// @Summary 获取营销活动列表
// @Description 获取所有营销活动的列表
// @Tags CRM-营销活动管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/campaigns [get]
func (h *CRMHandler) GetCampaignList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	campaigns, err := h.crmService.GetCampaignList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get campaign list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    campaigns,
	})
}

// @Summary 获取营销活动详情
// @Description 根据ID获取营销活动详情
// @Tags CRM-营销活动管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "营销活动ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/campaigns/{id} [get]
func (h *CRMHandler) GetCampaignDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	campaign, err := h.crmService.GetCampaignDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get campaign detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    campaign,
	})
}

// @Summary 创建营销活动
// @Description 创建新营销活动
// @Tags CRM-营销活动管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param campaign body schemas.CampaignCreateRequest true "营销活动信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/campaigns [post]
func (h *CRMHandler) CreateCampaign(c *gin.Context) {
	// 解析请求体
	var req schemas.CampaignCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	campaign, err := h.crmService.CreateCampaign(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create campaign: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    campaign,
	})
}

// @Summary 更新营销活动
// @Description 根据ID更新营销活动
// @Tags CRM-营销活动管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "营销活动ID"
// @Param campaign body schemas.CampaignUpdateRequest true "营销活动信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/campaigns/{id} [put]
func (h *CRMHandler) UpdateCampaign(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 解析请求体
	var req schemas.CampaignUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	campaign, err := h.crmService.UpdateCampaign(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update campaign: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    campaign,
	})
}

// @Summary 删除营销活动
// @Description 根据ID删除营销活动
// @Tags CRM-营销活动管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "营销活动ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/campaigns/{id} [delete]
func (h *CRMHandler) DeleteCampaign(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.crmService.DeleteCampaign(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete campaign: " + err.Error(),
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

// 客户活动管理路由处理函数
// @Summary 获取客户活动列表
// @Description 获取所有客户活动的列表
// @Tags CRM-客户活动管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/activities [get]
func (h *CRMHandler) GetActivityList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	activities, err := h.crmService.GetActivityList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get activity list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    activities,
	})
}

// @Summary 获取客户活动详情
// @Description 根据ID获取客户活动详情
// @Tags CRM-客户活动管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "客户活动ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/activities/{id} [get]
func (h *CRMHandler) GetActivityDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	activity, err := h.crmService.GetActivityDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get activity detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    activity,
	})
}

// @Summary 创建客户活动
// @Description 创建新客户活动
// @Tags CRM-客户活动管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param activity body schemas.ActivityCreateRequest true "客户活动信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/activities [post]
func (h *CRMHandler) CreateActivity(c *gin.Context) {
	// 解析请求体
	var req schemas.ActivityCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	activity, err := h.crmService.CreateActivity(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create activity: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    activity,
	})
}

// @Summary 更新客户活动
// @Description 根据ID更新客户活动
// @Tags CRM-客户活动管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "客户活动ID"
// @Param activity body schemas.ActivityUpdateRequest true "客户活动信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/activities/{id} [put]
func (h *CRMHandler) UpdateActivity(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 解析请求体
	var req schemas.ActivityUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	activity, err := h.crmService.UpdateActivity(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update activity: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    activity,
	})
}

// @Summary 删除客户活动
// @Description 根据ID删除客户活动
// @Tags CRM-客户活动管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "客户活动ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/activities/{id} [delete]
func (h *CRMHandler) DeleteActivity(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.crmService.DeleteActivity(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete activity: " + err.Error(),
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

// 客户服务管理路由处理函数
// @Summary 获取客户服务请求列表
// @Description 获取所有客户服务请求的列表
// @Tags CRM-客户服务管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/service-requests [get]
func (h *CRMHandler) GetServiceRequestList(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	serviceRequests, err := h.crmService.GetServiceRequestList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get service request list: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    serviceRequests,
	})
}

// @Summary 获取客户服务请求详情
// @Description 根据ID获取客户服务请求详情
// @Tags CRM-客户服务管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "客户服务请求ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/service-requests/{id} [get]
func (h *CRMHandler) GetServiceRequestDetail(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	serviceRequest, err := h.crmService.GetServiceRequestDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get service request detail: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    serviceRequest,
	})
}

// @Summary 创建客户服务请求
// @Description 创建新客户服务请求
// @Tags CRM-客户服务管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param serviceRequest body schemas.ServiceRequestCreateRequest true "客户服务请求信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/service-requests [post]
func (h *CRMHandler) CreateServiceRequest(c *gin.Context) {
	// 解析请求体
	var req schemas.ServiceRequestCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	serviceRequest, err := h.crmService.CreateServiceRequest(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create service request: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    serviceRequest,
	})
}

// @Summary 更新客户服务请求
// @Description 根据ID更新客户服务请求
// @Tags CRM-客户服务管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "客户服务请求ID"
// @Param serviceRequest body schemas.ServiceRequestUpdateRequest true "客户服务请求信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/service-requests/{id} [put]
func (h *CRMHandler) UpdateServiceRequest(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 解析请求体
	var req schemas.ServiceRequestUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service方法
	serviceRequest, err := h.crmService.UpdateServiceRequest(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update service request: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    serviceRequest,
	})
}

// @Summary 删除客户服务请求
// @Description 根据ID删除客户服务请求
// @Tags CRM-客户服务管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "客户服务请求ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/service-requests/{id} [delete]
func (h *CRMHandler) DeleteServiceRequest(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.crmService.DeleteServiceRequest(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete service request: " + err.Error(),
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

// @Summary 解决客户服务请求
// @Description 解决客户服务请求
// @Tags CRM-客户服务管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "客户服务请求ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/service-requests/{id}/resolve [post]
func (h *CRMHandler) ResolveServiceRequest(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.crmService.ResolveServiceRequest(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to resolve service request: " + err.Error(),
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

// CRM报表路由处理函数
// @Summary 获取客户报表
// @Description 获取客户的报表
// @Tags CRM-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/reports/customers [get]
func (h *CRMHandler) GetCustomerReport(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	report, err := h.crmService.GetCustomerReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get customer report: " + err.Error(),
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

// @Summary 获取销售漏斗报表
// @Description 获取销售漏斗的报表
// @Tags CRM-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/reports/sales-pipeline [get]
func (h *CRMHandler) GetSalesPipelineReport(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	report, err := h.crmService.GetSalesPipelineReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get sales pipeline report: " + err.Error(),
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

// @Summary 获取营销活动效果报表
// @Description 获取营销活动效果的报表
// @Tags CRM-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/reports/campaign-effectiveness [get]
func (h *CRMHandler) GetCampaignEffectivenessReport(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	report, err := h.crmService.GetCampaignEffectivenessReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get campaign effectiveness report: " + err.Error(),
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

// @Summary 获取客户服务请求报表
// @Description 获取客户服务请求的报表
// @Tags CRM-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/reports/service-requests [get]
func (h *CRMHandler) GetServiceRequestReport(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	report, err := h.crmService.GetServiceRequestReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get service request report: " + err.Error(),
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

// @Summary 导出CRM报表
// @Description 导出CRM相关的报表
// @Tags CRM-报表管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/crm/reports/export [get]
func (h *CRMHandler) ExportCRMReport(c *gin.Context) {
	// 从查询参数获取过滤条件
	req := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		req[k] = v[0]
	}

	// 调用service方法
	data, err := h.crmService.ExportCRMReport(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to export CRM report: " + err.Error(),
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
