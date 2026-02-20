package models

import (
	"time"

	"gorm.io/gorm"
)

// PurchaseVendor 供应商表模型
type PurchaseVendor struct {
	ID            string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	VendorNo      string         `json:"vendor_no" gorm:"unique;not null;type:varchar(20)"`
	Name          string         `json:"name" gorm:"not null;type:varchar(100)"`
	ContactPerson string         `json:"contact_person" gorm:"type:varchar(50)"`
	Phone         string         `json:"phone" gorm:"type:varchar(20)"`
	Email         string         `json:"email" gorm:"type:varchar(100)"`
	Address       string         `json:"address" gorm:"type:varchar(255)"`
	TaxNo         string         `json:"tax_no" gorm:"type:varchar(30)"`
	BankName      string         `json:"bank_name" gorm:"type:varchar(100)"`
	BankAccount   string         `json:"bank_account" gorm:"type:varchar(50)"`
	Category      string         `json:"category" gorm:"type:varchar(50)"`
	CreditLimit   float64        `json:"credit_limit" gorm:"type:decimal(18,2);default:0"`
	LeadTime      int            `json:"lead_time" gorm:"type:int;default:0"`
	PaymentTerms  string         `json:"payment_terms" gorm:"type:varchar(50)"`
	Remarks       string         `json:"remarks" gorm:"type:text"`
	Status        string         `json:"status" gorm:"type:varchar(20);default:'active'"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy     string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Orders    []PurchaseOrder    `json:"orders,omitempty" gorm:"foreignKey:VendorID"`
	Receipts  []PurchaseReceipt  `json:"receipts,omitempty" gorm:"foreignKey:VendorID"`
	Invoices  []PurchaseInvoice  `json:"invoices,omitempty" gorm:"foreignKey:VendorID"`
	Returns   []PurchaseReturn   `json:"returns,omitempty" gorm:"foreignKey:VendorID"`
}

// TableName 指定表名
func (PurchaseVendor) TableName() string {
	return "purchase_vendors"
}

// PurchasePlan 采购计划体表模型
type PurchasePlan struct {
	ID           string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	PlanNo       string         `json:"plan_no" gorm:"unique;not null;type:varchar(20)"`
	Name         string         `json:"name" gorm:"not null;type:varchar(100)"`
	Description  string         `json:"description" gorm:"type:text"`
	StartDate    time.Time      `json:"start_date" gorm:"not null;type:date"`
	EndDate      time.Time      `json:"end_date" gorm:"not null;type:date"`
	TotalAmount  float64        `json:"total_amount" gorm:"not null;type:decimal(18,2)"`
	Status       string         `json:"status" gorm:"type:varchar(20);default:'draft'"`
	CreatedBy    string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt    time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy    string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Items []PurchasePlanItem `json:"items,omitempty" gorm:"foreignKey:PlanID"`
}

// TableName 指定表名
func (PurchasePlan) TableName() string {
	return "purchase_plans"
}

// PurchasePlanItem 采购计划表明细表模型
type PurchasePlanItem struct {
	ID             string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	PlanID         string         `json:"plan_id" gorm:"not null;type:varchar(36)"`
	ItemID         string         `json:"item_id" gorm:"not null;type:varchar(36)"`
	Quantity       float64        `json:"quantity" gorm:"not null;type:decimal(18,4)"`
	EstimatedPrice float64        `json:"estimated_price" gorm:"not null;type:decimal(18,2)"`
	EstimatedAmount float64       `json:"estimated_amount" gorm:"not null;type:decimal(18,2)"`
	NeedDate       time.Time      `json:"need_date" gorm:"not null;type:date"`
	CreatedBy      string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt      time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy      string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Plan  PurchasePlan   `json:"plan,omitempty" gorm:"foreignKey:PlanID"`
	Item  InventoryItem  `json:"item,omitempty" gorm:"foreignKey:ItemID"`
}

// TableName 指定表名
func (PurchasePlanItem) TableName() string {
	return "purchase_plan_items"
}

// PurchaseOrder 采购订单表模型
type PurchaseOrder struct {
	ID           string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	OrderNo      string         `json:"order_no" gorm:"unique;not null;type:varchar(20)"`
	VendorID     string         `json:"vendor_id" gorm:"not null;type:varchar(36)"`
	OrderDate    time.Time      `json:"order_date" gorm:"not null;type:date"`
	DeliveryDate *time.Time     `json:"delivery_date" gorm:"type:date"`
	TotalAmount  float64        `json:"total_amount" gorm:"not null;type:decimal(18,2)"`
	Status       string         `json:"status" gorm:"type:varchar(20);default:'pending'"`
	PaymentTerms string         `json:"payment_terms" gorm:"type:varchar(50)"`
	Remarks      string         `json:"remarks" gorm:"type:text"`
	CreatedBy    string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt    time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy    string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Vendor   PurchaseVendor     `json:"vendor,omitempty" gorm:"foreignKey:VendorID"`
	Items    []PurchaseOrderItem `json:"items,omitempty" gorm:"foreignKey:OrderID"`
	Receipts []PurchaseReceipt   `json:"receipts,omitempty" gorm:"foreignKey:OrderID"`
	Invoices []PurchaseInvoice   `json:"invoices,omitempty" gorm:"foreignKey:OrderID"`
	Returns  []PurchaseReturn    `json:"returns,omitempty" gorm:"foreignKey:OrderID"`
}

// TableName 指定表名
func (PurchaseOrder) TableName() string {
	return "purchase_orders"
}

// PurchaseOrderItem 采购订单明细表模型
type PurchaseOrderItem struct {
	ID              string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	OrderID         string         `json:"order_id" gorm:"not null;type:varchar(36)"`
	ItemID          string         `json:"item_id" gorm:"not null;type:varchar(36)"`
	Quantity        float64        `json:"quantity" gorm:"not null;type:decimal(18,4)"`
	UnitPrice       float64        `json:"unit_price" gorm:"not null;type:decimal(18,2)"`
	Discount        float64        `json:"discount" gorm:"type:decimal(18,2);default:0"`
	Amount          float64        `json:"amount" gorm:"not null;type:decimal(18,2)"`
	ReceivedQuantity float64       `json:"received_quantity" gorm:"type:decimal(18,4);default:0"`
	CreatedBy       string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt       time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy       string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Order   PurchaseOrder    `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Item    InventoryItem    `json:"item,omitempty" gorm:"foreignKey:ItemID"`
	ReceiptItems []PurchaseReceiptItem `json:"receipt_items,omitempty" gorm:"foreignKey:OrderItemID"`
}

