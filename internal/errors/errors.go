package errors

import (
	"fmt"
	"net/http"
)

// AppError 应用错误类型
type AppError struct {
	Code    int    `json:"code"`    // 错误码
	Message string `json:"message"` // 错误信息
	Err     error  `json:"err"`     // 原始错误
}

// Error 实现error接口
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NewAppError 创建新的应用错误
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// BadRequest 创建400错误
func BadRequest(message string, err error) *AppError {
	return NewAppError(http.StatusBadRequest, message, err)
}

// Unauthorized 创建401错误
func Unauthorized(message string, err error) *AppError {
	return NewAppError(http.StatusUnauthorized, message, err)
}

// Forbidden 创建403错误
func Forbidden(message string, err error) *AppError {
	return NewAppError(http.StatusForbidden, message, err)
}

// NotFound 创建404错误
func NotFound(message string, err error) *AppError {
	return NewAppError(http.StatusNotFound, message, err)
}

// InternalServerError 创建500错误
func InternalServerError(message string, err error) *AppError {
	return NewAppError(http.StatusInternalServerError, message, err)
}
