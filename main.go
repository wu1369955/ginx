package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log"

	"github.com/wu136995/ginx/internal/api/middlewares"
	"github.com/wu136995/ginx/internal/api/routes"

	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/config"
	"github.com/wu136995/ginx/internal/database"
	"github.com/wu136995/ginx/internal/models"
)

func main() {
	// 加载配置
	err := config.LoadConfig("")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 连接数据库
	err = database.Connect()
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 自动迁移数据库表结构
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("迁移数据库表结构失败: %v", err)
	}

	// 设置Gin模式
	if config.AppConfig.Server.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
	}

	// 创建路由
	router := gin.New()

	// 设置中间件
	middlewares.SetupMiddlewares(router)

	// 设置路由
	routes.SetupRoutes(router)

	// 创建HTTP服务器
	port := config.AppConfig.Server.Port
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	// 启动服务器（非阻塞）
	go func() {
		log.Printf("服务器启动在 http://localhost:%s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("启动服务器失败: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务器...")

	// 设置关闭超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 优雅关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("服务器强制关闭: %v", err)
	}

	// 关闭数据库连接
	log.Println("正在关闭数据库连接...")
	if err := database.Close(); err != nil {
		log.Printf("关闭数据库连接失败: %v", err)
	}

	log.Println("服务器已退出")
}
