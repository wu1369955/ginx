package schemas

import "time"

// 通用响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 分页响应结构
type PaginatedResponse struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

// 客户管理相关

// GetCustomerListRequest 获取客户列表请求

type GetCustomerListRequest struct {
	Page       int    `form:"page" binding:"omitempty,min=1"`
	PageSize   int    `form:"pageSize" binding:"omitempty,min=1,max=100"`
	CustomerNo string `form:"customerNo" binding:"omitempty"`
	Name       string `form:"name" binding:"omitempty"`
	Contact    string `form:"contact" binding:"omitempty"`
	Phone      string `form:"phone" binding:"omitempty"`
	Status     string `form:"status" binding:"omitempty,oneof=active inactive"`
	Category   string `form:"category" binding:"omitempty"`
	Region     string `form:"region" binding:"omitempty"`
}

// CustomerResponse 客户响应

type CustomerResponse struct {
	ID           string    `json:"id"`
	CustomerNo   string    `json:"customerNo"`
	Name         string    `json:"name"`
	Contact      string    `json:"contact"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Address      string    `json:"address"`
	TaxNo        string    `json:"taxNo"`
	BankName     string    `json:"bankName,omitempty"`
	BankAccount  string    `json:"bankAccount,omitempty"`
	Category     string    `json:"category"`
	Region       string    `json:"region"`
	CreditLimit  float64   `json:"creditLimit,omitempty"`
	CreditDays   int       `json:"creditDays,omitempty"`
	Remarks      string    `json:"remarks,omitempty"`
	Status       string    `json:"status"`
	CreatedBy    string    `json:"createdBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedBy    string    `json:"updatedBy"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// CreateCustomerRequest 创建客户请求

type CreateCustomerRequest struct {
	Name        string  `json:"name" binding:"required"`
	Contact     string  `json:"contact" binding:"required"`
	Phone       string  `json:"phone" binding:"required"`
	Email       string  `json:"email" binding:"omitempty,email"`
	Address     string  `json:"address" binding:"omitempty"`
	TaxNo       string  `json:"taxNo" binding:"omitempty"`
	BankName    string  `json:"bankName" binding:"omitempty"`
	BankAccount string  `json:"bankAccount" binding:"omitempty"`
	Category    string  `json:"category" binding:"required"`
	Region      string  `json:"region" binding:"required"`
	CreditLimit float64 `json:"creditLimit" binding:"omitempty,min=0"`
	CreditDays  int     `json:"creditDays" binding:"omitempty,min=0"`
	Remarks     string  `json:"remarks" binding:"omitempty"`
}

// UpdateCustomerRequest 更新客户请求

type UpdateCustomerRequest struct {
	Name        string  `json:"name" binding:"omitempty"`
	Contact     string  `json:"contact" binding:"omitempty"`
	Phone       string  `json:"phone" binding:"omitempty"`
	Email       string  `json:"email" binding:"omitempty,email"`
	Address     string  `json:"address" binding:"omitempty"`
	TaxNo       string  `json:"taxNo" binding:"omitempty"`
	BankName    string  `json:"bankName" binding:"omitempty"`
	BankAccount string  `json:"bankAccount" binding:"omitempty"`
	Category    string  `json:"category" binding:"omitempty"`
	Region      string  `json:"region" binding:"omitempty"`
	CreditLimit float64 `json:"creditLimit" binding:"omitempty,min=0"`
	CreditDays  int     `json:"creditDays" binding:"omitempty,min=0"`
	Remarks     string  `json:"remarks" binding:"omitempty"`
	Status      string  `json:"status" binding:"omitempty,oneof=active inactive"`
}

// CustomerCreditEvaluateRequest 客户信用评估请求

type CustomerCreditEvaluateRequest struct {
	EvaluationDate string `json:"evaluationDate" binding:"required,datetime=2006-01-02"`
	CreditScore    int    `json:"creditScore" binding:"required,min=0,max=100"`
	RiskLevel      string `json:"riskLevel" binding:"required,oneof=low medium high"`
	Remarks        string `json:"remarks" binding:"omitempty"`
}

// 销售报价相关

// GetQuotationListRequest 获取销售报价列表请求

type GetQuotationListRequest struct {
	Page             int    `form:"page" binding:"omitempty,min=1"`
	PageSize         int    `form:"pageSize" binding:"omitempty,min=1,max=100"`
	QuotationNo      string `form:"quotationNo" binding:"omitempty"`
	CustomerId       string `form:"customerId" binding:"omitempty"`
	CustomerName     string `form:"customerName" binding:"omitempty"`
	Status           string `form:"status" binding:"omitempty,oneof=draft submitted approved rejected expired"`
	CreateDateStart  string `form:"createDateStart" binding:"omitempty,datetime=2006-01-02"`
	CreateDateEnd    string `form:"createDateEnd" binding:"omitempty,datetime=2006-01-02"`
	ExpiryDateStart  string `form:"expiryDateStart" binding:"omitempty,datetime=2006-01-02"`
	ExpiryDateEnd    string `form:"expiryDateEnd" binding:"omitempty,datetime=2006-01-02"`
}

// QuotationItem 报价单明细

type QuotationItem struct {
	ID            string  `json:"id,omitempty"`
	ProductId     string  `json:"productId" binding:"required"`
	ProductCode   string  `json:"productCode,omitempty"`
	ProductName   string  `json:"productName,omitempty"`
	Specification string  `json:"specification,omitempty"`
	Unit          string  `json:"unit,omitempty"`
	Quantity      float64 `json:"quantity" binding:"required,min=0.0001"`
	UnitPrice     float64 `json:"unitPrice" binding:"required,min=0.01"`
	Amount        float64 `json:"amount,omitempty"`
	Remark        string  `json:"remark" binding:"omitempty"`
}

// QuotationResponse 报价单响应

type QuotationResponse struct {
	ID             string          `json:"id"`
	QuotationNo    string          `json:"quotationNo"`
	CustomerId     string          `json:"customerId"`
	CustomerName   string          `json:"customerName"`
	Contact        string          `json:"contact"`
	Phone          string          `json:"phone"`
	Address        string          `json:"address,omitempty"`
	QuotationDate  string          `json:"quotationDate"`
	ExpiryDate     string          `json:"expiryDate"`
	PaymentTerms   string          `json:"paymentTerms,omitempty"`
	ShippingTerms  string          `json:"shippingTerms,omitempty"`
	Remarks        string          `json:"remarks,omitempty"`
	TotalAmount    float64         `json:"totalAmount"`
	Status         string          `json:"status"`
	Items          []QuotationItem `json:"items,omitempty"`
	CreatedBy      string          `json:"createdBy"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedBy      string          `json:"updatedBy"`
	UpdatedAt      time.Time       `json:"updatedAt"`
}

// CreateQuotationRequest 创建销售报价请求

type CreateQuotationRequest struct {
	CustomerId    string          `json:"customerId" binding:"required"`
	QuotationDate string          `json:"quotationDate" binding:"required,datetime=2006-01-02"`
	ExpiryDate    string          `json:"expiryDate" binding:"required,datetime=2006-01-02"`
	PaymentTerms  string          `json:"paymentTerms" binding:"omitempty"`
	ShippingTerms string          `json:"shippingTerms" binding:"omitempty"`
	Remarks       string          `json:"remarks" binding:"omitempty"`
	Items         []QuotationItem `json:"items" binding:"required,dive"`
}

// UpdateQuotationRequest 更新销售报价请求

type UpdateQuotationRequest struct {
	CustomerId    string          `json:"customerId" binding:"omitempty"`
	QuotationDate string          `json:"quotationDate" binding:"omitempty,datetime=2006-01-02"`
	ExpiryDate    string          `json:"expiryDate" binding:"omitempty,datetime=2006-01-02"`
	PaymentTerms  string          `json:"paymentTerms" binding:"omitempty"`
	ShippingTerms string          `json:"shippingTerms" binding:"omitempty"`
	Remarks       string          `json:"remarks" binding:"omitempty"`
	Items         []QuotationItem `json:"items" binding:"omitempty,dive"`
}

// ApproveQuotationRequest 审批销售报价请求

type ApproveQuotationRequest struct {
	Approved bool   `json:"approved" binding:"required"`
	Remarks  string `json:"remarks" binding:"omitempty"`
}

// GenerateOrderFromQuotationRequest 从报价单生成销售订单请求

type GenerateOrderFromQuotationRequest struct {
	OrderDate    string `json:"orderDate" binding:"required,datetime=2006-01-02"`
	DeliveryDate string `json:"deliveryDate" binding:"required,datetime=2006-01-02"`
}

// 销售订单相关

// GetOrderListRequest 获取销售订单列表请求

type GetOrderListRequest struct {
	Page             int     `form:"page" binding:"omitempty,min=1"`
	PageSize         int     `form:"pageSize" binding:"omitempty,min=1,max=100"`
	OrderNo          string  `form:"orderNo" binding:"omitempty"`
	CustomerId       string  `form:"customerId" binding:"omitempty"`
	CustomerName     string  `form:"customerName" binding:"omitempty"`
	Status           string  `form:"status" binding:"omitempty,oneof=draft approved partially_delivered fully_delivered cancelled"`
	CreateDateStart  string  `form:"createDateStart" binding:"omitempty,datetime=2006-01-02"`
	CreateDateEnd    string  `form:"createDateEnd" binding:"omitempty,datetime=2006-01-02"`
	DeliveryDateStart string `form:"deliveryDateStart" binding:"omitempty,datetime=2006-01-02"`
	DeliveryDateEnd   string `form:"deliveryDateEnd" binding:"omitempty,datetime=2006-01-02"`
	TotalAmountStart  float64 `form:"totalAmountStart" binding:"omitempty,min=0"`
	TotalAmountEnd    float64 `form:"totalAmountEnd" binding:"omitempty,min=0"`
}

// OrderItem 订单明细

type OrderItem struct {
	ID                string  `json:"id,omitempty"`
	ProductId         string  `json:"productId" binding:"required"`
	ProductCode       string  `json:"productCode,omitempty"`
	ProductName       string  `json:"productName,omitempty"`
	Specification     string  `json:"specification,omitempty"`
	Unit              string  `json:"unit,omitempty"`
	Quantity          float64 `json:"quantity" binding:"required,min=0.0001"`
	UnitPrice         float64 `json:"unitPrice" binding:"required,min=0.01"`
	Amount            float64 `json:"amount,omitempty"`
	DeliveredQuantity float64 `json:"deliveredQuantity,omitempty"`
	Remark            string  `json:"remark" binding:"omitempty"`
}

// OrderResponse 订单响应

type OrderResponse struct {
	ID            string      `json:"id"`
	OrderNo       string      `json:"orderNo"`
	CustomerId    string      `json:"customerId"`
	CustomerName  string      `json:"customerName"`
	Contact       string      `json:"contact"`
	Phone         string      `json:"phone"`
	Address       string      `json:"address,omitempty"`
	OrderDate     string      `json:"orderDate"`
	DeliveryDate  string      `json:"deliveryDate"`
	PaymentTerms  string      `json:"paymentTerms,omitempty"`
	ShippingTerms string      `json:"shippingTerms,omitempty"`
	Remarks       string      `json:"remarks,omitempty"`
	TotalAmount   float64     `json:"totalAmount"`
	Status        string      `json:"status"`
	Items         []OrderItem `json:"items,omitempty"`
	CreatedBy     string      `json:"createdBy"`
	CreatedAt     time.Time   `json:"createdAt"`
	UpdatedBy     string      `json:"updatedBy"`
	UpdatedAt     time.Time   `json:"updatedAt"`
}

// CreateOrderRequest 创建销售订单请求

type CreateOrderRequest struct {
	CustomerId    string      `json:"customerId" binding:"required"`
	OrderDate     string      `json:"orderDate" binding:"required,datetime=2006-01-02"`
	DeliveryDate  string      `json:"deliveryDate" binding:"required,datetime=2006-01-02"`
	PaymentTerms  string      `json:"paymentTerms" binding:"omitempty"`
	ShippingTerms string      `json:"shippingTerms" binding:"omitempty"`
	Remarks       string      `json:"remarks" binding:"omitempty"`
	Items         []OrderItem `json:"items" binding:"required,dive"`
}

// UpdateOrderRequest 更新销售订单请求

type UpdateOrderRequest struct {
	CustomerId    string      `json:"customerId" binding:"omitempty"`
	OrderDate     string      `json:"orderDate" binding:"omitempty,datetime=2006-01-02"`
	DeliveryDate  string      `json:"deliveryDate" binding:"omitempty,datetime=2006-01-02"`
	PaymentTerms  string      `json:"paymentTerms" binding:"omitempty"`
	ShippingTerms string      `json:"shippingTerms" binding:"omitempty"`
	Remarks       string      `json:"remarks" binding:"omitempty"`
	Items         []OrderItem `json:"items" binding:"omitempty,dive"`
}

// ApproveOrderRequest 审批销售订单请求

type ApproveOrderRequest struct {
	Approved bool   `json:"approved" binding:"required"`
	Remarks  string `json:"remarks" binding:"omitempty"`
}

// CancelOrderRequest 取消销售订单请求

type CancelOrderRequest struct {
	Reason string `json:"reason" binding:"required"`
}

// DeliveryItem 发货单明细

type DeliveryItem struct {
	OrderItemId       string  `json:"orderItemId" binding:"required"`
	DeliveredQuantity float64 `json:"deliveredQuantity" binding:"required,min=0.0001"`
	Remark            string  `json:"remark" binding:"omitempty"`
}

// GenerateDeliveryFromOrderRequest 从销售订单生成发货单请求

type GenerateDeliveryFromOrderRequest struct {
	DeliveryDate string         `json:"deliveryDate" binding:"required,datetime=2006-01-02"`
	WarehouseId  string         `json:"warehouseId" binding:"required"`
	Items        []DeliveryItem `json:"items" binding:"required,dive"`
}

// 销售发货相关

// GetDeliveryListRequest 获取销售发货列表请求

type GetDeliveryListRequest struct {
	Page             int    `form:"page" binding:"omitempty,min=1"`
	PageSize         int    `form:"pageSize" binding:"omitempty,min=1,max=100"`
	DeliveryNo       string `form:"deliveryNo" binding:"omitempty"`
	OrderId          string `form:"orderId" binding:"omitempty"`
	OrderNo          string `form:"orderNo" binding:"omitempty"`
	CustomerId       string `form:"customerId" binding:"omitempty"`
	CustomerName     string `form:"customerName" binding:"omitempty"`
	DeliveryDateStart string `form:"deliveryDateStart" binding:"omitempty,datetime=2006-01-02"`
	DeliveryDateEnd   string `form:"deliveryDateEnd" binding:"omitempty,datetime=2006-01-02"`
}

// DeliveryItemResponse 发货单明细响应

type DeliveryItemResponse struct {
	ID                string  `json:"id,omitempty"`
	OrderItemId       string  `json:"orderItemId"`
	ProductId         string  `json:"productId"`
	ProductCode       string  `json:"productCode,omitempty"`
	ProductName       string  `json:"productName,omitempty"`
	Specification     string  `json:"specification,omitempty"`
	Unit              string  `json:"unit,omitempty"`
	OrderedQuantity   float64 `json:"orderedQuantity"`
	DeliveredQuantity float64 `json:"deliveredQuantity"`
	Remark            string  `json:"remark,omitempty"`
}

// DeliveryResponse 发货单响应

type DeliveryResponse struct {
	ID            string                `json:"id"`
	DeliveryNo    string                `json:"deliveryNo"`
	OrderId       string                `json:"orderId"`
	OrderNo       string                `json:"orderNo"`
	CustomerId    string                `json:"customerId"`
	CustomerName  string                `json:"customerName"`
	Contact       string                `json:"contact"`
	Phone         string                `json:"phone"`
	Address       string                `json:"address,omitempty"`
	DeliveryDate  string                `json:"deliveryDate"`
	WarehouseId   string                `json:"warehouseId"`
	WarehouseName string                `json:"warehouseName,omitempty"`
	Carrier       string                `json:"carrier,omitempty"`
	TrackingNo    string                `json:"trackingNo,omitempty"`
	Remarks       string                `json:"remarks,omitempty"`
	Status        string                `json:"status"`
	TotalQuantity float64               `json:"totalQuantity,omitempty"`
	Items         []DeliveryItemResponse `json:"items,omitempty"`
	CreatedBy     string                `json:"createdBy"`
	CreatedAt     time.Time             `json:"createdAt"`
	UpdatedBy     string                `json:"updatedBy"`
	UpdatedAt     time.Time             `json:"updatedAt"`
}

// CreateDeliveryRequest 创建销售发货请求

type CreateDeliveryRequest struct {
	OrderId     string         `json:"orderId" binding:"required"`
	DeliveryDate string         `json:"deliveryDate" binding:"required,datetime=2006-01-02"`
	WarehouseId  string         `json:"warehouseId" binding:"required"`
	Carrier      string         `json:"carrier" binding:"omitempty"`
	TrackingNo   string         `json:"trackingNo" binding:"omitempty"`
	Remarks      string         `json:"remarks" binding:"omitempty"`
	Items        []DeliveryItem `json:"items" binding:"required,dive"`
}

// UpdateDeliveryRequest 更新销售发货请求

type UpdateDeliveryRequest struct {
	Carrier    string         `json:"carrier" binding:"omitempty"`
	TrackingNo string         `json:"trackingNo" binding:"omitempty"`
	Remarks    string         `json:"remarks" binding:"omitempty"`
	Items      []DeliveryItem `json:"items" binding:"omitempty,dive"`
}

// 销售发票相关

// GetInvoiceListRequest 获取销售发票列表请求

type GetInvoiceListRequest struct {
	Page             int    `form:"page" binding:"omitempty,min=1"`
	PageSize         int    `form:"pageSize" binding:"omitempty,min=1,max=100"`
	InvoiceNo        string `form:"invoiceNo" binding:"omitempty"`
	CustomerId       string `form:"customerId" binding:"omitempty"`
	CustomerName     string `form:"customerName" binding:"omitempty"`
	OrderId          string `form:"orderId" binding:"omitempty"`
	OrderNo          string `form:"orderNo" binding:"omitempty"`
	InvoiceDateStart string `form:"invoiceDateStart" binding:"omitempty,datetime=2006-01-02"`
	InvoiceDateEnd   string `form:"invoiceDateEnd" binding:"omitempty,datetime=2006-01-02"`
	Status           string `form:"status" binding:"omitempty,oneof=unpaid partially_paid fully_paid"`
}

// InvoiceItem 发票明细

type InvoiceItem struct {
	ID          string  `json:"id,omitempty"`
	ProductId   string  `json:"productId" binding:"required"`
	ProductCode string  `json:"productCode,omitempty"`
	ProductName string  `json:"productName,omitempty"`
	Specification string `json:"specification,omitempty"`
	Quantity    float64 `json:"quantity" binding:"required,min=0.0001"`
	UnitPrice   float64 `json:"unitPrice" binding:"required,min=0.01"`
	TaxRate     float64 `json:"taxRate" binding:"required,min=0"`
	TaxAmount   float64 `json:"taxAmount,omitempty"`
	TotalAmount float64 `json:"totalAmount,omitempty"`
}

// InvoiceResponse 发票响应

type InvoiceResponse struct {
	ID             string        `json:"id"`
	InvoiceNo      string        `json:"invoiceNo"`
	CustomerId     string        `json:"customerId"`
	CustomerName   string        `json:"customerName"`
	CustomerTaxNo  string        `json:"customerTaxNo,omitempty"`
	OrderId        string        `json:"orderId"`
	OrderNo        string        `json:"orderNo"`
	InvoiceDate    string        `json:"invoiceDate"`
	DueDate        string        `json:"dueDate"`
	TotalAmount    float64       `json:"totalAmount"`
	TaxAmount      float64       `json:"taxAmount,omitempty"`
	NetAmount      float64       `json:"netAmount,omitempty"`
	PaidAmount     float64       `json:"paidAmount"`
	BalanceAmount  float64       `json:"balanceAmount"`
	Status         string        `json:"status"`
	Items          []InvoiceItem `json:"items,omitempty"`
	CreatedBy      string        `json:"createdBy"`
	CreatedAt      time.Time     `json:"createdAt"`
	UpdatedBy      string        `json:"updatedBy"`
	UpdatedAt      time.Time     `json:"updatedAt"`
}

// CreateInvoiceRequest 创建销售发票请求

type CreateInvoiceRequest struct {
	CustomerId  string        `json:"customerId" binding:"required"`
	OrderId     string        `json:"orderId" binding:"required"`
	InvoiceDate string        `json:"invoiceDate" binding:"required,datetime=2006-01-02"`
	DueDate     string        `json:"dueDate" binding:"required,datetime=2006-01-02"`
	Items       []InvoiceItem `json:"items" binding:"required,dive"`
}

// UpdateInvoiceRequest 更新销售发票请求

type UpdateInvoiceRequest struct {
	DueDate string        `json:"dueDate" binding:"omitempty,datetime=2006-01-02"`
	Items   []InvoiceItem `json:"items" binding:"omitempty,dive"`
}

// ReceiveInvoicePaymentRequest 标记发票已收款请求

type ReceiveInvoicePaymentRequest struct {
	Amount        float64 `json:"amount" binding:"required,min=0.01"`
	ReceiveDate   string  `json:"receiveDate" binding:"required,datetime=2006-01-02"`
	PaymentMethod string  `json:"paymentMethod" binding:"required"`
	Remarks       string  `json:"remarks" binding:"omitempty"`
}

// 销售退货相关

// GetReturnListRequest 获取销售退货列表请求

type GetReturnListRequest struct {
	Page             int    `form:"page" binding:"omitempty,min=1"`
	PageSize         int    `form:"pageSize" binding:"omitempty,min=1,max=100"`
	ReturnNo         string `form:"returnNo" binding:"omitempty"`
	CustomerId       string `form:"customerId" binding:"omitempty"`
	CustomerName     string `form:"customerName" binding:"omitempty"`
	OrderId          string `form:"orderId" binding:"omitempty"`
	OrderNo          string `form:"orderNo" binding:"omitempty"`
	ReturnDateStart string `form:"returnDateStart" binding:"omitempty,datetime=2006-01-02"`
	ReturnDateEnd   string `form:"returnDateEnd" binding:"omitempty,datetime=2006-01-02"`
}

// ReturnItem 退货单明细

type ReturnItem struct {
	OrderItemId string  `json:"orderItemId" binding:"required"`
	Quantity    float64 `json:"quantity" binding:"required,min=0.0001"`
	UnitPrice   float64 `json:"unitPrice" binding:"required,min=0.01"`
	Amount      float64 `json:"amount,omitempty"`
	Reason      string  `json:"reason" binding:"required"`
}

// ReturnItemResponse 退货单明细响应

type ReturnItemResponse struct {
	ID          string  `json:"id,omitempty"`
	OrderItemId string  `json:"orderItemId"`
	ProductId   string  `json:"productId"`
	ProductCode string  `json:"productCode,omitempty"`
	ProductName string  `json:"productName,omitempty"`
	Specification string `json:"specification,omitempty"`
	Unit        string  `json:"unit,omitempty"`
	Quantity    float64 `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice"`
	Amount      float64 `json:"amount"`
	Reason      string  `json:"reason"`
}

// ReturnResponse 退货单响应

type ReturnResponse struct {
	ID          string              `json:"id"`
	ReturnNo    string              `json:"returnNo"`
	CustomerId  string              `json:"customerId"`
	CustomerName string             `json:"customerName"`
	Contact     string              `json:"contact"`
	Phone       string              `json:"phone"`
	OrderId     string              `json:"orderId"`
	OrderNo     string              `json:"orderNo"`
	ReturnDate  string              `json:"returnDate"`
	Reason      string              `json:"reason"`
	Remarks     string              `json:"remarks,omitempty"`
	TotalAmount float64             `json:"totalAmount"`
	Status      string              `json:"status"`
	Items       []ReturnItemResponse `json:"items,omitempty"`
	CreatedBy   string              `json:"createdBy"`
	CreatedAt   time.Time           `json:"createdAt"`
	UpdatedBy   string              `json:"updatedBy"`
	UpdatedAt   time.Time           `json:"updatedAt"`
}

// CreateReturnRequest 创建销售退货请求

type CreateReturnRequest struct {
	CustomerId  string       `json:"customerId" binding:"required"`
	OrderId     string       `json:"orderId" binding:"required"`
	ReturnDate  string       `json:"returnDate" binding:"required,datetime=2006-01-02"`
	Reason      string       `json:"reason" binding:"required"`
	Remarks     string       `json:"remarks" binding:"omitempty"`
	Items       []ReturnItem `json:"items" binding:"required,dive"`
}

// UpdateReturnRequest 更新销售退货请求

type UpdateReturnRequest struct {
	Reason  string       `json:"reason" binding:"omitempty"`
	Remarks string       `json:"remarks" binding:"omitempty"`
	Items   []ReturnItem `json:"items" binding:"omitempty,dive"`
}

// 销售报表相关

// GetOrderExecutionReportRequest 获取销售订单执行报表请求

type GetOrderExecutionReportRequest struct {
	StartDate  string `form:"startDate" binding:"required,datetime=2006-01-02"`
	EndDate    string `form:"endDate" binding:"required,datetime=2006-01-02"`
	CustomerId string `form:"customerId" binding:"omitempty"`
	ProductId  string `form:"productId" binding:"omitempty"`
}

// OrderExecutionReportItem 订单执行报表项

type OrderExecutionReportItem struct {
	OrderNo         string  `json:"orderNo"`
	CustomerName    string  `json:"customerName"`
	OrderDate       string  `json:"orderDate"`
	Amount          float64 `json:"amount"`
	Status          string  `json:"status"`
	DeliveredAmount float64 `json:"deliveredAmount"`
	DeliveredRate   float64 `json:"deliveredRate"`
}

// OrderExecutionReportResponse 订单执行报表响应

type OrderExecutionReportResponse struct {
	Period       string                    `json:"period"`
	TotalOrders  int                       `json:"totalOrders"`
	TotalAmount  float64                   `json:"totalAmount"`
	Orders       []OrderExecutionReportItem `json:"orders"`
}

// GetCustomerAnalysisReportRequest 获取客户销售分析报表请求

type GetCustomerAnalysisReportRequest struct {
	StartDate string `form:"startDate" binding:"required,datetime=2006-01-02"`
	EndDate   string `form:"endDate" binding:"required,datetime=2006-01-02"`
}

// CustomerAnalysisReportItem 客户销售分析报表项

type CustomerAnalysisReportItem struct {
	CustomerId      string  `json:"customerId"`
	CustomerName    string  `json:"customerName"`
	TotalOrders     int     `json:"totalOrders"`
	TotalAmount     float64 `json:"totalAmount"`
	AvgOrderAmount  float64 `json:"avgOrderAmount"`
	PaymentRate     float64 `json:"paymentRate"`
}

// CustomerAnalysisReportResponse 客户销售分析报表响应

type CustomerAnalysisReportResponse struct {
	Period     string                        `json:"period"`
	Customers  []CustomerAnalysisReportItem  `json:"customers"`
}

// GetProductTrendReportRequest 获取产品销售趋势报表请求

type GetProductTrendReportRequest struct {
	StartDate string `form:"startDate" binding:"required,datetime=2006-01-02"`
	EndDate   string `form:"endDate" binding:"required,datetime=2006-01-02"`
	ProductId string `form:"productId" binding:"omitempty"`
}

// ProductTrendReportItem 产品销售趋势报表项

type ProductTrendReportItem struct {
	Period   string  `json:"period"`
	Quantity float64 `json:"quantity"`
	Amount   float64 `json:"amount"`
}

// ProductTrendReportResponse 产品销售趋势报表响应

type ProductTrendReportResponse struct {
	Period  string                      `json:"period"`
	Trends  []ProductTrendReportItem    `json:"trends"`
}

// ExportSalesReportRequest 导出销售报表请求

type ExportSalesReportRequest struct {
	Type       string `form:"type" binding:"required,oneof=order_execution customer_analysis product_trend"`
	StartDate  string `form:"startDate" binding:"required,datetime=2006-01-02"`
	EndDate    string `form:"endDate" binding:"required,datetime=2006-01-02"`
	CustomerId string `form:"customerId" binding:"omitempty"`
	ProductId  string `form:"productId" binding:"omitempty"`
}
