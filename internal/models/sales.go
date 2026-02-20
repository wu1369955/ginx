package models

import (
	"time"

	"gorm.io/gorm"
)

// SalesCustomer 客户表模型
type SalesCustomer struct {
	ID            string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	CustomerNo    string         `json:"customer_no" gorm:"unique;not null;type:varchar(20)"`
	Name          string         `json:"name" gorm:"not null;type:varchar(100)"`
	ContactPerson string         `json:"contact_person" gorm:"type:varchar(50)"`
	Phone         string         `json:"phone" gorm:"type:varchar(20)"`
	Email         string         `json:"email" gorm:"type:varchar(100)"`
	Address       string         `json:"address" gorm:"type:varchar(255)"`
	TaxNo         string         `json:"tax_no" gorm:"type:varchar(30)"`
	BankName      string         `json:"bank_name" gorm:"type:varchar(100)"`
	BankAccount   string         `json:"bank_account" gorm:"type:varchar(50)"`
	Category      string         `json:"category" gorm:"type:varchar(50)"`
	Region        string         `json:"region" gorm:"type:varchar(50)"`
	CreditLimit   float64        `json:"credit_limit" gorm:"type:decimal(18,2);default:0"`
	CreditDays    int            `json:"credit_days" gorm:"type:int;default:0"`
	Remarks       string         `json:"remarks" gorm:"type:text"`
	Status        string         `json:"status" gorm:"type:varchar(20);default:'active'"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy     string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Quotes     []SalesQuote     `json:"quotes,omitempty" gorm:"foreignKey:CustomerID"`
	Orders     []SalesOrder     `json:"orders,omitempty" gorm:"foreignKey:CustomerID"`
	Deliveries []SalesDelivery  `json:"deliveries,omitempty" gorm:"foreignKey:CustomerID"`
	Invoices   []SalesInvoice   `json:"invoices,omitempty" gorm:"foreignKey:CustomerID"`
	Returns    []SalesReturn    `json:"returns,omitempty" gorm:"foreignKey:CustomerID"`
	Credits    []SalesCustomerCredit `json:"credits,omitempty" gorm:"foreignKey:CustomerID"`
}

// TableName 指定表名
func (SalesCustomer) TableName() string {
	return "sales_customers"
}

// SalesProduct 产品表模型
type SalesProduct struct {
	ID          string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	ProductNo   string         `json:"product_no" gorm:"unique;not null;type:varchar(20)"`
	Name        string         `json:"name" gorm:"not null;type:varchar(100)"`
	Description string         `json:"description" gorm:"type:text"`
	CategoryID  string         `json:"category_id" gorm:"type:varchar(36)"`
	Unit        string         `json:"unit" gorm:"not null;type:varchar(10)"`
	Price       float64        `json:"price" gorm:"not null;type:decimal(18,2)"`
	Status      string         `json:"status" gorm:"type:varchar(20);default:'active'"`
	CreatedBy   string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt   time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy   string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Category    *SalesProductCategory `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	QuoteItems  []SalesQuoteItem      `json:"quote_items,omitempty" gorm:"foreignKey:ProductID"`
	OrderItems  []SalesOrderItem      `json:"order_items,omitempty" gorm:"foreignKey:ProductID"`
	DeliveryItems []SalesDeliveryItem `json:"delivery_items,omitempty" gorm:"foreignKey:ProductID"`
	InvoiceItems []SalesInvoiceItem   `json:"invoice_items,omitempty" gorm:"foreignKey:ProductID"`
	ReturnItems  []SalesReturnItem    `json:"return_items,omitempty" gorm:"foreignKey:ProductID"`
}

// TableName 指定表名
func (SalesProduct) TableName() string {
	return "sales_products"
}

// SalesProductCategory 产品类别表模型
type SalesProductCategory struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Code      string         `json:"code" gorm:"unique;not null;type:varchar(20)"`
	Name      string         `json:"name" gorm:"not null;type:varchar(100)"`
	ParentID  string         `json:"parent_id" gorm:"type:varchar(36)"`
	CreatedBy string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Parent   *SalesProductCategory `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children []SalesProductCategory `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	Products []SalesProduct        `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
}

// TableName 指定表名
func (SalesProductCategory) TableName() string {
	return "sales_product_categories"
}

// SalesQuote 销售报价单表模型
type SalesQuote struct {
	ID         string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	QuoteNo    string         `json:"quote_no" gorm:"unique;not null;type:varchar(20)"`
	CustomerID string         `json:"customer_id" gorm:"not null;type:varchar(36)"`
	QuoteDate  time.Time      `json:"quote_date" gorm:"not null;type:date"`
	ValidUntil time.Time      `json:"valid_until" gorm:"not null;type:date"`
	TotalAmount float64       `json:"total_amount" gorm:"not null;type:decimal(18,2)"`
	Status     string         `json:"status" gorm:"type:varchar(20);default:'draft'"`
	Remarks    string         `json:"remarks" gorm:"type:text"`
	CreatedBy  string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt  time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy  string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Customer  SalesCustomer    `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
	Items     []SalesQuoteItem `json:"items,omitempty" gorm:"foreignKey:QuoteID"`
	Order     *SalesOrder      `json:"order,omitempty" gorm:"foreignKey:QuoteID"`
}

