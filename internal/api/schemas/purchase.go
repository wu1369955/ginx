package schemas

// 供应商相关结构体

// SupplierCreateRequest 创建供应商请求
 type SupplierCreateRequest struct {
	Code        string `json:"code" binding:"required"`
	Name        string `json:"name" binding:"required"`
	ContactName string `json:"contact_name" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Address     string `json:"address" binding:"required"`
	TaxNumber   string `json:"tax_number"`
	BankName    string `json:"bank_name"`
	BankAccount string `json:"bank_account"`
	Category    string `json:"category"`
	Rating      int    `json:"rating"`
	Remarks     string `json:"remarks"`
}

// SupplierUpdateRequest 更新供应商请求
 type SupplierUpdateRequest struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	ContactName string `json:"contact_name"`
	Phone       string `json:"phone"`
	Email       string `json:"email" binding:"omitempty,email"`
	Address     string `json:"address"`
	TaxNumber   string `json:"tax_number"`
	BankName    string `json:"bank_name"`
	BankAccount string `json:"bank_account"`
	Category    string `json:"category"`
	Rating      int    `json:"rating"`
	Remarks     string `json:"remarks"`
}

// SupplierResponse 供应商响应
 type SupplierResponse struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	ContactName string `json:"contact_name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	TaxNumber   string `json:"tax_number"`
	BankName    string `json:"bank_name"`
	BankAccount string `json:"bank_account"`
	Category    string `json:"category"`
	Rating      int    `json:"rating"`
	Remarks     string `json:"remarks"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// SupplierContactCreateRequest 创建供应商联系人请求
 type SupplierContactCreateRequest struct {
	SupplierID  uint   `json:"supplier_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Position    string `json:"position"`
	Phone       string `json:"phone" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	IsPrimary   bool   `json:"is_primary"`
	Remarks     string `json:"remarks"`
}

// SupplierContactResponse 供应商联系人响应
 type SupplierContactResponse struct {
	ID         uint   `json:"id"`
	SupplierID uint   `json:"supplier_id"`
	Name       string `json:"name"`
	Position   string `json:"position"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	IsPrimary  bool   `json:"is_primary"`
	Remarks    string `json:"remarks"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

// 采购申请相关结构体

// RequisitionCreateRequest 创建采购申请请求
 type RequisitionCreateRequest struct {
	Code        string                    `json:"code" binding:"required"`
	DepartmentID uint                     `json:"department_id" binding:"required"`
	ApplicantID  uint                     `json:"applicant_id" binding:"required"`
	ApplicationDate string                 `json:"application_date" binding:"required"`
	EstimatedAmount float64               `json:"estimated_amount" binding:"required"`
	Reason      string                    `json:"reason" binding:"required"`
	Items       []RequisitionItemRequest `json:"items" binding:"required,dive"`
	Remarks     string                    `json:"remarks"`
}

// RequisitionItemRequest 采购申请项目请求
 type RequisitionItemRequest struct {
	MaterialID  uint    `json:"material_id" binding:"required"`
	Quantity    float64 `json:"quantity" binding:"required"`
	EstimatedPrice float64 `json:"estimated_price" binding:"required"`
	RequiredDate string `json:"required_date" binding:"required"`
	Purpose     string `json:"purpose" binding:"required"`
}

// RequisitionUpdateRequest 更新采购申请请求
 type RequisitionUpdateRequest struct {
	DepartmentID uint                     `json:"department_id"`
	ApplicantID  uint                     `json:"applicant_id"`
	ApplicationDate string                 `json:"application_date"`
	EstimatedAmount float64               `json:"estimated_amount"`
	Reason      string                    `json:"reason"`
	Items       []RequisitionItemRequest `json:"items" binding:"omitempty,dive"`
	Remarks     string                    `json:"remarks"`
}

// RequisitionResponse 采购申请响应
 type RequisitionResponse struct {
	ID             uint                    `json:"id"`
	Code           string                  `json:"code"`
	DepartmentID   uint                    `json:"department_id"`
	DepartmentName string                  `json:"department_name"`
	ApplicantID    uint                    `json:"applicant_id"`
	ApplicantName  string                  `json:"applicant_name"`
	ApplicationDate string                  `json:"application_date"`
	EstimatedAmount float64                `json:"estimated_amount"`
	ActualAmount   float64                 `json:"actual_amount"`
	Status         string                  `json:"status"`
	Reason         string                  `json:"reason"`
	Items          []RequisitionItemResponse `json:"items"`
	Remarks        string                  `json:"remarks"`
	CreatedAt      string                  `json:"created_at"`
	UpdatedAt      string                  `json:"updated_at"`
}

// RequisitionItemResponse 采购申请项目响应
 type RequisitionItemResponse struct {
	ID           uint    `json:"id"`
	RequisitionID uint    `json:"requisition_id"`
	MaterialID   uint    `json:"material_id"`
	MaterialCode string  `json:"material_code"`
	MaterialName string  `json:"material_name"`
	Quantity     float64 `json:"quantity"`
	EstimatedPrice float64 `json:"estimated_price"`
	ActualPrice  float64 `json:"actual_price"`
	RequiredDate string  `json:"required_date"`
	Purpose      string  `json:"purpose"`
}

// 采购订单相关结构体

// PurchaseOrderCreateRequest 创建采购订单请求
 type PurchaseOrderCreateRequest struct {
	Code        string                     `json:"code" binding:"required"`
	SupplierID  uint                       `json:"supplier_id" binding:"required"`
	RequisitionID uint                     `json:"requisition_id"`
	OrderDate   string                     `json:"order_date" binding:"required"`
	ExpectedDeliveryDate string            `json:"expected_delivery_date" binding:"required"`
	TotalAmount float64                    `json:"total_amount" binding:"required"`
	TaxAmount   float64                    `json:"tax_amount"`
	GrandTotal  float64                    `json:"grand_total" binding:"required"`
	PaymentTerm string                     `json:"payment_term"`
	DeliveryAddress string                 `json:"delivery_address"`
	Items       []PurchaseOrderItemRequest `json:"items" binding:"required,dive"`
	Remarks     string                     `json:"remarks"`
}

// PurchaseOrderItemRequest 采购订单项目请求
 type PurchaseOrderItemRequest struct {
	MaterialID  uint    `json:"material_id" binding:"required"`
	Quantity    float64 `json:"quantity" binding:"required"`
	UnitPrice   float64 `json:"unit_price" binding:"required"`
	TaxRate     float64 `json:"tax_rate"`
	Subtotal    float64 `json:"subtotal" binding:"required"`
	DeliveryDate string `json:"delivery_date"`
}

// PurchaseOrderUpdateRequest 更新采购订单请求
 type PurchaseOrderUpdateRequest struct {
	SupplierID  uint                       `json:"supplier_id"`
	OrderDate   string                     `json:"order_date"`
	ExpectedDeliveryDate string            `json:"expected_delivery_date"`
	TotalAmount float64                    `json:"total_amount"`
	TaxAmount   float64                    `json:"tax_amount"`
	GrandTotal  float64                    `json:"grand_total"`
	PaymentTerm string                     `json:"payment_term"`
	DeliveryAddress string                 `json:"delivery_address"`
	Items       []PurchaseOrderItemRequest `json:"items" binding:"omitempty,dive"`
	Remarks     string                     `json:"remarks"`
}

// PurchaseOrderResponse 采购订单响应
 type PurchaseOrderResponse struct {
	ID                   uint                      `json:"id"`
	Code                 string                    `json:"code"`
	SupplierID           uint                      `json:"supplier_id"`
	SupplierName         string                    `json:"supplier_name"`
	RequisitionID        uint                      `json:"requisition_id"`
	OrderDate            string                    `json:"order_date"`
	ExpectedDeliveryDate string                    `json:"expected_delivery_date"`
	ActualDeliveryDate   string                    `json:"actual_delivery_date"`
	TotalAmount          float64                   `json:"total_amount"`
	TaxAmount            float64                   `json:"tax_amount"`
	GrandTotal           float64                   `json:"grand_total"`
	PaidAmount           float64                   `json:"paid_amount"`
	Status               string                    `json:"status"`
	PaymentTerm          string                    `json:"payment_term"`
	DeliveryAddress      string                    `json:"delivery_address"`
	Items                []PurchaseOrderItemResponse `json:"items"`
	Remarks              string                    `json:"remarks"`
	CreatedAt            string                    `json:"created_at"`
	UpdatedAt            string                    `json:"updated_at"`
}

// PurchaseOrderItemResponse 采购订单项目响应
 type PurchaseOrderItemResponse struct {
	ID            uint    `json:"id"`
	PurchaseOrderID uint    `json:"purchase_order_id"`
	MaterialID    uint    `json:"material_id"`
	MaterialCode  string  `json:"material_code"`
	MaterialName  string  `json:"material_name"`
	Quantity      float64 `json:"quantity"`
	ReceivedQuantity float64 `json:"received_quantity"`
	UnitPrice     float64 `json:"unit_price"`
	TaxRate       float64 `json:"tax_rate"`
	Subtotal      float64 `json:"subtotal"`
	DeliveryDate  string  `json:"delivery_date"`
}

// 采购收货相关结构体

// ReceiptCreateRequest 创建采购收货请求
 type ReceiptCreateRequest struct {
	Code           string                   `json:"code" binding:"required"`
	PurchaseOrderID uint                     `json:"purchase_order_id" binding:"required"`
	ReceiptDate    string                   `json:"receipt_date" binding:"required"`
	WarehouseID    uint                     `json:"warehouse_id" binding:"required"`
	TotalQuantity  float64                  `json:"total_quantity" binding:"required"`
	Items          []ReceiptItemRequest     `json:"items" binding:"required,dive"`
	Remarks        string                   `json:"remarks"`
}

// ReceiptItemRequest 采购收货项目请求
 type ReceiptItemRequest struct {
	PurchaseOrderItemID uint    `json:"purchase_order_item_id" binding:"required"`
	Quantity            float64 `json:"quantity" binding:"required"`
	BatchNumber         string  `json:"batch_number"`
	ExpiryDate          string  `json:"expiry_date"`
	Remarks             string  `json:"remarks"`
}

// ReceiptUpdateRequest 更新采购收货请求
 type ReceiptUpdateRequest struct {
	ReceiptDate    string                   `json:"receipt_date"`
	WarehouseID    uint                     `json:"warehouse_id"`
	TotalQuantity  float64                  `json:"total_quantity"`
	Items          []ReceiptItemRequest     `json:"items" binding:"omitempty,dive"`
	Remarks        string                   `json:"remarks"`
}

// ReceiptResponse 采购收货响应
 type ReceiptResponse struct {
	ID              uint                      `json:"id"`
	Code            string                    `json:"code"`
	PurchaseOrderID uint                      `json:"purchase_order_id"`
	PurchaseOrderCode string                  `json:"purchase_order_code"`
	SupplierName    string                    `json:"supplier_name"`
	ReceiptDate     string                    `json:"receipt_date"`
	WarehouseID     uint                      `json:"warehouse_id"`
	WarehouseName   string                    `json:"warehouse_name"`
	TotalQuantity   float64                   `json:"total_quantity"`
	Status          string                    `json:"status"`
	Items           []ReceiptItemResponse     `json:"items"`
	Remarks         string                    `json:"remarks"`
	CreatedAt       string                    `json:"created_at"`
	UpdatedAt       string                    `json:"updated_at"`
}

// ReceiptItemResponse 采购收货项目响应
 type ReceiptItemResponse struct {
	ID                  uint    `json:"id"`
	ReceiptID           uint    `json:"receipt_id"`
	PurchaseOrderItemID uint    `json:"purchase_order_item_id"`
	MaterialID          uint    `json:"material_id"`
	MaterialCode        string  `json:"material_code"`
	MaterialName        string  `json:"material_name"`
	Quantity            float64 `json:"quantity"`
	BatchNumber         string  `json:"batch_number"`
	ExpiryDate          string  `json:"expiry_date"`
	Remarks             string  `json:"remarks"`
}

// 采购发票相关结构体

// PurchaseInvoiceCreateRequest 创建采购发票请求
 type PurchaseInvoiceCreateRequest struct {
	Code           string                     `json:"code" binding:"required"`
	PurchaseOrderID uint                       `json:"purchase_order_id" binding:"required"`
	InvoiceDate    string                     `json:"invoice_date" binding:"required"`
	InvoiceNumber  string                     `json:"invoice_number" binding:"required"`
	TotalAmount    float64                    `json:"total_amount" binding:"required"`
	TaxAmount      float64                    `json:"tax_amount"`
	GrandTotal     float64                    `json:"grand_total" binding:"required"`
	DueDate        string                     `json:"due_date"`
	Items          []PurchaseInvoiceItemRequest `json:"items" binding:"required,dive"`
	Remarks        string                     `json:"remarks"`
}

// PurchaseInvoiceItemRequest 采购发票项目请求
 type PurchaseInvoiceItemRequest struct {
	MaterialID  uint    `json:"material_id" binding:"required"`
	Quantity    float64 `json:"quantity" binding:"required"`
	UnitPrice   float64 `json:"unit_price" binding:"required"`
	TaxRate     float64 `json:"tax_rate"`
	Subtotal    float64 `json:"subtotal" binding:"required"`
}

// PurchaseInvoiceUpdateRequest 更新采购发票请求
 type PurchaseInvoiceUpdateRequest struct {
	InvoiceDate    string                     `json:"invoice_date"`
	InvoiceNumber  string                     `json:"invoice_number"`
	TotalAmount    float64                    `json:"total_amount"`
	TaxAmount      float64                    `json:"tax_amount"`
	GrandTotal     float64                    `json:"grand_total"`
	DueDate        string                     `json:"due_date"`
	Items          []PurchaseInvoiceItemRequest `json:"items" binding:"omitempty,dive"`
	Remarks        string                     `json:"remarks"`
}

// PurchaseInvoiceResponse 采购发票响应
 type PurchaseInvoiceResponse struct {
	ID              uint                        `json:"id"`
	Code            string                      `json:"code"`
	PurchaseOrderID uint                        `json:"purchase_order_id"`
	PurchaseOrderCode string                    `json:"purchase_order_code"`
	SupplierID      uint                        `json:"supplier_id"`
	SupplierName    string                      `json:"supplier_name"`
	InvoiceDate     string                      `json:"invoice_date"`
	InvoiceNumber   string                      `json:"invoice_number"`
	TotalAmount     float64                     `json:"total_amount"`
	TaxAmount       float64                     `json:"tax_amount"`
	GrandTotal      float64                     `json:"grand_total"`
	PaidAmount      float64                     `json:"paid_amount"`
	DueDate         string                      `json:"due_date"`
	Status          string                      `json:"status"`
	Items           []PurchaseInvoiceItemResponse `json:"items"`
	Remarks         string                      `json:"remarks"`
	CreatedAt       string                      `json:"created_at"`
	UpdatedAt       string                      `json:"updated_at"`
}

// PurchaseInvoiceItemResponse 采购发票项目响应
 type PurchaseInvoiceItemResponse struct {
	ID            uint    `json:"id"`
	PurchaseInvoiceID uint    `json:"purchase_invoice_id"`
	MaterialID    uint    `json:"material_id"`
	MaterialCode  string  `json:"material_code"`
	MaterialName  string  `json:"material_name"`
	Quantity      float64 `json:"quantity"`
	UnitPrice     float64 `json:"unit_price"`
	TaxRate       float64 `json:"tax_rate"`
	Subtotal      float64 `json:"subtotal"`
}

// 采购退货相关结构体

// PurchaseReturnCreateRequest 创建采购退货请求
 type PurchaseReturnCreateRequest struct {
	Code           string                      `json:"code" binding:"required"`
	PurchaseOrderID uint                        `json:"purchase_order_id" binding:"required"`
	ReceiptID      uint                        `json:"receipt_id" binding:"required"`
	ReturnDate     string                      `json:"return_date" binding:"required"`
	TotalQuantity  float64                     `json:"total_quantity" binding:"required"`
	TotalAmount    float64                     `json:"total_amount" binding:"required"`
	Reason         string                      `json:"reason" binding:"required"`
	Items          []PurchaseReturnItemRequest `json:"items" binding:"required,dive"`
	Remarks        string                      `json:"remarks"`
}

// PurchaseReturnItemRequest 采购退货项目请求
 type PurchaseReturnItemRequest struct {
	ReceiptItemID uint    `json:"receipt_item_id" binding:"required"`
	Quantity      float64 `json:"quantity" binding:"required"`
	UnitPrice     float64 `json:"unit_price" binding:"required"`
	Subtotal      float64 `json:"subtotal" binding:"required"`
	Reason        string  `json:"reason"`
}

// PurchaseReturnUpdateRequest 更新采购退货请求
 type PurchaseReturnUpdateRequest struct {
	ReturnDate     string                      `json:"return_date"`
	TotalQuantity  float64                     `json:"total_quantity"`
	TotalAmount    float64                     `json:"total_amount"`
	Reason         string                      `json:"reason"`
	Items          []PurchaseReturnItemRequest `json:"items" binding:"omitempty,dive"`
	Remarks        string                      `json:"remarks"`
}

// PurchaseReturnResponse 采购退货响应
 type PurchaseReturnResponse struct {
	ID              uint                         `json:"id"`
	Code            string                       `json:"code"`
	PurchaseOrderID uint                         `json:"purchase_order_id"`
	PurchaseOrderCode string                     `json:"purchase_order_code"`
	ReceiptID       uint                         `json:"receipt_id"`
	ReceiptCode     string                       `json:"receipt_code"`
	SupplierID      uint                         `json:"supplier_id"`
	SupplierName    string                       `json:"supplier_name"`
	ReturnDate      string                       `json:"return_date"`
	TotalQuantity   float64                      `json:"total_quantity"`
	TotalAmount     float64                      `json:"total_amount"`
	Status          string                       `json:"status"`
	Reason          string                       `json:"reason"`
	Items           []PurchaseReturnItemResponse `json:"items"`
	Remarks         string                       `json:"remarks"`
	CreatedAt       string                       `json:"created_at"`
	UpdatedAt       string                       `json:"updated_at"`
}

// PurchaseReturnItemResponse 采购退货项目响应
 type PurchaseReturnItemResponse struct {
	ID             uint    `json:"id"`
	PurchaseReturnID uint    `json:"purchase_return_id"`
	ReceiptItemID  uint    `json:"receipt_item_id"`
	MaterialID     uint    `json:"material_id"`
	MaterialCode   string  `json:"material_code"`
	MaterialName   string  `json:"material_name"`
	Quantity       float64 `json:"quantity"`
	UnitPrice      float64 `json:"unit_price"`
	Subtotal       float64 `json:"subtotal"`
	Reason         string  `json:"reason"`
}

// 采购报表相关结构体

// PurchaseSummaryReportRequest 采购汇总报表请求
 type PurchaseSummaryReportRequest struct {
	StartDate     string `json:"start_date" binding:"required"`
	EndDate       string `json:"end_date" binding:"required"`
	SupplierID    uint   `json:"supplier_id"`
	Category      string `json:"category"`
}

// PurchaseSummaryReportResponse 采购汇总报表响应
 type PurchaseSummaryReportResponse struct {
	Period        string  `json:"period"`
	TotalOrders   int     `json:"total_orders"`
	TotalAmount   float64 `json:"total_amount"`
	TotalTax      float64 `json:"total_tax"`
	GrandTotal    float64 `json:"grand_total"`
	SupplierStats []SupplierStatResponse `json:"supplier_stats"`
}

// SupplierStatResponse 供应商统计响应
 type SupplierStatResponse struct {
	SupplierID   uint    `json:"supplier_id"`
	SupplierName string  `json:"supplier_name"`
	OrderCount   int     `json:"order_count"`
	TotalAmount  float64 `json:"total_amount"`
	Percentage   float64 `json:"percentage"`
}