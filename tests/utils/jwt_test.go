package utils

import (
	"testing"

	"github.com/wu136995/ginx/internal/utils"
)

// TestGenerateToken 测试生成JWT token
func TestGenerateToken(t *testing.T) {
	userID := uint(1)
	username := "testuser"
	role := "user"
	secret := "testsecret"
	expireHours := 24

	token, err := utils.GenerateToken(userID, username, role, secret, expireHours)
	if err != nil {
		t.Errorf("GenerateToken failed: %v", err)
	}

	if token == "" {
		t.Error("GenerateToken returned empty string")
	}
}

// TestParseToken 测试解析JWT token
func TestParseToken(t *testing.T) {
	userID := uint(1)
	username := "testuser"
	role := "user"
	secret := "testsecret"
	expireHours := 24

	// 生成token
	token, err := utils.GenerateToken(userID, username, role, secret, expireHours)
	if err != nil {
		t.Errorf("GenerateToken failed: %v", err)
	}

	// 解析token
	claims, err := utils.ParseToken(token, secret)
	if err != nil {
		t.Errorf("ParseToken failed: %v", err)
	}

	// 验证token内容
	if claims.UserID != userID {
		t.Errorf("Expected userID %d, got %d", userID, claims.UserID)
	}

	if claims.Username != username {
		t.Errorf("Expected username %s, got %s", username, claims.Username)
	}

	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}