// TableName 指定表名
func (SalesQuote) TableName() string {
	return "sales_quotes"
}

// SalesQuoteItem 销售报价单明细表模型
type SalesQuoteItem struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	QuoteID   string         `json:"quote_id" gorm:"not null;type:varchar(36)"`
	ProductID string         `json:"product_id" gorm:"not null;type:varchar(36)"`
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
	Quote   SalesQuote    `json:"quote,omitempty" gorm:"foreignKey:QuoteID"`
	Product SalesProduct  `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

// TableName 指定表名
func (SalesQuoteItem) TableName() string {
	return "sales_quote_items"
}

// SalesOrder 销售订单表模型
type SalesOrder struct {
	ID         string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	OrderNo    string         `json:"order_no" gorm:"unique;not null;type:varchar(20)"`
	CustomerID string         `json:"customer_id" gorm:"not null;type:varchar(36)"`
	QuoteID    string         `json:"quote_id" gorm:"type:varchar(36)"`
	OrderDate  time.Time      `json:"order_date" gorm:"not null;type:date"`
	DeliveryDate *time.Time   `json:"delivery_date" gorm:"type:date"`
	TotalAmount float64       `json:"total_amount" gorm:"not null;type:decimal(18,2)"`
	Status     string         `json:"status" gorm:"type:varchar(20);default:'pending'"`
	Remarks    string         `json:"remarks" gorm:"type:text"`
	CreatedBy  string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt  time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy  string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Customer   SalesCustomer    `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
	Quote      *SalesQuote      `json:"quote,omitempty" gorm:"foreignKey:QuoteID"`
	Items      []SalesOrderItem `json:"items,omitempty" gorm:"foreignKey:OrderID"`
	Deliveries []SalesDelivery  `json:"deliveries,omitempty" gorm:"foreignKey:OrderID"`
	Invoices   []SalesInvoice   `json:"invoices,omitempty" gorm:"foreignKey:OrderID"`
	Returns    []SalesReturn    `json:"returns,omitempty" gorm:"foreignKey:OrderID"`
}

// TableName 指定表名
func (SalesOrder) TableName() string {
	return "sales_orders"
}

// SalesOrderItem 销售订单明细表模型
type SalesOrderItem struct {
	ID              string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	OrderID         string         `json:"order_id" gorm:"not null;type:varchar(36)"`
	ProductID       string         `json:"product_id" gorm:"not null;type:varchar(36)"`
	Quantity        float64        `json:"quantity" gorm:"not null;type:decimal(18,4)"`
	UnitPrice       float64        `json:"unit_price" gorm:"not null;type:decimal(18,2)"`
	Discount        float64        `json:"discount" gorm:"type:decimal(18,2);default:0"`
	Amount          float64        `json:"amount" gorm:"not null;type:decimal(18,2)"`
	ShippedQuantity float64        `json:"shipped_quantity" gorm:"type:decimal(18,4);default:0"`
	CreatedBy       string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt       time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy       string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Order    SalesOrder      `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Product  SalesProduct    `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	DeliveryItems []SalesDeliveryItem `json:"delivery_items,omitempty" gorm:"foreignKey:OrderItemID"`
}

// TableName 指定表名
func (SalesOrderItem) TableName() string {
	return "sales_order_items"
}

