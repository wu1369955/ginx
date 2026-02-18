package config

import (
	"testing"

	"github.com/wu136995/ginx/internal/config"
)

// TestLoadConfig 测试加载配置
func TestLoadConfig(t *testing.T) {
	err := config.LoadConfig("")
	if err != nil {
		t.Errorf("LoadConfig failed: %v", err)
	}

	// 验证默认配置
	cfg := config.GetConfig()
	if cfg.Server.Port != "8080" {
		t.Errorf("Expected server port 8080, got %s", cfg.Server.Port)
	}

	if cfg.Server.Environment != "development" {
		t.Errorf("Expected server environment development, got %s", cfg.Server.Environment)
	}

	if cfg.Database.Host != "localhost" {
		t.Errorf("Expected database host localhost, got %s", cfg.Database.Host)
	}

	if cfg.Database.Port != "5432" {
		t.Errorf("Expected database port 5432, got %s", cfg.Database.Port)
	}
}
