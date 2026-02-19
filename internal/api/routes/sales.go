package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/api/handlers"
)

// SetupSalesRoutes 设置销售模块路由
func SetupSalesRoutes(router *gin.Engine) {
	// 销售模块API路由组
	sales := router.Group("/api/v1/sales")
	{
		// 客户管理
		customers := sales.Group("/customers")
		{
			customers.GET("", handlers.GetCustomerList)
			customers.GET("/:id", handlers.GetCustomerDetail)
			customers.POST("", handlers.CreateCustomer)
			customers.PUT("/:id", handlers.UpdateCustomer)
			customers.DELETE("/:id", handlers.DeleteCustomer)
			customers.POST("/:id/credit-evaluate", handlers.CustomerCreditEvaluate)
		}

		// 销售报价管理
		quotations := sales.Group("/quotations")
		{
			quotations.GET("", handlers.GetQuotationList)
			quotations.GET("/:id", handlers.GetQuotationDetail)
			quotations.POST("", handlers.CreateQuotation)
			quotations.PUT("/:id", handlers.UpdateQuotation)
			quotations.DELETE("/:id", handlers.DeleteQuotation)
			quotations.POST("/:id/approve", handlers.ApproveQuotation)
			quotations.POST("/:id/generate-order", handlers.GenerateOrderFromQuotation)
		}

		// 销售订单管理
		orders := sales.Group("/orders")
		{
			orders.GET("", handlers.GetOrderList)
			orders.GET("/:id", handlers.GetOrderDetail)
			orders.POST("", handlers.CreateOrder)
			orders.PUT("/:id", handlers.UpdateOrder)
			orders.DELETE("/:id", handlers.DeleteOrder)
			orders.POST("/:id/approve", handlers.ApproveOrder)
			orders.POST("/:id/cancel", handlers.CancelOrder)
			orders.POST("/:id/generate-delivery", handlers.GenerateDeliveryFromOrder)
		}

		// 销售发货管理
		deliveries := sales.Group("/deliveries")
		{
			deliveries.GET("", handlers.GetDeliveryList)
			deliveries.GET("/:id", handlers.GetDeliveryDetail)
			deliveries.POST("", handlers.CreateDelivery)
			deliveries.PUT("/:id", handlers.UpdateDelivery)
			deliveries.DELETE("/:id", handlers.DeleteDelivery)
		}

		// 销售发票管理
		invoices := sales.Group("/invoices")
		{
			invoices.GET("", handlers.GetInvoiceList)
			invoices.GET("/:id", handlers.GetInvoiceDetail)
			invoices.POST("", handlers.CreateInvoice)
			invoices.PUT("/:id", handlers.UpdateInvoice)
			invoices.DELETE("/:id", handlers.DeleteInvoice)
			invoices.POST("/:id/receive", handlers.ReceiveInvoicePayment)
		}

		// 销售退货管理
		returns := sales.Group("/returns")
		{
			returns.GET("", handlers.GetReturnList)
			returns.GET("/:id", handlers.GetReturnDetail)
			returns.POST("", handlers.CreateReturn)
			returns.PUT("/:id", handlers.UpdateReturn)
			returns.DELETE("/:id", handlers.DeleteReturn)
		}

		// 销售报表管理
	reports := sales.Group("/reports")
		{
			reports.GET("/order-execution", handlers.GetOrderExecutionReport)
			reports.GET("/customer-analysis", handlers.GetCustomerAnalysisReport)
			reports.GET("/product-trend", handlers.GetProductTrendReport)
			reports.GET("/export", handlers.ExportSalesReport)
		}
	}
}