// TableName 指定表名
func (PurchaseOrderItem) TableName() string {
	return "purchase_order_items"
}

// PurchaseReceipt 采购收货单表模型
type PurchaseReceipt struct {
	ID            string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	ReceiptNo     string         `json:"receipt_no" gorm:"unique;not null;type:varchar(20)"`
	OrderID       string         `json:"order_id" gorm:"not null;type:varchar(36)"`
	VendorID      string         `json:"vendor_id" gorm:"not null;type:varchar(36)"`
	ReceiptDate   time.Time      `json:"receipt_date" gorm:"not null;type:date"`
	TotalQuantity float64        `json:"total_quantity" gorm:"not null;type:decimal(18,4)"`
	TotalAmount   float64        `json:"total_amount" gorm:"not null;type:decimal(18,2)"`
	Status        string         `json:"status" gorm:"type:varchar(20);default:'pending'"`
	Remarks       string         `json:"remarks" gorm:"type:text"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy     string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Order    PurchaseOrder       `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Vendor   PurchaseVendor      `json:"vendor,omitempty" gorm:"foreignKey:VendorID"`
	Items    []PurchaseReceiptItem `json:"items,omitempty" gorm:"foreignKey:ReceiptID"`
}

// TableName 指定表名
func (PurchaseReceipt) TableName() string {
	return "purchase_receipts"
}

// PurchaseReceiptItem 采购收货单明细表模型
type PurchaseReceiptItem struct {
	ID          string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	ReceiptID   string         `json:"receipt_id" gorm:"not null;type:varchar(36)"`
	OrderItemID string         `json:"order_item_id" gorm:"not null;type:varchar(36)"`
	ItemID      string         `json:"item_id" gorm:"not null;type:varchar(36)"`
	Quantity    float64        `json:"quantity" gorm:"not null;type:decimal(18,4)"`
	UnitPrice   float64        `json:"unit_price" gorm:"not null;type:decimal(18,2)"`
	Amount      float64        `json:"amount" gorm:"not null;type:decimal(18,2)"`
	WarehouseID string         `json:"warehouse_id" gorm:"not null;type:varchar(36)"`
	LocationID  string         `json:"location_id" gorm:"type:varchar(36)"`
	CreatedBy   string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt   time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy   string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Receipt   PurchaseReceipt   `json:"receipt,omitempty" gorm:"foreignKey:ReceiptID"`
	OrderItem PurchaseOrderItem `json:"order_item,omitempty" gorm:"foreignKey:OrderItemID"`
	Item      InventoryItem     `json:"item,omitempty" gorm:"foreignKey:ItemID"`
	Warehouse InventoryWarehouse `json:"warehouse,omitempty" gorm:"foreignKey:WarehouseID"`
	Location  *InventoryLocation `json:"location,omitempty" gorm:"foreignKey:LocationID"`
}

// TableName 指定表名
func (PurchaseReceiptItem) TableName() string {
	return "purchase_receipt_items"
}

