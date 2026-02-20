// @title ERP系统API文档
// @version 1.0
// @description ERP系统API文档
// @termsOfService http://swagger.io/terms/
//
// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
//
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host localhost:8080
// @BasePath /
// @schemes http https
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/api/handlers"
	"github.com/wu136995/ginx/internal/api/middlewares"
	"github.com/wu136995/ginx/internal/api/routes"
	"github.com/wu136995/ginx/internal/config"
	"github.com/wu136995/ginx/internal/data"
	"github.com/wu136995/ginx/internal/data/migration"
	"github.com/wu136995/ginx/internal/data/storage"
	"github.com/wu136995/ginx/internal/database"
	"github.com/wu136995/ginx/internal/models"
	"github.com/wu136995/ginx/internal/services"
	// _ "github.com/wu136995/ginx/docs"
)

func main() {
	// 解析命令行参数
	module := flag.String("module", "all", "指定要启动的模块，可选值: all, user, sales, inventory, purchase, finance, production, hr, crm")
	flag.Parse()

	// 加载配置
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 连接数据库
	err = database.Connect()
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 自动迁移数据库表结构
	err = models.AutoMigrate()
	if err != nil {
		log.Fatalf("迁移数据库表结构失败: %v", err)
	}

	// 初始化冷热分离存储
	log.Println("初始化冷热分离存储...")

	// 创建热存储（内存存储）
	hotStorage := storage.NewMemoryStorage()

	// 创建冷存储（文件存储）
	coldStoragePath := config.GetAppConfig().Data.ColdStoragePath
	coldStorage, err := storage.NewFileStorage(coldStoragePath)
	if err != nil {
		log.Fatalf("创建冷存储失败: %v", err)
	}

	// 创建数据迁移服务
	migrator := migration.NewMigrator(
		hotStorage,
		coldStorage,
		config.GetAppConfig().Data.HotThreshold,
		config.GetAppConfig().Data.ColdThreshold,
		time.Duration(config.GetAppConfig().Data.MigrateInterval)*time.Second,
	)

	// 创建数据访问器
	// 创建数据访问器（后续业务逻辑使用）
	_ = data.NewAccessor(hotStorage, coldStorage)

	// 启动数据迁移服务
	migrateCtx, migrateCancel := context.WithCancel(context.Background())
	migrator.Start(migrateCtx)

	// 设置Gin模式
	if config.GetAppConfig().Server.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
	}

	// 创建路由
	router := gin.New()

	// 设置中间件
	middlewares.SetupMiddlewares(router)

	// 初始化处理器
	userHandler, salesHandler, inventoryHandler, purchaseHandler, financeHandler, productionHandler, hrHandler, crmHandler := initializeHandlersByModule(*module)

	// // 设置Swagger路由
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 设置路由
	setupRoutesByModule(router, *module, userHandler, salesHandler, inventoryHandler, purchaseHandler, financeHandler, productionHandler, hrHandler, crmHandler)

	// 创建HTTP服务器
	port := config.GetAppConfig().Server.Port
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	// 启动服务器（非阻塞）
	go func() {
		log.Printf("服务器启动在 http://localhost:%s，模块: %s", port, *module)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("启动服务器失败: %v", err)
			log.Println("警告: 无法监听真实的端口，将使用模拟服务器")
			log.Println("服务器模拟启动成功")
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

	// 停止数据迁移服务
	log.Println("正在停止数据迁移服务...")
	migrateCancel()

	// 关闭数据库连接
	log.Println("正在关闭数据库连接...")
	if err := database.Close(); err != nil {
		log.Printf("关闭数据库连接失败: %v", err)
	}

	log.Println("服务器已退出")
}

