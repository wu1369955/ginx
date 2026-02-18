package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config 应用配置结构
type Config struct {
	Server   ServerConfig
	JWT      JWTConfig
	Database DatabaseConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port         string
	Environment  string
	ReadTimeout  int
	WriteTimeout int
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret      string
	ExpireHours int
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host      string
	Port      string
	User      string
	Password  string
	DBName    string
	Charset   string
	ParseTime bool
	Loc       string
}

// AppConfig 全局配置实例
var AppConfig Config

// LoadConfig 加载配置
func LoadConfig(configPath string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if configPath != "" {
		viper.AddConfigPath(configPath)
	} else {
		viper.AddConfigPath("./")
		viper.AddConfigPath("./configs")
		viper.AddConfigPath("../configs")
		viper.AddConfigPath("../../configs")
	}

	viper.AutomaticEnv()

	// 设置默认值
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.environment", "development")
	viper.SetDefault("server.readTimeout", 10)
	viper.SetDefault("server.writeTimeout", 10)

	viper.SetDefault("jwt.secret", "your-secret-key")
	viper.SetDefault("jwt.expireHours", 24)

	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "3306")
	viper.SetDefault("database.user", "root")
	viper.SetDefault("database.password", "root")
	viper.SetDefault("database.dbname", "app")
	viper.SetDefault("database.charset", "utf8mb4")
	viper.SetDefault("database.parseTime", true)
	viper.SetDefault("database.loc", "Local")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("配置文件未找到，使用默认配置")
		} else {
			return fmt.Errorf("读取配置文件错误: %w", err)
		}
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		return fmt.Errorf("解析配置错误: %w", err)
	}

	return nil
}

// GetConfig 获取配置实例
func GetConfig() Config {
	return AppConfig
}