// PurchaseInvoice 采购发票表模型
type PurchaseInvoice struct {
	ID          string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	InvoiceNo   string         `json:"invoice_no" gorm:"unique;not null;type:varchar(20)"`
	OrderID     string         `json:"order_id" gorm:"not null;type:varchar(36)"`
	VendorID    string         `json:"vendor_id" gorm:"not null;type:varchar(36)"`
	InvoiceDate time.Time      `json:"invoice_date" gorm:"not null;type:date"`
	DueDate     time.Time      `json:"due_date" gorm:"not null;type:date"`
	TotalAmount float64        `json:"total_amount" gorm:"not null;type:decimal(18,2)"`
	PaidAmount  float64        `json:"paid_amount" gorm:"type:decimal(18,2);default:0"`
	Status      string         `json:"status" gorm:"type:varchar(20);default:'unpaid'"`
	Remarks     string         `json:"remarks" gorm:"type:text"`
	CreatedBy   string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt   time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy   string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Order    PurchaseOrder      `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Vendor   PurchaseVendor     `json:"vendor,omitempty" gorm:"foreignKey:VendorID"`
	Items    []PurchaseInvoiceItem `json:"items,omitempty" gorm:"foreignKey:InvoiceID"`
}

// TableName 指定表名
func (PurchaseInvoice) TableName() string {
	return "purchase_invoices"
}

// PurchaseInvoiceItem 采购发票明细表模型
type PurchaseInvoiceItem struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	InvoiceID string         `json:"invoice_id" gorm:"not null;type:varchar(36)"`
	ItemID    string         `json:"item_id" gorm:"not null;type:varchar(36)"`
	Quantity  float64        `json:"quantity" gorm:"not null;type:decimal(18,4)"`
	UnitPrice float64        `json:"unit_price" gorm:"not null;type:decimal(18,2)"`
	Discount  float64        `json:"discount" gorm:"type:decimal(18,2);default:0"`
	Amount    float64        `json:"amount" gorm:"not null;type:decimal(18,2)"`
	CreatedBy string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Invoice PurchaseInvoice `json:"invoice,omitempty" gorm:"foreignKey:InvoiceID"`
	Item    InventoryItem   `json:"item,omitempty" gorm:"foreignKey:ItemID"`
}

// TableName 指定表名
func (PurchaseInvoiceItem) TableName() string {
	return "purchase_invoice_items"
}

// PurchaseReturn 采购退货单表模型
type PurchaseReturn struct {
	ID         string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	ReturnNo   string         `json:"return_no" gorm:"unique;not null;type:varchar(20)"`
	OrderID    string         `json:"order_id" gorm:"not null;type:varchar(36)"`
	VendorID   string         `json:"vendor_id" gorm:"not null;type:varchar(36)"`
	ReturnDate time.Time      `json:"return_date" gorm:"not null;type:date"`
	TotalAmount float64       `json:"total_amount" gorm:"not null;type:decimal(18,2)"`
	Status     string         `json:"status" gorm:"type:varchar(20);default:'pending'"`
	Reason     string         `json:"reason" gorm:"type:text"`
	Remarks    string         `json:"remarks" gorm:"type:text"`
	CreatedBy  string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt  time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy  string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Order   PurchaseOrder      `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Vendor  PurchaseVendor     `json:"vendor,omitempty" gorm:"foreignKey:VendorID"`
	Items   []PurchaseReturnItem `json:"items,omitempty" gorm:"foreignKey:ReturnID"`
}

// TableName 指定表名
func (PurchaseReturn) TableName() string {
	return "purchase_returns"
}

// PurchaseReturnItem 采购退货单明细表模型
type PurchaseReturnItem struct {
	ID          string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	ReturnID    string         `json:"return_id" gorm:"not null;type:varchar(36)"`
	ItemID      string         `json:"item_id" gorm:"not null;type:varchar(36)"`
	Quantity    float64        `json:"quantity" gorm:"not null;type:decimal(18,4)"`
	UnitPrice   float64        `json:"unit_price" gorm:"not null;type:decimal(18,2)"`
	Amount      float64        `json:"amount" gorm:"not null;type:decimal(18,2)"`
	WarehouseID string         `json:"warehouse_id" gorm:"not null;type:varchar(36)"`
	CreatedBy   string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt   time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy   string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Return    PurchaseReturn     `json:"return,omitempty" gorm:"foreignKey:ReturnID"`
	Item      InventoryItem      `json:"item,omitempty" gorm:"foreignKey:ItemID"`
	Warehouse InventoryWarehouse `json:"warehouse,omitempty" gorm:"foreignKey:WarehouseID"`
}

// TableName 指定表名
func (PurchaseReturnItem) TableName() string {
	return "purchase_return_items"
}