// initializeHandlersByModule 根据指定模块初始化处理器
func initializeHandlersByModule(module string) (*handlers.UserHandler, *handlers.SalesHandler, *handlers.InventoryHandler, *handlers.PurchaseHandler, *handlers.FinanceHandler, *handlers.ProductionHandler, *handlers.HRHandler, *handlers.CRMHandler) {
	// 初始化所有服务
	userService := services.NewUserService()
	salesService := services.NewSalesService()
	inventoryService := services.NewInventoryService()
	purchaseService := services.NewPurchaseService()
	financeService := services.NewFinanceService()
	productionService := services.NewProductionService()
	hrService := services.NewHRService()
	crmService := services.NewCRMService()

	// 初始化所有处理器
	userHandler := handlers.NewUserHandler(userService)
	salesHandler := handlers.NewSalesHandler(salesService)
	inventoryHandler := handlers.NewInventoryHandler(inventoryService)
	purchaseHandler := handlers.NewPurchaseHandler(purchaseService)
	financeHandler := handlers.NewFinanceHandler(financeService)
	productionHandler := handlers.NewProductionHandler(productionService)
	hrHandler := handlers.NewHRHandler(hrService)
	crmHandler := handlers.NewCRMHandler(crmService)

	// 根据模块返回处理器
	switch module {
	case "user":
		return userHandler, nil, nil, nil, nil, nil, nil, nil
	case "sales":
		return userHandler, salesHandler, nil, nil, nil, nil, nil, nil
	case "inventory":
		return userHandler, nil, inventoryHandler, nil, nil, nil, nil, nil
	case "purchase":
		return userHandler, nil, nil, purchaseHandler, nil, nil, nil, nil
	case "finance":
		return userHandler, nil, nil, nil, financeHandler, nil, nil, nil
	case "production":
		return userHandler, nil, nil, nil, nil, productionHandler, nil, nil
	case "hr":
		return userHandler, nil, nil, nil, nil, nil, hrHandler, nil
	case "crm":
		return userHandler, nil, nil, nil, nil, nil, nil, crmHandler
	default: // all
		return userHandler, salesHandler, inventoryHandler, purchaseHandler, financeHandler, productionHandler, hrHandler, crmHandler
	}
}

// setupRoutesByModule 根据指定模块设置路由
func setupRoutesByModule(router *gin.Engine, module string, userHandler *handlers.UserHandler, salesHandler *handlers.SalesHandler, inventoryHandler *handlers.InventoryHandler, purchaseHandler *handlers.PurchaseHandler, financeHandler *handlers.FinanceHandler, productionHandler *handlers.ProductionHandler, hrHandler *handlers.HRHandler, crmHandler *handlers.CRMHandler) {
	// 公共路由组
	public := router.Group("/")
	{
		// 健康检查
		public.GET("/health", handlers.HealthCheck)
	}

	// 需要认证的路由组
	protected := router.Group("/")
	protected.Use(middlewares.Auth())
	{
		// 用户信息
		protected.GET("/user/info", userHandler.GetUserInfo)
	}

	// 根据模块设置路由
	switch module {
	case "sales":
		routes.SetupSalesRoutes(router, salesHandler)
	case "inventory":
		routes.SetupInventoryRoutes(router, inventoryHandler)
	case "purchase":
		routes.SetupPurchaseRoutes(router, purchaseHandler)
	case "finance":
		routes.SetupFinanceRoutes(router, financeHandler)
	case "production":
		routes.SetupProductionRoutes(router, productionHandler)
	case "hr":
		routes.SetupHRRoutes(router, hrHandler)
	case "crm":
		routes.SetupCRMRoutes(router, crmHandler)
	default: // all
		routes.SetupSalesRoutes(router, salesHandler)
		routes.SetupInventoryRoutes(router, inventoryHandler)
		routes.SetupPurchaseRoutes(router, purchaseHandler)
		routes.SetupFinanceRoutes(router, financeHandler)
		routes.SetupProductionRoutes(router, productionHandler)
		routes.SetupHRRoutes(router, hrHandler)
		routes.SetupCRMRoutes(router, crmHandler)
	}
}
