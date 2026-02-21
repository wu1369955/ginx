package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/api/handlers"
)

// SetupSalesRoutes 设置销售模块路由
func SetupSalesRoutes(router *gin.Engine, salesHandler *handlers.SalesHandler) {
	// 销售模块API路由组
	sales := router.Group("/api/v1/sales")
	{
		// 客户管理
		customers := sales.Group("/customers")
		{
			customers.GET("", salesHandler.GetSalesCustomerList)
			customers.GET("/:id", salesHandler.GetSalesCustomerDetail)
			customers.POST("", salesHandler.CreateSalesCustomer)
			customers.PUT("/:id", salesHandler.UpdateSalesCustomer)
			customers.DELETE("/:id", salesHandler.DeleteSalesCustomer)
			customers.POST("/:id/credit-evaluate", salesHandler.CustomerCreditEvaluate)
		}

		// 销售报价管理
		quotations := sales.Group("/quotations")
		{
			quotations.GET("", salesHandler.GetQuotationList)
			quotations.GET("/:id", salesHandler.GetQuotationDetail)
			quotations.POST("", salesHandler.CreateQuotation)
			quotations.PUT("/:id", salesHandler.UpdateQuotation)
			quotations.DELETE("/:id", salesHandler.DeleteQuotation)
			quotations.POST("/:id/approve", salesHandler.ApproveQuotation)
			quotations.POST("/:id/generate-order", salesHandler.GenerateOrderFromQuotation)
		}

		// 销售订单管理
		orders := sales.Group("/orders")
		{
			orders.GET("", salesHandler.GetOrderList)
			orders.GET("/:id", salesHandler.GetOrderDetail)
			orders.POST("", salesHandler.CreateOrder)
			orders.PUT("/:id", salesHandler.UpdateOrder)
			orders.DELETE("/:id", salesHandler.DeleteOrder)
			orders.POST("/:id/approve", salesHandler.ApproveOrder)
			orders.POST("/:id/cancel", salesHandler.CancelOrder)
			orders.POST("/:id/generate-delivery", salesHandler.GenerateDeliveryFromOrder)
		}

		// 销售发货管理
		deliveries := sales.Group("/deliveries")
		{
			deliveries.GET("", salesHandler.GetDeliveryList)
			deliveries.GET("/:id", salesHandler.GetDeliveryDetail)
			deliveries.POST("", salesHandler.CreateDelivery)
			deliveries.PUT("/:id", salesHandler.UpdateDelivery)
			deliveries.DELETE("/:id", salesHandler.DeleteDelivery)
		}

		// 销售发票管理
		invoices := sales.Group("/invoices")
		{
			invoices.GET("", salesHandler.GetInvoiceList)
			invoices.GET("/:id", salesHandler.GetInvoiceDetail)
			invoices.POST("", salesHandler.CreateInvoice)
			invoices.PUT("/:id", salesHandler.UpdateInvoice)
			invoices.DELETE("/:id", salesHandler.DeleteInvoice)
			invoices.POST("/:id/receive", salesHandler.ReceiveInvoicePayment)
		}

		// 销售退货管理
		returns := sales.Group("/returns")
		{
			returns.GET("", salesHandler.GetReturnList)
			returns.GET("/:id", salesHandler.GetReturnDetail)
			returns.POST("", salesHandler.CreateReturn)
			returns.PUT("/:id", salesHandler.UpdateReturn)
			returns.DELETE("/:id", salesHandler.DeleteReturn)
		}

		// 销售报表管理
		reports := sales.Group("/reports")
		{
			reports.GET("/order-execution", salesHandler.GetOrderExecutionReport)
			reports.GET("/customer-analysis", salesHandler.GetSalesCustomerAnalysisReport)
			reports.GET("/product-trend", salesHandler.GetProductTrendReport)
			reports.GET("/export", salesHandler.ExportSalesReport)
		}
	}
}
