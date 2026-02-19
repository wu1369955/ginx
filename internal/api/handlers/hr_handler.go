package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/services"
)

// HRHandler 人力资源模块处理器
type HRHandler struct {
	hrService services.HRService
}

// NewHRHandler 创建人力资源模块处理器实例
func NewHRHandler(hrService services.HRService) *HRHandler {
	return &HRHandler{
		hrService: hrService,
	}
}

// 员工管理

// @Summary 获取员工列表
// @Description 获取所有员工的列表
// @Tags 人力资源-员工管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/employees [get]
func (h *HRHandler) GetEmployeeList(c *gin.Context) {
	// 解析查询参数
	req := make(map[string]interface{})
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	// 调用服务
	employees, err := h.hrService.GetEmployeeList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取员工列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  employees,
		"total": len(employees),
	})
}

// @Summary 获取员工详情
// @Description 根据ID获取员工的详细信息
// @Tags 人力资源-员工管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "员工ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/employees/{id} [get]
func (h *HRHandler) GetEmployeeDetail(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	employee, err := h.hrService.GetEmployeeDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取员工详情失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// @Summary 创建员工
// @Description 创建新的员工
// @Tags 人力资源-员工管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param employee body map[string]interface{} true "员工信息"
// @Success 201 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/employees [post]
func (h *HRHandler) CreateEmployee(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	employee, err := h.hrService.CreateEmployee(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建员工失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": employee})
}

// @Summary 更新员工
// @Description 根据ID更新员工信息
// @Tags 人力资源-员工管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "员工ID"
// @Param employee body map[string]interface{} true "员工信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/employees/{id} [put]
func (h *HRHandler) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	employee, err := h.hrService.UpdateEmployee(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新员工失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// @Summary 删除员工
// @Description 根据ID删除员工
// @Tags 人力资源-员工管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "员工ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/employees/{id} [delete]
func (h *HRHandler) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	if err := h.hrService.DeleteEmployee(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除员工失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除员工成功"})
}

// 部门管理

// @Summary 获取部门列表
// @Description 获取所有部门的列表
// @Tags 人力资源-部门管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/departments [get]
func (h *HRHandler) GetDepartmentList(c *gin.Context) {
	// 解析查询参数
	req := make(map[string]interface{})
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	// 调用服务
	departments, err := h.hrService.GetDepartmentList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取部门列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  departments,
		"total": len(departments),
	})
}

// @Summary 获取部门详情
// @Description 根据ID获取部门的详细信息
// @Tags 人力资源-部门管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "部门ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/departments/{id} [get]
func (h *HRHandler) GetDepartmentDetail(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	department, err := h.hrService.GetDepartmentDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取部门详情失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": department})
}

// @Summary 创建部门
// @Description 创建新的部门
// @Tags 人力资源-部门管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param department body map[string]interface{} true "部门信息"
// @Success 201 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/departments [post]
func (h *HRHandler) CreateDepartment(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	department, err := h.hrService.CreateDepartment(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建部门失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": department})
}

// @Summary 更新部门
// @Description 根据ID更新部门信息
// @Tags 人力资源-部门管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "部门ID"
// @Param department body map[string]interface{} true "部门信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/departments/{id} [put]
func (h *HRHandler) UpdateDepartment(c *gin.Context) {
	id := c.Param("id")

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	department, err := h.hrService.UpdateDepartment(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新部门失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": department})
}

// @Summary 删除部门
// @Description 根据ID删除部门
// @Tags 人力资源-部门管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "部门ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/departments/{id} [delete]
func (h *HRHandler) DeleteDepartment(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	if err := h.hrService.DeleteDepartment(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除部门失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除部门成功"})
}

// 职位管理

// @Summary 获取职位列表
// @Description 获取所有职位的列表
// @Tags 人力资源-职位管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/positions [get]
func (h *HRHandler) GetPositionList(c *gin.Context) {
	// 解析查询参数
	req := make(map[string]interface{})
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	// 调用服务
	positions, err := h.hrService.GetPositionList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取职位列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  positions,
		"total": len(positions),
	})
}

// @Summary 获取职位详情
// @Description 根据ID获取职位的详细信息
// @Tags 人力资源-职位管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "职位ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/positions/{id} [get]
func (h *HRHandler) GetPositionDetail(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	position, err := h.hrService.GetPositionDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取职位详情失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": position})
}

// @Summary 创建职位
// @Description 创建新的职位
// @Tags 人力资源-职位管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param position body map[string]interface{} true "职位信息"
// @Success 201 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/positions [post]
func (h *HRHandler) CreatePosition(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	position, err := h.hrService.CreatePosition(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建职位失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": position})
}

// @Summary 更新职位
// @Description 根据ID更新职位信息
// @Tags 人力资源-职位管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "职位ID"
// @Param position body map[string]interface{} true "职位信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/positions/{id} [put]
func (h *HRHandler) UpdatePosition(c *gin.Context) {
	id := c.Param("id")

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	position, err := h.hrService.UpdatePosition(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新职位失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": position})
}

// @Summary 删除职位
// @Description 根据ID删除职位
// @Tags 人力资源-职位管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "职位ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/positions/{id} [delete]
func (h *HRHandler) DeletePosition(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	if err := h.hrService.DeletePosition(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除职位失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除职位成功"})
}

// 考勤管理

// @Summary 获取考勤列表
// @Description 获取所有考勤记录的列表
// @Tags 人力资源-考勤管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/attendance [get]
func (h *HRHandler) GetAttendanceList(c *gin.Context) {
	// 解析查询参数
	req := make(map[string]interface{})
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	// 调用服务
	attendances, err := h.hrService.GetAttendanceList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取考勤列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  attendances,
		"total": len(attendances),
	})
}

// @Summary 获取考勤详情
// @Description 根据ID获取考勤记录的详细信息
// @Tags 人力资源-考勤管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "考勤记录ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/attendance/{id} [get]
func (h *HRHandler) GetAttendanceDetail(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	attendance, err := h.hrService.GetAttendanceDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取考勤详情失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": attendance})
}

// @Summary 创建考勤记录
// @Description 创建新的考勤记录
// @Tags 人力资源-考勤管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param attendance body map[string]interface{} true "考勤记录信息"
// @Success 201 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/attendance [post]
func (h *HRHandler) CreateAttendance(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	attendance, err := h.hrService.CreateAttendance(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建考勤记录失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": attendance})
}

// @Summary 更新考勤记录
// @Description 根据ID更新考勤记录信息
// @Tags 人力资源-考勤管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "考勤记录ID"
// @Param attendance body map[string]interface{} true "考勤记录信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/attendance/{id} [put]
func (h *HRHandler) UpdateAttendance(c *gin.Context) {
	id := c.Param("id")

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	attendance, err := h.hrService.UpdateAttendance(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新考勤记录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": attendance})
}

// @Summary 删除考勤记录
// @Description 根据ID删除考勤记录
// @Tags 人力资源-考勤管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "考勤记录ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/attendance/{id} [delete]
func (h *HRHandler) DeleteAttendance(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	if err := h.hrService.DeleteAttendance(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除考勤记录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除考勤记录成功"})
}

// 薪资管理

// @Summary 获取薪资列表
// @Description 获取所有薪资记录的列表
// @Tags 人力资源-薪资管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/salary [get]
func (h *HRHandler) GetSalaryList(c *gin.Context) {
	// 解析查询参数
	req := make(map[string]interface{})
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	// 调用服务
	salaries, err := h.hrService.GetSalaryList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取薪资列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  salaries,
		"total": len(salaries),
	})
}

// @Summary 获取薪资详情
// @Description 根据ID获取薪资记录的详细信息
// @Tags 人力资源-薪资管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "薪资记录ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/salary/{id} [get]
func (h *HRHandler) GetSalaryDetail(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	salary, err := h.hrService.GetSalaryDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取薪资详情失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": salary})
}

// @Summary 创建薪资记录
// @Description 创建新的薪资记录
// @Tags 人力资源-薪资管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param salary body map[string]interface{} true "薪资记录信息"
// @Success 201 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/salary [post]
func (h *HRHandler) CreateSalary(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	salary, err := h.hrService.CreateSalary(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建薪资记录失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": salary})
}

// @Summary 更新薪资记录
// @Description 根据ID更新薪资记录信息
// @Tags 人力资源-薪资管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "薪资记录ID"
// @Param salary body map[string]interface{} true "薪资记录信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/salary/{id} [put]
func (h *HRHandler) UpdateSalary(c *gin.Context) {
	id := c.Param("id")

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	salary, err := h.hrService.UpdateSalary(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新薪资记录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": salary})
}

// @Summary 删除薪资记录
// @Description 根据ID删除薪资记录
// @Tags 人力资源-薪资管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "薪资记录ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/salary/{id} [delete]
func (h *HRHandler) DeleteSalary(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	if err := h.hrService.DeleteSalary(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除薪资记录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除薪资记录成功"})
}

// 培训管理

// @Summary 获取培训列表
// @Description 获取所有培训记录的列表
// @Tags 人力资源-培训管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/training [get]
func (h *HRHandler) GetTrainingList(c *gin.Context) {
	// 解析查询参数
	req := make(map[string]interface{})
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	// 调用服务
	trainings, err := h.hrService.GetTrainingList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取培训列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  trainings,
		"total": len(trainings),
	})
}

// @Summary 获取培训详情
// @Description 根据ID获取培训记录的详细信息
// @Tags 人力资源-培训管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "培训记录ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/training/{id} [get]
func (h *HRHandler) GetTrainingDetail(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	training, err := h.hrService.GetTrainingDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取培训详情失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": training})
}

// @Summary 创建培训记录
// @Description 创建新的培训记录
// @Tags 人力资源-培训管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param training body map[string]interface{} true "培训记录信息"
// @Success 201 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/training [post]
func (h *HRHandler) CreateTraining(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	training, err := h.hrService.CreateTraining(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建培训记录失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": training})
}

// @Summary 更新培训记录
// @Description 根据ID更新培训记录信息
// @Tags 人力资源-培训管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "培训记录ID"
// @Param training body map[string]interface{} true "培训记录信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/training/{id} [put]
func (h *HRHandler) UpdateTraining(c *gin.Context) {
	id := c.Param("id")

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	training, err := h.hrService.UpdateTraining(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新培训记录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": training})
}

// @Summary 删除培训记录
// @Description 根据ID删除培训记录
// @Tags 人力资源-培训管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "培训记录ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/training/{id} [delete]
func (h *HRHandler) DeleteTraining(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	if err := h.hrService.DeleteTraining(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除培训记录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除培训记录成功"})
}

// 招聘管理

// @Summary 获取招聘列表
// @Description 获取所有招聘记录的列表
// @Tags 人力资源-招聘管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/recruitment [get]
func (h *HRHandler) GetRecruitmentList(c *gin.Context) {
	// 解析查询参数
	req := make(map[string]interface{})
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	// 调用服务
	recruitments, err := h.hrService.GetRecruitmentList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取招聘列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  recruitments,
		"total": len(recruitments),
	})
}

// @Summary 获取招聘详情
// @Description 根据ID获取招聘记录的详细信息
// @Tags 人力资源-招聘管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "招聘记录ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/recruitment/{id} [get]
func (h *HRHandler) GetRecruitmentDetail(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	recruitment, err := h.hrService.GetRecruitmentDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取招聘详情失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recruitment})
}

// @Summary 创建招聘记录
// @Description 创建新的招聘记录
// @Tags 人力资源-招聘管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param recruitment body map[string]interface{} true "招聘记录信息"
// @Success 201 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/recruitment [post]
func (h *HRHandler) CreateRecruitment(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	recruitment, err := h.hrService.CreateRecruitment(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建招聘记录失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": recruitment})
}

// @Summary 更新招聘记录
// @Description 根据ID更新招聘记录信息
// @Tags 人力资源-招聘管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "招聘记录ID"
// @Param recruitment body map[string]interface{} true "招聘记录信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/recruitment/{id} [put]
func (h *HRHandler) UpdateRecruitment(c *gin.Context) {
	id := c.Param("id")

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	recruitment, err := h.hrService.UpdateRecruitment(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新招聘记录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recruitment})
}

// @Summary 删除招聘记录
// @Description 根据ID删除招聘记录
// @Tags 人力资源-招聘管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "招聘记录ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/recruitment/{id} [delete]
func (h *HRHandler) DeleteRecruitment(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	if err := h.hrService.DeleteRecruitment(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除招聘记录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除招聘记录成功"})
}

// 绩效评估管理

// @Summary 获取绩效评估列表
// @Description 获取所有绩效评估记录的列表
// @Tags 人力资源-绩效评估管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/performance [get]
func (h *HRHandler) GetPerformanceList(c *gin.Context) {
	// 解析查询参数
	req := make(map[string]interface{})
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	// 调用服务
	performances, err := h.hrService.GetPerformanceList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取绩效评估列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  performances,
		"total": len(performances),
	})
}

// @Summary 获取绩效评估详情
// @Description 根据ID获取绩效评估记录的详细信息
// @Tags 人力资源-绩效评估管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "绩效评估记录ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/performance/{id} [get]
func (h *HRHandler) GetPerformanceDetail(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	performance, err := h.hrService.GetPerformanceDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取绩效评估详情失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": performance})
}

// @Summary 创建绩效评估记录
// @Description 创建新的绩效评估记录
// @Tags 人力资源-绩效评估管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param performance body map[string]interface{} true "绩效评估记录信息"
// @Success 201 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/performance [post]
func (h *HRHandler) CreatePerformance(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	performance, err := h.hrService.CreatePerformance(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建绩效评估记录失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": performance})
}

// @Summary 更新绩效评估记录
// @Description 根据ID更新绩效评估记录信息
// @Tags 人力资源-绩效评估管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "绩效评估记录ID"
// @Param performance body map[string]interface{} true "绩效评估记录信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/performance/{id} [put]
func (h *HRHandler) UpdatePerformance(c *gin.Context) {
	id := c.Param("id")

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务
	performance, err := h.hrService.UpdatePerformance(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新绩效评估记录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": performance})
}

// @Summary 删除绩效评估记录
// @Description 根据ID删除绩效评估记录
// @Tags 人力资源-绩效评估管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "绩效评估记录ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/v1/hr/performance/{id} [delete]
func (h *HRHandler) DeletePerformance(c *gin.Context) {
	id := c.Param("id")

	// 调用服务
	if err := h.hrService.DeletePerformance(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除绩效评估记录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除绩效评估记录成功"})
}
