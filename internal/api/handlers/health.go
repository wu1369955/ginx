package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck 健康检查处理程序
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "服务运行正常",
	})
}
