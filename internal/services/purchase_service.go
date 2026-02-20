package services

import (
	"errors"
	"time"

	"github.com/wu136995/ginx/internal/database"
	"github.com/wu136995/ginx/internal/models"
	"gorm.io/gorm"
)

// PurchaseService 采购服务接口
type PurchaseService interface {
	// 供应商管理
	GetSupplierList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetSupplierDetail(id string) (map[string]interface{}, error)
	CreateSupplier(req map[string]interface{}) (map[string]interface{}, error)
	UpdateSupplier(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeleteSupplier(id string) error
	GetSupplierContacts(id string) ([]map[string]interface{}, error)
	AddSupplierContact(id string, req map[string]interface{}) (map[string]interface{}, error)

	// 采购申请管理
	GetRequisitionList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetRequisitionDetail(id string) (map[string]interface{}, error)
	CreateRequisition(req map[string]interface{}) (map[string]interface{}, error)
	UpdateRequisition(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeleteRequisition(id string) error
	SubmitRequisition(id string) error
	ApproveRequisition(id string) error
	RejectRequisition(id string) error

	// 采购订单管理
	GetPurchaseOrderList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetPurchaseOrderDetail(id string) (map[string]interface{}, error)
	CreatePurchaseOrder(req map[string]interface{}) (map[string]interface{}, error)
	UpdatePurchaseOrder(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeletePurchaseOrder(id string) error
	SubmitPurchaseOrder(id string) error
	ApprovePurchaseOrder(id string) error
	RejectPurchaseOrder(id string) error
	ClosePurchaseOrder(id string) error

	// 采购收货管理
	GetPurchaseReceiptList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetPurchaseReceiptDetail(id string) (map[string]interface{}, error)
	CreatePurchaseReceipt(req map[string]interface{}) (map[string]interface{}, error)
	UpdatePurchaseReceipt(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeletePurchaseReceipt(id string) error
	CompletePurchaseReceipt(id string) error

	// 采购发票管理
	GetPurchaseInvoiceList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetPurchaseInvoiceDetail(id string) (map[string]interface{}, error)
	CreatePurchaseInvoice(req map[string]interface{}) (map[string]interface{}, error)
	UpdatePurchaseInvoice(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeletePurchaseInvoice(id string) error
	VerifyPurchaseInvoice(id string) error
	PayPurchaseInvoice(id string, req map[string]interface{}) error

	// 采购退货管理
	GetPurchaseReturnList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetPurchaseReturnDetail(id string) (map[string]interface{}, error)
	CreatePurchaseReturn(req map[string]interface{}) (map[string]interface{}, error)
	UpdatePurchaseReturn(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeletePurchaseReturn(id string) error
	CompletePurchaseReturn(id string) error

	// 采购报表管理
	GetPurchaseSummaryReport(req map[string]interface{}) (map[string]interface{}, error)
	GetPurchaseDetailReport(req map[string]interface{}) (map[string]interface{}, error)
	GetSupplierAnalysisReport(req map[string]interface{}) (map[string]interface{}, error)
	GetPriceAnalysisReport(req map[string]interface{}) (map[string]interface{}, error)
	GetPurchaseForecastReport(req map[string]interface{}) (map[string]interface{}, error)
	ExportPurchaseReport(req map[string]interface{}) ([]byte, error)
}

// purchaseService 采购服务实现
type purchaseService struct {
	db *gorm.DB
}

// NewPurchaseService 创建采购服务实例
func NewPurchaseService() PurchaseService {
	return &purchaseService{
		db: database.GetDB(),
	}
}

// 供应商管理方法
func (s *purchaseService) GetSupplierList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取供应商数据
	var vendors []models.PurchaseVendor
	result := s.db.Find(&vendors)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	supplierList := make([]map[string]interface{}, len(vendors))
	for i, vendor := range vendors {
		supplierList[i] = map[string]interface{}{
			"id":             vendor.ID,
			"vendor_no":      vendor.VendorNo,
			"name":           vendor.Name,
			"contact_person": vendor.ContactPerson,
			"phone":          vendor.Phone,
			"email":          vendor.Email,
			"address":        vendor.Address,
			"tax_no":         vendor.TaxNo,
			"bank_name":      vendor.BankName,
			"bank_account":   vendor.BankAccount,
			"category":       vendor.Category,
			"credit_limit":   vendor.CreditLimit,
			"lead_time":      vendor.LeadTime,
			"payment_terms":  vendor.PaymentTerms,
			"remarks":        vendor.Remarks,
			"status":         vendor.Status,
			"created_at":     vendor.CreatedAt,
			"created_by":     vendor.CreatedBy,
			"updated_at":     vendor.UpdatedAt,
			"updated_by":     vendor.UpdatedBy,
		}
	}

	return supplierList, nil
}

func (s *purchaseService) GetSupplierDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取供应商详情
	var vendor models.PurchaseVendor
	result := s.db.First(&vendor, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	supplierDetail := map[string]interface{}{
		"id":             vendor.ID,
		"vendor_no":      vendor.VendorNo,
		"name":           vendor.Name,
		"contact_person": vendor.ContactPerson,
		"phone":          vendor.Phone,
		"email":          vendor.Email,
		"address":        vendor.Address,
		"tax_no":         vendor.TaxNo,
		"bank_name":      vendor.BankName,
		"bank_account":   vendor.BankAccount,
		"category":       vendor.Category,
		"credit_limit":   vendor.CreditLimit,
		"lead_time":      vendor.LeadTime,
		"payment_terms":  vendor.PaymentTerms,
		"remarks":        vendor.Remarks,
		"status":         vendor.Status,
		"created_at":     vendor.CreatedAt,
		"created_by":     vendor.CreatedBy,
		"updated_at":     vendor.UpdatedAt,
		"updated_by":     vendor.UpdatedBy,
	}

	return supplierDetail, nil
}

func (s *purchaseService) CreateSupplier(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从请求中获取数据
	vendor := models.PurchaseVendor{
		ID:            req["id"].(string),
		VendorNo:      req["vendor_no"].(string),
		Name:          req["name"].(string),
		ContactPerson: req["contact_person"].(string),
		Phone:         req["phone"].(string),
		Email:         req["email"].(string),
		Address:       req["address"].(string),
		TaxNo:         req["tax_no"].(string),
		BankName:      req["bank_name"].(string),
		BankAccount:   req["bank_account"].(string),
		Category:      req["category"].(string),
		CreditLimit:   req["credit_limit"].(float64),
		LeadTime:      req["lead_time"].(int),
		PaymentTerms:  req["payment_terms"].(string),
		Remarks:       req["remarks"].(string),
		Status:        req["status"].(string),
		CreatedAt:     time.Now(),
		CreatedBy:     req["created_by"].(string),
		UpdatedAt:     time.Now(),
		UpdatedBy:     req["created_by"].(string),
	}

	// 保存到数据库
	result := s.db.Create(&vendor)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	createdSupplier := map[string]interface{}{
		"id":             vendor.ID,
		"vendor_no":      vendor.VendorNo,
		"name":           vendor.Name,
		"contact_person": vendor.ContactPerson,
		"phone":          vendor.Phone,
		"email":          vendor.Email,
		"address":        vendor.Address,
		"tax_no":         vendor.TaxNo,
		"bank_name":      vendor.BankName,
		"bank_account":   vendor.BankAccount,
		"category":       vendor.Category,
		"credit_limit":   vendor.CreditLimit,
		"lead_time":      vendor.LeadTime,
		"payment_terms":  vendor.PaymentTerms,
		"remarks":        vendor.Remarks,
		"status":         vendor.Status,
		"created_at":     vendor.CreatedAt,
		"created_by":     vendor.CreatedBy,
		"updated_at":     vendor.UpdatedAt,
		"updated_by":     vendor.UpdatedBy,
	}

	return createdSupplier, nil
}

func (s *purchaseService) UpdateSupplier(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取供应商
	var vendor models.PurchaseVendor
	result := s.db.First(&vendor, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	vendor.Name = req["name"].(string)
	vendor.ContactPerson = req["contact_person"].(string)
	vendor.Phone = req["phone"].(string)
	vendor.Email = req["email"].(string)
	vendor.Address = req["address"].(string)
	vendor.TaxNo = req["tax_no"].(string)
	vendor.BankName = req["bank_name"].(string)
	vendor.BankAccount = req["bank_account"].(string)
	vendor.Category = req["category"].(string)
	vendor.CreditLimit = req["credit_limit"].(float64)
	vendor.LeadTime = req["lead_time"].(int)
	vendor.PaymentTerms = req["payment_terms"].(string)
	vendor.Remarks = req["remarks"].(string)
	vendor.Status = req["status"].(string)
	vendor.UpdatedAt = time.Now()
	vendor.UpdatedBy = req["updated_by"].(string)

	// 保存到数据库
	result = s.db.Save(&vendor)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	updatedSupplier := map[string]interface{}{
		"id":             vendor.ID,
		"vendor_no":      vendor.VendorNo,
		"name":           vendor.Name,
		"contact_person": vendor.ContactPerson,
		"phone":          vendor.Phone,
		"email":          vendor.Email,
		"address":        vendor.Address,
		"tax_no":         vendor.TaxNo,
		"bank_name":      vendor.BankName,
		"bank_account":   vendor.BankAccount,
		"category":       vendor.Category,
		"credit_limit":   vendor.CreditLimit,
		"lead_time":      vendor.LeadTime,
		"payment_terms":  vendor.PaymentTerms,
		"remarks":        vendor.Remarks,
		"status":         vendor.Status,
		"created_at":     vendor.CreatedAt,
		"created_by":     vendor.CreatedBy,
		"updated_at":     vendor.UpdatedAt,
		"updated_by":     vendor.UpdatedBy,
	}

	return updatedSupplier, nil
}

func (s *purchaseService) DeleteSupplier(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除供应商
	result := s.db.Delete(&models.PurchaseVendor{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *purchaseService) GetSupplierContacts(id string) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要从数据库读取供应商联系人数据
	// 由于模型中没有定义联系人表，暂时返回空列表
	return []map[string]interface{}{}, nil
}

func (s *purchaseService) AddSupplierContact(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要创建供应商联系人
	// 由于模型中没有定义联系人表，暂时返回空map
	return map[string]interface{}{}, nil
}

// 采购申请管理方法
func (s *purchaseService) GetRequisitionList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要从数据库读取采购申请数据
	// 由于模型中没有定义采购申请表，暂时返回空列表
	return []map[string]interface{}{}, nil
}

func (s *purchaseService) GetRequisitionDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要从数据库读取采购申请详情
	// 由于模型中没有定义采购申请表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *purchaseService) CreateRequisition(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要创建采购申请
	// 由于模型中没有定义采购申请表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *purchaseService) UpdateRequisition(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要更新采购申请
	// 由于模型中没有定义采购申请表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *purchaseService) DeleteRequisition(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 实际应用中，这里需要删除采购申请
	// 由于模型中没有定义采购申请表，暂时返回nil
	return nil
}

func (s *purchaseService) SubmitRequisition(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 实际应用中，这里需要提交采购申请
	// 由于模型中没有定义采购申请表，暂时返回nil
	return nil
}

func (s *purchaseService) ApproveRequisition(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 实际应用中，这里需要批准采购申请
	// 由于模型中没有定义采购申请表，暂时返回nil
	return nil
}

func (s *purchaseService) RejectRequisition(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 实际应用中，这里需要拒绝采购申请
	// 由于模型中没有定义采购申请表，暂时返回nil
	return nil
}

// 采购订单管理方法
func (s *purchaseService) GetPurchaseOrderList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取采购订单数据
	var orders []models.PurchaseOrder
	result := s.db.Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	orderList := make([]map[string]interface{}, len(orders))
	for i, order := range orders {
		orderList[i] = map[string]interface{}{
			"id":            order.ID,
			"order_no":      order.OrderNo,
			"vendor_id":     order.VendorID,
			"order_date":    order.OrderDate,
			"delivery_date": order.DeliveryDate,
			"total_amount":  order.TotalAmount,
			"status":        order.Status,
			"payment_terms": order.PaymentTerms,
			"remarks":       order.Remarks,
			"created_at":    order.CreatedAt,
			"created_by":    order.CreatedBy,
			"updated_at":    order.UpdatedAt,
			"updated_by":    order.UpdatedBy,
		}
	}

	return orderList, nil
}

func (s *purchaseService) GetPurchaseOrderDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取采购订单详情
	var order models.PurchaseOrder
	result := s.db.First(&order, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	orderDetail := map[string]interface{}{
		"id":            order.ID,
		"order_no":      order.OrderNo,
		"vendor_id":     order.VendorID,
		"order_date":    order.OrderDate,
		"delivery_date": order.DeliveryDate,
		"total_amount":  order.TotalAmount,
		"status":        order.Status,
		"payment_terms": order.PaymentTerms,
		"remarks":       order.Remarks,
		"created_at":    order.CreatedAt,
		"created_by":    order.CreatedBy,
		"updated_at":    order.UpdatedAt,
		"updated_by":    order.UpdatedBy,
	}

	return orderDetail, nil
}

func (s *purchaseService) CreatePurchaseOrder(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从请求中获取数据
	order := models.PurchaseOrder{
		ID:           req["id"].(string),
		OrderNo:      req["order_no"].(string),
		VendorID:     req["vendor_id"].(string),
		OrderDate:    time.Now(),
		TotalAmount:  req["total_amount"].(float64),
		Status:       req["status"].(string),
		PaymentTerms: req["payment_terms"].(string),
		Remarks:      req["remarks"].(string),
		CreatedAt:    time.Now(),
		CreatedBy:    req["created_by"].(string),
		UpdatedAt:    time.Now(),
		UpdatedBy:    req["created_by"].(string),
	}

	// 处理可选的delivery_date
	if deliveryDate, ok := req["delivery_date"].(time.Time); ok {
		order.DeliveryDate = &deliveryDate
	}

	// 保存到数据库
	result := s.db.Create(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	createdOrder := map[string]interface{}{
		"id":            order.ID,
		"order_no":      order.OrderNo,
		"vendor_id":     order.VendorID,
		"order_date":    order.OrderDate,
		"delivery_date": order.DeliveryDate,
		"total_amount":  order.TotalAmount,
		"status":        order.Status,
		"payment_terms": order.PaymentTerms,
		"remarks":       order.Remarks,
		"created_at":    order.CreatedAt,
		"created_by":    order.CreatedBy,
		"updated_at":    order.UpdatedAt,
		"updated_by":    order.UpdatedBy,
	}

	return createdOrder, nil
}

func (s *purchaseService) UpdatePurchaseOrder(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取采购订单
	var order models.PurchaseOrder
	result := s.db.First(&order, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	order.VendorID = req["vendor_id"].(string)
	order.OrderDate = time.Now()
	order.TotalAmount = req["total_amount"].(float64)
	order.Status = req["status"].(string)
	order.PaymentTerms = req["payment_terms"].(string)
	order.Remarks = req["remarks"].(string)
	order.UpdatedAt = time.Now()
	order.UpdatedBy = req["updated_by"].(string)

	// 处理可选的delivery_date
	if deliveryDate, ok := req["delivery_date"].(time.Time); ok {
		order.DeliveryDate = &deliveryDate
	}

	// 保存到数据库
	result = s.db.Save(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	updatedOrder := map[string]interface{}{
		"id":            order.ID,
		"order_no":      order.OrderNo,
		"vendor_id":     order.VendorID,
		"order_date":    order.OrderDate,
		"delivery_date": order.DeliveryDate,
		"total_amount":  order.TotalAmount,
		"status":        order.Status,
		"payment_terms": order.PaymentTerms,
		"remarks":       order.Remarks,
		"created_at":    order.CreatedAt,
		"created_by":    order.CreatedBy,
		"updated_at":    order.UpdatedAt,
		"updated_by":    order.UpdatedBy,
	}

	return updatedOrder, nil
}

func (s *purchaseService) DeletePurchaseOrder(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除采购订单
	result := s.db.Delete(&models.PurchaseOrder{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *purchaseService) SubmitPurchaseOrder(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取采购订单
	var order models.PurchaseOrder
	result := s.db.First(&order, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为submitted
	order.Status = "submitted"
	order.UpdatedAt = time.Now()
	order.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *purchaseService) ApprovePurchaseOrder(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取采购订单
	var order models.PurchaseOrder
	result := s.db.First(&order, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为approved
	order.Status = "approved"
	order.UpdatedAt = time.Now()
	order.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *purchaseService) RejectPurchaseOrder(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取采购订单
	var order models.PurchaseOrder
	result := s.db.First(&order, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为rejected
	order.Status = "rejected"
	order.UpdatedAt = time.Now()
	order.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *purchaseService) ClosePurchaseOrder(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取采购订单
	var order models.PurchaseOrder
	result := s.db.First(&order, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为closed
	order.Status = "closed"
	order.UpdatedAt = time.Now()
	order.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 采购收货管理方法
func (s *purchaseService) GetPurchaseReceiptList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取采购收货单数据
	var receipts []models.PurchaseReceipt
	result := s.db.Find(&receipts)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	receiptList := make([]map[string]interface{}, len(receipts))
	for i, receipt := range receipts {
		receiptList[i] = map[string]interface{}{
			"id":             receipt.ID,
			"receipt_no":     receipt.ReceiptNo,
			"order_id":       receipt.OrderID,
			"vendor_id":      receipt.VendorID,
			"receipt_date":   receipt.ReceiptDate,
			"total_quantity": receipt.TotalQuantity,
			"total_amount":   receipt.TotalAmount,
			"status":         receipt.Status,
			"remarks":        receipt.Remarks,
			"created_at":     receipt.CreatedAt,
			"created_by":     receipt.CreatedBy,
			"updated_at":     receipt.UpdatedAt,
			"updated_by":     receipt.UpdatedBy,
		}
	}

	return receiptList, nil
}

func (s *purchaseService) GetPurchaseReceiptDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取采购收货单详情
	var receipt models.PurchaseReceipt
	result := s.db.First(&receipt, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	receiptDetail := map[string]interface{}{
		"id":             receipt.ID,
		"receipt_no":     receipt.ReceiptNo,
		"order_id":       receipt.OrderID,
		"vendor_id":      receipt.VendorID,
		"receipt_date":   receipt.ReceiptDate,
		"total_quantity": receipt.TotalQuantity,
		"total_amount":   receipt.TotalAmount,
		"status":         receipt.Status,
		"remarks":        receipt.Remarks,
		"created_at":     receipt.CreatedAt,
		"created_by":     receipt.CreatedBy,
		"updated_at":     receipt.UpdatedAt,
		"updated_by":     receipt.UpdatedBy,
	}

	return receiptDetail, nil
}

func (s *purchaseService) CreatePurchaseReceipt(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从请求中获取数据
	receipt := models.PurchaseReceipt{
		ID:            req["id"].(string),
		ReceiptNo:     req["receipt_no"].(string),
		OrderID:       req["order_id"].(string),
		VendorID:      req["vendor_id"].(string),
		ReceiptDate:   time.Now(),
		TotalQuantity: req["total_quantity"].(float64),
		TotalAmount:   req["total_amount"].(float64),
		Status:        req["status"].(string),
		Remarks:       req["remarks"].(string),
		CreatedAt:     time.Now(),
		CreatedBy:     req["created_by"].(string),
		UpdatedAt:     time.Now(),
		UpdatedBy:     req["created_by"].(string),
	}

	// 保存到数据库
	result := s.db.Create(&receipt)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	createdReceipt := map[string]interface{}{
		"id":             receipt.ID,
		"receipt_no":     receipt.ReceiptNo,
		"order_id":       receipt.OrderID,
		"vendor_id":      receipt.VendorID,
		"receipt_date":   receipt.ReceiptDate,
		"total_quantity": receipt.TotalQuantity,
		"total_amount":   receipt.TotalAmount,
		"status":         receipt.Status,
		"remarks":        receipt.Remarks,
		"created_at":     receipt.CreatedAt,
		"created_by":     receipt.CreatedBy,
		"updated_at":     receipt.UpdatedAt,
		"updated_by":     receipt.UpdatedBy,
	}

	return createdReceipt, nil
}

func (s *purchaseService) UpdatePurchaseReceipt(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取采购收货单
	var receipt models.PurchaseReceipt
	result := s.db.First(&receipt, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	receipt.OrderID = req["order_id"].(string)
	receipt.VendorID = req["vendor_id"].(string)
	receipt.ReceiptDate = time.Now()
	receipt.TotalQuantity = req["total_quantity"].(float64)
	receipt.TotalAmount = req["total_amount"].(float64)
	receipt.Status = req["status"].(string)
	receipt.Remarks = req["remarks"].(string)
	receipt.UpdatedAt = time.Now()
	receipt.UpdatedBy = req["updated_by"].(string)

	// 保存到数据库
	result = s.db.Save(&receipt)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	updatedReceipt := map[string]interface{}{
		"id":             receipt.ID,
		"receipt_no":     receipt.ReceiptNo,
		"order_id":       receipt.OrderID,
		"vendor_id":      receipt.VendorID,
		"receipt_date":   receipt.ReceiptDate,
		"total_quantity": receipt.TotalQuantity,
		"total_amount":   receipt.TotalAmount,
		"status":         receipt.Status,
		"remarks":        receipt.Remarks,
		"created_at":     receipt.CreatedAt,
		"created_by":     receipt.CreatedBy,
		"updated_at":     receipt.UpdatedAt,
		"updated_by":     receipt.UpdatedBy,
	}

	return updatedReceipt, nil
}

func (s *purchaseService) DeletePurchaseReceipt(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除采购收货单
	result := s.db.Delete(&models.PurchaseReceipt{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *purchaseService) CompletePurchaseReceipt(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取采购收货单
	var receipt models.PurchaseReceipt
	result := s.db.First(&receipt, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为completed
	receipt.Status = "completed"
	receipt.UpdatedAt = time.Now()
	receipt.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&receipt)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 采购发票管理方法
func (s *purchaseService) GetPurchaseInvoiceList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取采购发票数据
	var invoices []models.PurchaseInvoice
	result := s.db.Find(&invoices)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	invoiceList := make([]map[string]interface{}, len(invoices))
	for i, invoice := range invoices {
		invoiceList[i] = map[string]interface{}{
			"id":            invoice.ID,
			"invoice_no":    invoice.InvoiceNo,
			"order_id":      invoice.OrderID,
			"receipt_id":    invoice.ReceiptID,
			"vendor_id":     invoice.VendorID,
			"invoice_date":  invoice.InvoiceDate,
			"due_date":      invoice.DueDate,
			"total_amount":  invoice.TotalAmount,
			"tax_amount":    invoice.TaxAmount,
			"status":        invoice.Status,
			"payment_terms": invoice.PaymentTerms,
			"remarks":       invoice.Remarks,
			"created_at":    invoice.CreatedAt,
			"created_by":    invoice.CreatedBy,
			"updated_at":    invoice.UpdatedAt,
			"updated_by":    invoice.UpdatedBy,
		}
	}

	return invoiceList, nil
}

func (s *purchaseService) GetPurchaseInvoiceDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取采购发票详情
	var invoice models.PurchaseInvoice
	result := s.db.First(&invoice, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	invoiceDetail := map[string]interface{}{
		"id":            invoice.ID,
		"invoice_no":    invoice.InvoiceNo,
		"order_id":      invoice.OrderID,
		"receipt_id":    invoice.ReceiptID,
		"vendor_id":     invoice.VendorID,
		"invoice_date":  invoice.InvoiceDate,
		"due_date":      invoice.DueDate,
		"total_amount":  invoice.TotalAmount,
		"tax_amount":    invoice.TaxAmount,
		"status":        invoice.Status,
		"payment_terms": invoice.PaymentTerms,
		"remarks":       invoice.Remarks,
		"created_at":    invoice.CreatedAt,
		"created_by":    invoice.CreatedBy,
		"updated_at":    invoice.UpdatedAt,
		"updated_by":    invoice.UpdatedBy,
	}

	return invoiceDetail, nil
}

func (s *purchaseService) CreatePurchaseInvoice(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从请求中获取数据
	invoice := models.PurchaseInvoice{
		ID:           req["id"].(string),
		InvoiceNo:    req["invoice_no"].(string),
		OrderID:      req["order_id"].(string),
		ReceiptID:    req["receipt_id"].(string),
		VendorID:     req["vendor_id"].(string),
		InvoiceDate:  time.Now(),
		DueDate:      time.Now().AddDate(0, 1, 0), // 默认30天
		TotalAmount:  req["total_amount"].(float64),
		TaxAmount:    req["tax_amount"].(float64),
		Status:       req["status"].(string),
		PaymentTerms: req["payment_terms"].(string),
		Remarks:      req["remarks"].(string),
		CreatedAt:    time.Now(),
		CreatedBy:    req["created_by"].(string),
		UpdatedAt:    time.Now(),
		UpdatedBy:    req["created_by"].(string),
	}

	// 保存到数据库
	result := s.db.Create(&invoice)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	createdInvoice := map[string]interface{}{
		"id":            invoice.ID,
		"invoice_no":    invoice.InvoiceNo,
		"order_id":      invoice.OrderID,
		"receipt_id":    invoice.ReceiptID,
		"vendor_id":     invoice.VendorID,
		"invoice_date":  invoice.InvoiceDate,
		"due_date":      invoice.DueDate,
		"total_amount":  invoice.TotalAmount,
		"tax_amount":    invoice.TaxAmount,
		"status":        invoice.Status,
		"payment_terms": invoice.PaymentTerms,
		"remarks":       invoice.Remarks,
		"created_at":    invoice.CreatedAt,
		"created_by":    invoice.CreatedBy,
		"updated_at":    invoice.UpdatedAt,
		"updated_by":    invoice.UpdatedBy,
	}

	return createdInvoice, nil
}

func (s *purchaseService) UpdatePurchaseInvoice(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取采购发票
	var invoice models.PurchaseInvoice
	result := s.db.First(&invoice, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	invoice.OrderID = req["order_id"].(string)
	invoice.ReceiptID = req["receipt_id"].(string)
	invoice.VendorID = req["vendor_id"].(string)
	invoice.InvoiceDate = time.Now()
	invoice.DueDate = time.Now().AddDate(0, 1, 0) // 默认30天
	invoice.TotalAmount = req["total_amount"].(float64)
	invoice.TaxAmount = req["tax_amount"].(float64)
	invoice.Status = req["status"].(string)
	invoice.PaymentTerms = req["payment_terms"].(string)
	invoice.Remarks = req["remarks"].(string)
	invoice.UpdatedAt = time.Now()
	invoice.UpdatedBy = req["updated_by"].(string)

	// 保存到数据库
	result = s.db.Save(&invoice)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	updatedInvoice := map[string]interface{}{
		"id":            invoice.ID,
		"invoice_no":    invoice.InvoiceNo,
		"order_id":      invoice.OrderID,
		"receipt_id":    invoice.ReceiptID,
		"vendor_id":     invoice.VendorID,
		"invoice_date":  invoice.InvoiceDate,
		"due_date":      invoice.DueDate,
		"total_amount":  invoice.TotalAmount,
		"tax_amount":    invoice.TaxAmount,
		"status":        invoice.Status,
		"payment_terms": invoice.PaymentTerms,
		"remarks":       invoice.Remarks,
		"created_at":    invoice.CreatedAt,
		"created_by":    invoice.CreatedBy,
		"updated_at":    invoice.UpdatedAt,
		"updated_by":    invoice.UpdatedBy,
	}

	return updatedInvoice, nil
}

func (s *purchaseService) DeletePurchaseInvoice(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除采购发票
	result := s.db.Delete(&models.PurchaseInvoice{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *purchaseService) VerifyPurchaseInvoice(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取采购发票
	var invoice models.PurchaseInvoice
	result := s.db.First(&invoice, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为verified
	invoice.Status = "verified"
	invoice.UpdatedAt = time.Now()
	invoice.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&invoice)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *purchaseService) PayPurchaseInvoice(id string, req map[string]interface{}) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取采购发票
	var invoice models.PurchaseInvoice
	result := s.db.First(&invoice, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为paid
	invoice.Status = "paid"
	invoice.UpdatedAt = time.Now()
	invoice.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&invoice)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 采购退货管理方法
func (s *purchaseService) GetPurchaseReturnList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取采购退货单数据
	var returns []models.PurchaseReturn
	result := s.db.Find(&returns)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	returnList := make([]map[string]interface{}, len(returns))
	for i, returnOrder := range returns {
		returnList[i] = map[string]interface{}{
			"id":             returnOrder.ID,
			"return_no":      returnOrder.ReturnNo,
			"order_id":       returnOrder.OrderID,
			"receipt_id":     returnOrder.ReceiptID,
			"vendor_id":      returnOrder.VendorID,
			"return_date":    returnOrder.ReturnDate,
			"total_quantity": returnOrder.TotalQuantity,
			"total_amount":   returnOrder.TotalAmount,
			"status":         returnOrder.Status,
			"reason":         returnOrder.Reason,
			"remarks":        returnOrder.Remarks,
			"created_at":     returnOrder.CreatedAt,
			"created_by":     returnOrder.CreatedBy,
			"updated_at":     returnOrder.UpdatedAt,
			"updated_by":     returnOrder.UpdatedBy,
		}
	}

	return returnList, nil
}

func (s *purchaseService) GetPurchaseReturnDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取采购退货单详情
	var returnOrder models.PurchaseReturn
	result := s.db.First(&returnOrder, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	returnDetail := map[string]interface{}{
		"id":             returnOrder.ID,
		"return_no":      returnOrder.ReturnNo,
		"order_id":       returnOrder.OrderID,
		"receipt_id":     returnOrder.ReceiptID,
		"vendor_id":      returnOrder.VendorID,
		"return_date":    returnOrder.ReturnDate,
		"total_quantity": returnOrder.TotalQuantity,
		"total_amount":   returnOrder.TotalAmount,
		"status":         returnOrder.Status,
		"reason":         returnOrder.Reason,
		"remarks":        returnOrder.Remarks,
		"created_at":     returnOrder.CreatedAt,
		"created_by":     returnOrder.CreatedBy,
		"updated_at":     returnOrder.UpdatedAt,
		"updated_by":     returnOrder.UpdatedBy,
	}

	return returnDetail, nil
}

func (s *purchaseService) CreatePurchaseReturn(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从请求中获取数据
	returnOrder := models.PurchaseReturn{
		ID:            req["id"].(string),
		ReturnNo:      req["return_no"].(string),
		OrderID:       req["order_id"].(string),
		ReceiptID:     req["receipt_id"].(string),
		VendorID:      req["vendor_id"].(string),
		ReturnDate:    time.Now(),
		TotalQuantity: req["total_quantity"].(float64),
		TotalAmount:   req["total_amount"].(float64),
		Status:        req["status"].(string),
		Reason:        req["reason"].(string),
		Remarks:       req["remarks"].(string),
		CreatedAt:     time.Now(),
		CreatedBy:     req["created_by"].(string),
		UpdatedAt:     time.Now(),
		UpdatedBy:     req["created_by"].(string),
	}

	// 保存到数据库
	result := s.db.Create(&returnOrder)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	createdReturn := map[string]interface{}{
		"id":             returnOrder.ID,
		"return_no":      returnOrder.ReturnNo,
		"order_id":       returnOrder.OrderID,
		"receipt_id":     returnOrder.ReceiptID,
		"vendor_id":      returnOrder.VendorID,
		"return_date":    returnOrder.ReturnDate,
		"total_quantity": returnOrder.TotalQuantity,
		"total_amount":   returnOrder.TotalAmount,
		"status":         returnOrder.Status,
		"reason":         returnOrder.Reason,
		"remarks":        returnOrder.Remarks,
		"created_at":     returnOrder.CreatedAt,
		"created_by":     returnOrder.CreatedBy,
		"updated_at":     returnOrder.UpdatedAt,
		"updated_by":     returnOrder.UpdatedBy,
	}

	return createdReturn, nil
}

func (s *purchaseService) UpdatePurchaseReturn(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取采购退货单
	var returnOrder models.PurchaseReturn
	result := s.db.First(&returnOrder, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	returnOrder.OrderID = req["order_id"].(string)
	returnOrder.ReceiptID = req["receipt_id"].(string)
	returnOrder.VendorID = req["vendor_id"].(string)
	returnOrder.ReturnDate = time.Now()
	returnOrder.TotalQuantity = req["total_quantity"].(float64)
	returnOrder.TotalAmount = req["total_amount"].(float64)
	returnOrder.Status = req["status"].(string)
	returnOrder.Reason = req["reason"].(string)
	returnOrder.Remarks = req["remarks"].(string)
	returnOrder.UpdatedAt = time.Now()
	returnOrder.UpdatedBy = req["updated_by"].(string)

	// 保存到数据库
	result = s.db.Save(&returnOrder)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	updatedReturn := map[string]interface{}{
		"id":             returnOrder.ID,
		"return_no":      returnOrder.ReturnNo,
		"order_id":       returnOrder.OrderID,
		"receipt_id":     returnOrder.ReceiptID,
		"vendor_id":      returnOrder.VendorID,
		"return_date":    returnOrder.ReturnDate,
		"total_quantity": returnOrder.TotalQuantity,
		"total_amount":   returnOrder.TotalAmount,
		"status":         returnOrder.Status,
		"reason":         returnOrder.Reason,
		"remarks":        returnOrder.Remarks,
		"created_at":     returnOrder.CreatedAt,
		"created_by":     returnOrder.CreatedBy,
		"updated_at":     returnOrder.UpdatedAt,
		"updated_by":     returnOrder.UpdatedBy,
	}

	return updatedReturn, nil
}

func (s *purchaseService) DeletePurchaseReturn(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除采购退货单
	result := s.db.Delete(&models.PurchaseReturn{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *purchaseService) CompletePurchaseReturn(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取采购退货单
	var returnOrder models.PurchaseReturn
	result := s.db.First(&returnOrder, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为completed
	returnOrder.Status = "completed"
	returnOrder.UpdatedAt = time.Now()
	returnOrder.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&returnOrder)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 采购报表管理方法
func (s *purchaseService) GetPurchaseSummaryReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要基于数据库数据生成采购汇总报表
	// 暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *purchaseService) GetPurchaseDetailReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要基于数据库数据生成采购明细报表
	// 暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *purchaseService) GetSupplierAnalysisReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要基于数据库数据生成供应商分析报表
	// 暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *purchaseService) GetPriceAnalysisReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要基于数据库数据生成价格分析报表
	// 暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *purchaseService) GetPurchaseForecastReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要基于数据库数据生成采购预测报表
	// 暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *purchaseService) ExportPurchaseReport(req map[string]interface{}) ([]byte, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要基于数据库数据导出采购报表
	// 暂时返回空字节数组
	return []byte{}, nil
}
