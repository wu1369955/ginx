package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/api/handlers"
)

// SetupFinanceRoutes 设置财务模块路由
func SetupFinanceRoutes(router *gin.Engine, financeHandler *handlers.FinanceHandler) {
	// 财务模块API路由组
	finance := router.Group("/api/v1/finance")
	{
		// 账户管理
		accounts := finance.Group("/accounts")
		{
			accounts.GET("", financeHandler.GetAccountList)
			accounts.GET("/:id", financeHandler.GetAccountDetail)
			accounts.POST("", financeHandler.CreateAccount)
			accounts.PUT("/:id", financeHandler.UpdateAccount)
			accounts.DELETE("/:id", financeHandler.DeleteAccount)
		}

		// 凭证管理
		vouchers := finance.Group("/vouchers")
		{
			vouchers.GET("", financeHandler.GetVoucherList)
			vouchers.GET("/:id", financeHandler.GetVoucherDetail)
			vouchers.POST("", financeHandler.CreateVoucher)
			vouchers.PUT("/:id", financeHandler.UpdateVoucher)
			vouchers.DELETE("/:id", financeHandler.DeleteVoucher)
		}

		// 付款管理
		payments := finance.Group("/payments")
		{
			payments.GET("", financeHandler.GetVoucherList)
			payments.GET("/:id", financeHandler.GetVoucherDetail)
			payments.POST("", financeHandler.CreateVoucher)
			payments.PUT("/:id", financeHandler.UpdateVoucher)
			payments.DELETE("/:id", financeHandler.DeleteVoucher)
		}

		// 收款管理
		receipts := finance.Group("/receipts")
		{
			receipts.GET("", financeHandler.GetVoucherList)
			receipts.GET("/:id", financeHandler.GetVoucherDetail)
			receipts.POST("", financeHandler.CreateVoucher)
			receipts.PUT("/:id", financeHandler.UpdateVoucher)
			receipts.DELETE("/:id", financeHandler.DeleteVoucher)
		}

		// 财务报表管理
		reports := finance.Group("/reports")
		{
			reports.GET("/balance", financeHandler.GetBalanceSheet)
			reports.GET("/profit", financeHandler.GetIncomeStatement)
			reports.GET("/cash-flow", financeHandler.GetCashFlowStatement)
			reports.GET("/trial", financeHandler.GetTrialBalance)
			reports.GET("/export", financeHandler.ExportFinancialReport)
		}
	}
}