// SalesDelivery 销售发货单表模型
type SalesDelivery struct {
	ID            string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	DeliveryNo    string         `json:"delivery_no" gorm:"unique;not null;type:varchar(20)"`
	OrderID       string         `json:"order_id" gorm:"not null;type:varchar(36)"`
	CustomerID    string         `json:"customer_id" gorm:"not null;type:varchar(36)"`
	DeliveryDate  time.Time      `json:"delivery_date" gorm:"not null;type:date"`
	TotalQuantity float64        `json:"total_quantity" gorm:"not null;type:decimal(18,4)"`
	Status        string         `json:"status" gorm:"type:varchar(20);default:'pending'"`
	Remarks       string         `json:"remarks" gorm:"type:text"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy     string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Order     SalesOrder         `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Customer  SalesCustomer      `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
	Items     []SalesDeliveryItem `json:"items,omitempty" gorm:"foreignKey:DeliveryID"`
}

// TableName 指定表名
func (SalesDelivery) TableName() string {
	return "sales_deliveries"
}

// SalesDeliveryItem 销售发货单明细表模型
type SalesDeliveryItem struct {
	ID          string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	DeliveryID  string         `json:"delivery_id" gorm:"not null;type:varchar(36)"`
	OrderItemID string         `json:"order_item_id" gorm:"not null;type:varchar(36)"`
	ProductID   string         `json:"product_id" gorm:"not null;type:varchar(36)"`
	Quantity    float64        `json:"quantity" gorm:"not null;type:decimal(18,4)"`
	CreatedBy   string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt   time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy   string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Delivery    SalesDelivery   `json:"delivery,omitempty" gorm:"foreignKey:DeliveryID"`
	OrderItem   SalesOrderItem  `json:"order_item,omitempty" gorm:"foreignKey:OrderItemID"`
	Product     SalesProduct    `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

// TableName 指定表名
func (SalesDeliveryItem) TableName() string {
	return "sales_delivery_items"
}

// SalesInvoice 销售发票表模型
type SalesInvoice struct {
	ID          string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	InvoiceNo   string         `json:"invoice_no" gorm:"unique;not null;type:varchar(20)"`
	OrderID     string         `json:"order_id" gorm:"not null;type:varchar(36)"`
	CustomerID  string         `json:"customer_id" gorm:"not null;type:varchar(36)"`
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
	Order     SalesOrder        `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Customer  SalesCustomer     `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
	Items     []SalesInvoiceItem `json:"items,omitempty" gorm:"foreignKey:InvoiceID"`
}

// TableName 指定表名
func (SalesInvoice) TableName() string {
	return "sales_invoices"
}

// SalesInvoiceItem 销售发票明细表模型
type SalesInvoiceItem struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	InvoiceID string         `json:"invoice_id" gorm:"not null;type:varchar(36)"`
	ProductID string         `json:"product_id" gorm:"not null;type:varchar(36)"`
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
	Invoice  SalesInvoice  `json:"invoice,omitempty" gorm:"foreignKey:InvoiceID"`
	Product  SalesProduct  `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

// TableName 指定表名
func (SalesInvoiceItem) TableName() string {
	return "sales_invoice_items"
}

// SalesReturn 销售退货单表模型
type SalesReturn struct {
	ID         string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	ReturnNo   string         `json:"return_no" gorm:"unique;not null;type:varchar(20)"`
	OrderID    string         `json:"order_id" gorm:"not null;type:varchar(36)"`
	CustomerID string         `json:"customer_id" gorm:"not null;type:varchar(36)"`
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
	Order     SalesOrder      `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Customer  SalesCustomer   `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
	Items     []SalesReturnItem `json:"items,omitempty" gorm:"foreignKey:ReturnID"`
}

// TableName 指定表名
func (SalesReturn) TableName() string {
	return "sales_returns"
}

// SalesReturnItem 销售退货单明细表模型
type SalesReturnItem struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	ReturnID  string         `json:"return_id" gorm:"not null;type:varchar(36)"`
	ProductID string         `json:"product_id" gorm:"not null;type:varchar(36)"`
	Quantity  float64        `json:"quantity" gorm:"not null;type:decimal(18,4)"`
	UnitPrice float64        `json:"unit_price" gorm:"not null;type:decimal(18,2)"`
	Amount    float64        `json:"amount" gorm:"not null;type:decimal(18,2)"`
	CreatedBy string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Return   SalesReturn   `json:"return,omitempty" gorm:"foreignKey:ReturnID"`
	Product  SalesProduct  `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

// TableName 指定表名
func (SalesReturnItem) TableName() string {
	return "sales_return_items"
}

// SalesCustomerCredit 客户信用评估表模型
type SalesCustomerCredit struct {
	ID            string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	CustomerID    string         `json:"customer_id" gorm:"not null;type:varchar(36)"`
	EvaluationDate time.Time     `json:"evaluation_date" gorm:"not null;type:date"`
	CreditScore   int            `json:"credit_score" gorm:"not null;type:int"`
	RiskLevel     string         `json:"risk_level" gorm:"not null;type:varchar(20)"`
	Remarks       string         `json:"remarks" gorm:"type:text"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy     string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Customer  SalesCustomer  `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
}

// TableName 指定表名
func (SalesCustomerCredit) TableName() string {
	return "sales_customer_credits"
}
