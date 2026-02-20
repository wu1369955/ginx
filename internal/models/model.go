package models

import (
	"log"

	"github.com/wu136995/ginx/internal/database"
)

// AllModels 所有模型列表
var AllModels = []interface{}{
	// 用户模型
	&User{},

	// CRM模型
	&CRMAccount{},
	&CRMContact{},
	&CRMActivity{},
	&CRMCase{},
	&CRMCaseComment{},
	&CRMLead{},
	&CRMOpportunity{},
	&CRMCampaign{},

	// 销售模型
	&SalesCustomer{},
	&SalesProduct{},
	&SalesProductCategory{},
	&SalesQuote{},
	&SalesQuoteItem{},
	&SalesOrder{},
	&SalesOrderItem{},
	&SalesDelivery{},
	&SalesDeliveryItem{},
	&SalesInvoice{},
	&SalesInvoiceItem{},
	&SalesReturn{},
	&SalesReturnItem{},
	&SalesCustomerCredit{},

	// 库存模型
	&InventoryWarehouse{},
	&InventoryLocation{},
	&InventoryItem{},
	&InventoryItemCategory{},
	&InventoryOnHand{},
	&InventoryTransaction{},
	&InventoryCount{},
	&InventoryCountItem{},

	// 采购模型
	&PurchaseVendor{},
	&PurchasePlan{},
	&PurchasePlanItem{},
	&PurchaseOrder{},
	&PurchaseOrderItem{},
	&PurchaseReceipt{},
	&PurchaseReceiptItem{},
	&PurchaseInvoice{},
	&PurchaseInvoiceItem{},
	&PurchaseReturn{},
	&PurchaseReturnItem{},

	// 财务模型
	&FinanceAccount{},
}

// AutoMigrate 自动迁移所有模型
func AutoMigrate() error {
	db := database.GetDB()
	if db == nil {
		log.Println("警告: 数据库连接为nil，跳过自动迁移")
		return nil
	}
	return database.AutoMigrate(AllModels...)
}
