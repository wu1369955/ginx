//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/wu136995/ginx/internal/api/handlers"
	"github.com/wu136995/ginx/internal/services"
)

// WireSet 依赖注入集合
var WireSet = wire.NewSet(
	services.NewUserService,
	services.NewSalesService,
	services.NewInventoryService,
	services.NewPurchaseService,
	services.NewFinanceService,
	services.NewProductionService,
	services.NewHRService,
	services.NewCRMService,
	handlers.NewUserHandler,
	handlers.NewSalesHandler,
	handlers.NewInventoryHandler,
	handlers.NewPurchaseHandler,
	handlers.NewFinanceHandler,
	handlers.NewProductionHandler,
	handlers.NewHRHandler,
	handlers.NewCRMHandler,
)

// InitializeHandlers 初始化所有处理器
func InitializeHandlers() (*handlers.UserHandler, *handlers.SalesHandler, *handlers.InventoryHandler, *handlers.PurchaseHandler, *handlers.FinanceHandler, *handlers.ProductionHandler, *handlers.HRHandler, *handlers.CRMHandler) {
	wire.Build(WireSet)
	return nil, nil, nil, nil, nil, nil, nil, nil
}
