package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 消息
	Data    interface{} `json:"data"`    // 数据
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

// Fail 失败响应
func Fail(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// BadRequest 400错误响应
func BadRequest(c *gin.Context, message string) {
	Fail(c, http.StatusBadRequest, message)
}

// Unauthorized 401错误响应
func Unauthorized(c *gin.Context, message string) {
	Fail(c, http.StatusUnauthorized, message)
}

// Forbidden 403错误响应
func Forbidden(c *gin.Context, message string) {
	Fail(c, http.StatusForbidden, message)
}

// NotFound 404错误响应
func NotFound(c *gin.Context, message string) {
	Fail(c, http.StatusNotFound, message)
}

// InternalServerError 500错误响应
func InternalServerError(c *gin.Context, message string) {
	Fail(c, http.StatusInternalServerError, message)
}
