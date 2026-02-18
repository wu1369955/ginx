package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserInfo 获取用户信息处理程序
func GetUserInfo(c *gin.Context) {
	// 从上下文中获取用户信息
	userID, _ := c.Get("userID")
	userName, _ := c.Get("userName")

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": gin.H{
			"userID":   userID,
			"userName": userName,
		},
	})
}
