package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/api/handlers"
)

// SetupPurchaseRoutes 设置采购模块路由
func SetupPurchaseRoutes(router *gin.Engine, purchaseHandler *handlers.PurchaseHandler) {
	// 采购模块API路由组
	purchase := router.Group("/api/v1/purchase")
	{
		// 供应商管理
		suppliers := purchase.Group("/suppliers")
		{
			suppliers.GET("", purchaseHandler.GetSupplierList)
			suppliers.GET("/:id", purchaseHandler.GetSupplierDetail)
			suppliers.POST("", purchaseHandler.CreateSupplier)
			suppliers.PUT("/:id", purchaseHandler.UpdateSupplier)
			suppliers.DELETE("/:id", purchaseHandler.DeleteSupplier)
			suppliers.GET("/:id/contacts", purchaseHandler.GetSupplierContacts)
			suppliers.POST("/:id/contacts", purchaseHandler.AddSupplierContact)
		}

		// 采购申请管理
		requisitions := purchase.Group("/requisitions")
		{
			requisitions.GET("", purchaseHandler.GetRequisitionList)
			requisitions.GET("/:id", purchaseHandler.GetRequisitionDetail)
			requisitions.POST("", purchaseHandler.CreateRequisition)
			requisitions.PUT("/:id", purchaseHandler.UpdateRequisition)
			requisitions.DELETE("/:id", purchaseHandler.DeleteRequisition)
			requisitions.POST("/:id/submit", purchaseHandler.SubmitRequisition)
			requisitions.POST("/:id/approve", purchaseHandler.ApproveRequisition)
			requisitions.POST("/:id/reject", purchaseHandler.RejectRequisition)
		}

		// 采购订单管理
		orders := purchase.Group("/orders")
		{
			orders.GET("", purchaseHandler.GetPurchaseOrderList)
			orders.GET("/:id", purchaseHandler.GetPurchaseOrderDetail)
			orders.POST("", purchaseHandler.CreatePurchaseOrder)
			orders.PUT("/:id", purchaseHandler.UpdatePurchaseOrder)
			orders.DELETE("/:id", purchaseHandler.DeletePurchaseOrder)
			orders.POST("/:id/submit", purchaseHandler.SubmitPurchaseOrder)
			orders.POST("/:id/approve", purchaseHandler.ApprovePurchaseOrder)
			orders.POST("/:id/reject", purchaseHandler.RejectPurchaseOrder)
			orders.POST("/:id/close", purchaseHandler.ClosePurchaseOrder)
		}

		// 采购收货管理
		receipts := purchase.Group("/receipts")
		{
			receipts.GET("", purchaseHandler.GetPurchaseReceiptList)
			receipts.GET("/:id", purchaseHandler.GetPurchaseReceiptDetail)
			receipts.POST("", purchaseHandler.CreatePurchaseReceipt)
			receipts.PUT("/:id", purchaseHandler.UpdatePurchaseReceipt)
			receipts.DELETE("/:id", purchaseHandler.DeletePurchaseReceipt)
			receipts.POST("/:id/complete", purchaseHandler.CompletePurchaseReceipt)
		}

		// 采购发票管理
		invoices := purchase.Group("/invoices")
		{
			invoices.GET("", purchaseHandler.GetPurchaseInvoiceList)
			invoices.GET("/:id", purchaseHandler.GetPurchaseInvoiceDetail)
			invoices.POST("", purchaseHandler.CreatePurchaseInvoice)
			invoices.PUT("/:id", purchaseHandler.UpdatePurchaseInvoice)
			invoices.DELETE("/:id", purchaseHandler.DeletePurchaseInvoice)
			invoices.POST("/:id/verify", purchaseHandler.VerifyPurchaseInvoice)
			invoices.POST("/:id/pay", purchaseHandler.PayPurchaseInvoice)
		}

		// 采购退货管理
		returns := purchase.Group("/returns")
		{
			returns.GET("", purchaseHandler.GetPurchaseReturnList)
			returns.GET("/:id", purchaseHandler.GetPurchaseReturnDetail)
			returns.POST("", purchaseHandler.CreatePurchaseReturn)
			returns.PUT("/:id", purchaseHandler.UpdatePurchaseReturn)
			returns.DELETE("/:id", purchaseHandler.DeletePurchaseReturn)
			returns.POST("/:id/complete", purchaseHandler.CompletePurchaseReturn)
		}

		// 采购报表管理
		reports := purchase.Group("/reports")
		{
			reports.GET("/summary", purchaseHandler.GetPurchaseSummaryReport)
			reports.GET("/detail", purchaseHandler.GetPurchaseDetailReport)
			reports.GET("/supplier", purchaseHandler.GetSupplierAnalysisReport)
			reports.GET("/price", purchaseHandler.GetPriceAnalysisReport)
			reports.GET("/forecast", purchaseHandler.GetPurchaseForecastReport)
			reports.GET("/export", purchaseHandler.ExportPurchaseReport)
		}
	}
}
