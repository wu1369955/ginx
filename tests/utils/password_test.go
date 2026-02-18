package utils

import (
	"testing"

	"github.com/wu136995/ginx/internal/utils"
)

// TestHashPassword 测试密码加密
func TestHashPassword(t *testing.T) {
	password := "test123"
	hash, err := utils.HashPassword(password)
	if err != nil {
		t.Errorf("HashPassword failed: %v", err)
	}

	if hash == "" {
		t.Error("HashPassword returned empty string")
	}
}

// TestCheckPassword 测试密码验证
func TestCheckPassword(t *testing.T) {
	password := "test123"
	hash, err := utils.HashPassword(password)
	if err != nil {
		t.Errorf("HashPassword failed: %v", err)
	}

	// 测试正确的密码
	if !utils.CheckPassword(password, hash) {
		t.Error("CheckPassword failed for correct password")
	}

	// 测试错误的密码
	if utils.CheckPassword("wrongpassword", hash) {
		t.Error("CheckPassword passed for wrong password")
	}
}
