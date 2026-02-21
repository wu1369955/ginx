package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthCheck 健康检查处理程序
func HealthCheck(c *gin.Context) {
	log.Println("健康检查请求")
	c.JSON(http.StatusOK, gin.H{
		"timestamp": time.Now().Format(time.RFC3339),
		"status":    "ok",
		"message":   "服务运行正常",
	})
}
