package config

import (
	"time"

	"github.com/spf13/viper"
)

// 应用配置
type AppConfig struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Data     DataConfig     `mapstructure:"data"` // 新增数据配置
}

// 服务器配置
type ServerConfig struct {
	Port string `mapstructure:"port"`
}

// 数据库配置
type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	Charset  string `mapstructure:"charset"`
}

// JWT配置
type JWTConfig struct {
	Secret      string        `mapstructure:"secret"`
	ExpireHours time.Duration `mapstructure:"expireHours"`
}

// 数据配置（新增）
type DataConfig struct {
	HotThreshold  int64         `mapstructure:"hotThreshold"`  // 热数据阈值（秒）
	ColdThreshold int64         `mapstructure:"coldThreshold"` // 冷数据阈值（秒）
	MigrateInterval time.Duration `mapstructure:"migrateInterval"` // 迁移检查间隔
	ColdStoragePath string       `mapstructure:"coldStoragePath"` // 冷存储路径
}

// 全局配置实例
var AppConfig AppConfig

// 加载配置
func LoadConfig() error {
	// 设置默认值
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("database.driver", "mysql")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "3306")
	viper.SetDefault("database.username", "root")
	viper.SetDefault("database.password", "123456")
	viper.SetDefault("database.dbname", "ginx")
	viper.SetDefault("database.charset", "utf8mb4")
	viper.SetDefault("jwt.secret", "your-secret-key")
	viper.SetDefault("jwt.expireHours", 24)
	// 新增数据配置默认值
	viper.SetDefault("data.hotThreshold", 3600)  // 1小时
	viper.SetDefault("data.coldThreshold", 86400) // 24小时
	viper.SetDefault("data.migrateInterval", 300) // 5分钟
	viper.SetDefault("data.coldStoragePath", "./cold_data")

	// 读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		// 配置文件不存在时使用默认值
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	// 解析配置
	if err := viper.Unmarshal(&AppConfig); err != nil {
		return err
	}

	return nil
}
