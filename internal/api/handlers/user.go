package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/models"
	"github.com/wu136995/ginx/internal/services"
)

// UserHandler 用户处理器
type UserHandler struct {
	userService services.UserService
}

// NewUserHandler 创建用户处理器实例
func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// @Summary 获取用户信息
// @Description 获取当前登录用户的详细信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/user/info [get]
// GetUserInfo 获取用户信息处理程序
func (h *UserHandler) GetUserInfo(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    nil,
		})
		return
	}

	// 调用service方法获取用户信息
	user, err := h.userService.GetUserByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get user info: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    user,
	})
}

// @Summary 创建用户
// @Description 创建新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body map[string]interface{} true "用户信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/user [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
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
	user := &models.User{
		Username: req["username"].(string),
		Email:    req["email"].(string),
		// 其他字段映射
	}

	err := h.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create user: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    user,
	})
}

// @Summary 更新用户
// @Description 更新用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body map[string]interface{} true "用户信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/user [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	// 检查是否已认证
	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    nil,
		})
		return
	}

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
	user := &models.User{
		Username: req["username"].(string),
		Email:    req["email"].(string),
		// 其他字段映射
	}

	err := h.userService.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update user: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    user,
	})
}

// @Summary 删除用户
// @Description 删除用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "用户ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Router /api/user/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 调用service方法
	err := h.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete user: " + err.Error(),
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
