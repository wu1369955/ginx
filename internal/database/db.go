package database

import (
	"fmt"
	"log"

	"github.com/wu136995/ginx/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 全局数据库连接实例
var DB *gorm.DB

// Connect 连接数据库
func Connect() error {
	cfg := config.GetConfig()
	dbConfig := cfg.Database

	// 构建DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName, dbConfig.Charset, dbConfig.ParseTime, dbConfig.Loc)

	// 配置GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return fmt.Errorf("连接数据库失败: %w", err)
	}

	log.Println("数据库连接成功")
	DB = db

	return nil
}

// GetDB 获取数据库连接实例
func GetDB() *gorm.DB {
	return DB
}

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate(models ...interface{}) error {
	if DB == nil {
		return fmt.Errorf("数据库未连接")
	}

	err := DB.AutoMigrate(models...)
	if err != nil {
		return fmt.Errorf("自动迁移数据库表结构失败: %w", err)
	}

	log.Println("数据库表结构迁移成功")
	return nil
}

// Close 关闭数据库连接
func Close() error {
	if DB == nil {
		return nil
	}

	// 获取底层的sql.DB对象
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("获取底层sql.DB对象失败: %w", err)
	}

	// 关闭数据库连接
	err = sqlDB.Close()
	if err != nil {
		return fmt.Errorf("关闭数据库连接失败: %w", err)
	}

	log.Println("数据库连接已关闭")
	return nil
}
