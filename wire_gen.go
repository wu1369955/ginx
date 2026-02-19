//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject

package main

import (
	"github.com/wu136995/ginx/internal/api/handlers"
	"github.com/wu136995/ginx/internal/services"
)

// Injectors from wire.go:

// InitializeHandlers 初始化所有处理器
func InitializeHandlers() (*handlers.UserHandler, *handlers.SalesHandler, *handlers.InventoryHandler, *handlers.PurchaseHandler, *handlers.FinanceHandler, *handlers.ProductionHandler, *handlers.HRHandler, *handlers.CRMHandler, error) {
	userService := services.NewUserService()
	salesService := services.NewSalesService()
	inventoryService := services.NewInventoryService()
	purchaseService := services.NewPurchaseService()
	financeService := services.NewFinanceService()
	productionService := services.NewProductionService()
	hrService := services.NewHRService()
	crmService := services.NewCRMService()
	userHandler := handlers.NewUserHandler(userService)
	salesHandler := handlers.NewSalesHandler(salesService)
	inventoryHandler := handlers.NewInventoryHandler(inventoryService)
	purchaseHandler := handlers.NewPurchaseHandler(purchaseService)
	financeHandler := handlers.NewFinanceHandler(financeService)
	productionHandler := handlers.NewProductionHandler(productionService)
	hrHandler := handlers.NewHRHandler(hrService)
	crmHandler := handlers.NewCRMHandler(crmService)
	return userHandler, salesHandler, inventoryHandler, purchaseHandler, financeHandler, productionHandler, hrHandler, crmHandler, nil
}
