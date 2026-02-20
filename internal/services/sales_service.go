package services

import (
	"errors"
	"time"

	"github.com/wu136995/ginx/internal/database"
	"github.com/wu136995/ginx/internal/models"
	"gorm.io/gorm"
)

// SalesService 销售服务接口
type SalesService interface {
	// 客户管理
	GetCustomerList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetCustomerDetail(id string) (map[string]interface{}, error)
	CreateCustomer(req map[string]interface{}) (map[string]interface{}, error)
	UpdateCustomer(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeleteCustomer(id string) error
	CustomerCreditEvaluate(id string) (map[string]interface{}, error)

	// 销售报价管理
	GetQuotationList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetQuotationDetail(id string) (map[string]interface{}, error)
	CreateQuotation(req map[string]interface{}) (map[string]interface{}, error)
	UpdateQuotation(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeleteQuotation(id string) error
	ApproveQuotation(id string) error
	GenerateOrderFromQuotation(id string) (map[string]interface{}, error)

	// 销售订单管理
	GetOrderList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetOrderDetail(id string) (map[string]interface{}, error)
	CreateOrder(req map[string]interface{}) (map[string]interface{}, error)
	UpdateOrder(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeleteOrder(id string) error
	ApproveOrder(id string) error
	CancelOrder(id string) error
	GenerateDeliveryFromOrder(id string) (map[string]interface{}, error)

	// 销售发货管理
	GetDeliveryList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetDeliveryDetail(id string) (map[string]interface{}, error)
	CreateDelivery(req map[string]interface{}) (map[string]interface{}, error)
	UpdateDelivery(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeleteDelivery(id string) error

	// 销售发票管理
	GetInvoiceList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetInvoiceDetail(id string) (map[string]interface{}, error)
	CreateInvoice(req map[string]interface{}) (map[string]interface{}, error)
	UpdateInvoice(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeleteInvoice(id string) error
	ReceiveInvoicePayment(id string, req map[string]interface{}) error

	// 销售退货管理
	GetReturnList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetReturnDetail(id string) (map[string]interface{}, error)
	CreateReturn(req map[string]interface{}) (map[string]interface{}, error)
	UpdateReturn(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeleteReturn(id string) error

	// 销售报表管理
	GetOrderExecutionReport(req map[string]interface{}) (map[string]interface{}, error)
	GetCustomerAnalysisReport(req map[string]interface{}) (map[string]interface{}, error)
	GetProductTrendReport(req map[string]interface{}) (map[string]interface{}, error)
	ExportSalesReport(req map[string]interface{}) ([]byte, error)
}

// salesService 销售服务实现
type salesService struct {
	db *gorm.DB
}

// NewSalesService 创建销售服务实例
func NewSalesService() SalesService {
	return &salesService{
		db: database.GetDB(),
	}
}

// 客户管理方法
func (s *salesService) GetCustomerList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取客户数据
	var customers []models.SalesCustomer
	result := s.db.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	customerList := make([]map[string]interface{}, len(customers))
	for i, customer := range customers {
		customerList[i] = map[string]interface{}{
			"id":             customer.ID,
			"customer_no":    customer.CustomerNo,
			"name":           customer.Name,
			"contact_person": customer.ContactPerson,
			"phone":          customer.Phone,
			"email":          customer.Email,
			"address":        customer.Address,
			"tax_no":         customer.TaxNo,
			"bank_name":      customer.BankName,
			"bank_account":   customer.BankAccount,
			"credit_limit":   customer.CreditLimit,
			"credit_balance": customer.CreditBalance,
			"payment_terms":  customer.PaymentTerms,
			"credit_days":    customer.CreditDays,
			"status":         customer.Status,
			"remarks":        customer.Remarks,
			"created_at":     customer.CreatedAt,
			"created_by":     customer.CreatedBy,
			"updated_at":     customer.UpdatedAt,
			"updated_by":     customer.UpdatedBy,
		}
	}

	return customerList, nil
}

func (s *salesService) GetCustomerDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取客户详情
	var customer models.SalesCustomer
	result := s.db.First(&customer, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	customerDetail := map[string]interface{}{
		"id":             customer.ID,
		"customer_no":    customer.CustomerNo,
		"name":           customer.Name,
		"contact_person": customer.ContactPerson,
		"phone":          customer.Phone,
		"email":          customer.Email,
		"address":        customer.Address,
		"tax_no":         customer.TaxNo,
		"bank_name":      customer.BankName,
		"bank_account":   customer.BankAccount,
		"credit_limit":   customer.CreditLimit,
		"credit_balance": customer.CreditBalance,
		"payment_terms":  customer.PaymentTerms,
		"credit_days":    customer.CreditDays,
		"status":         customer.Status,
		"remarks":        customer.Remarks,
		"created_at":     customer.CreatedAt,
		"created_by":     customer.CreatedBy,
		"updated_at":     customer.UpdatedAt,
		"updated_by":     customer.UpdatedBy,
	}

	return customerDetail, nil
}

func (s *salesService) CreateCustomer(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从请求中获取数据
	customer := models.SalesCustomer{
		ID:            req["id"].(string),
		CustomerNo:    req["customer_no"].(string),
		Name:          req["name"].(string),
		ContactPerson: req["contact_person"].(string),
		Phone:         req["phone"].(string),
		Email:         req["email"].(string),
		Address:       req["address"].(string),
		TaxNo:         req["tax_no"].(string),
		BankName:      req["bank_name"].(string),
		BankAccount:   req["bank_account"].(string),
		CreditLimit:   req["credit_limit"].(float64),
		CreditBalance: 0,
		PaymentTerms:  req["payment_terms"].(string),
		Status:        req["status"].(string),
		Remarks:       req["remarks"].(string),
		CreatedAt:     time.Now(),
		CreatedBy:     req["created_by"].(string),
		UpdatedAt:     time.Now(),
		UpdatedBy:     req["created_by"].(string),
	}

	// 保存到数据库
	result := s.db.Create(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	createdCustomer := map[string]interface{}{
		"id":             customer.ID,
		"customer_no":    customer.CustomerNo,
		"name":           customer.Name,
		"contact_person": customer.ContactPerson,
		"phone":          customer.Phone,
		"email":          customer.Email,
		"address":        customer.Address,
		"tax_no":         customer.TaxNo,
		"bank_name":      customer.BankName,
		"bank_account":   customer.BankAccount,
		"credit_limit":   customer.CreditLimit,
		"credit_balance": customer.CreditBalance,
		"payment_terms":  customer.PaymentTerms,
		"status":         customer.Status,
		"remarks":        customer.Remarks,
		"created_at":     customer.CreatedAt,
		"created_by":     customer.CreatedBy,
		"updated_at":     customer.UpdatedAt,
		"updated_by":     customer.UpdatedBy,
	}

	return createdCustomer, nil
}

func (s *salesService) UpdateCustomer(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取客户
	var customer models.SalesCustomer
	result := s.db.First(&customer, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	customer.Name = req["name"].(string)
	customer.ContactPerson = req["contact_person"].(string)
	customer.Phone = req["phone"].(string)
	customer.Email = req["email"].(string)
	customer.Address = req["address"].(string)
	customer.TaxNo = req["tax_no"].(string)
	customer.BankName = req["bank_name"].(string)
	customer.BankAccount = req["bank_account"].(string)
	customer.CreditLimit = req["credit_limit"].(float64)
	customer.PaymentTerms = req["payment_terms"].(string)
	customer.Status = req["status"].(string)
	customer.Remarks = req["remarks"].(string)
	customer.UpdatedAt = time.Now()
	customer.UpdatedBy = req["updated_by"].(string)

	// 保存到数据库
	result = s.db.Save(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	updatedCustomer := map[string]interface{}{
		"id":             customer.ID,
		"customer_no":    customer.CustomerNo,
		"name":           customer.Name,
		"contact_person": customer.ContactPerson,
		"phone":          customer.Phone,
		"email":          customer.Email,
		"address":        customer.Address,
		"tax_no":         customer.TaxNo,
		"bank_name":      customer.BankName,
		"bank_account":   customer.BankAccount,
		"credit_limit":   customer.CreditLimit,
		"credit_balance": customer.CreditBalance,
		"payment_terms":  customer.PaymentTerms,
		"credit_days":    customer.CreditDays,
		"status":         customer.Status,
		"remarks":        customer.Remarks,
		"created_at":     customer.CreatedAt,
		"created_by":     customer.CreatedBy,
		"updated_at":     customer.UpdatedAt,
		"updated_by":     customer.UpdatedBy,
	}

	return updatedCustomer, nil
}

func (s *salesService) DeleteCustomer(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除客户
	result := s.db.Delete(&models.SalesCustomer{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *salesService) CustomerCreditEvaluate(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取客户
	var customer models.SalesCustomer
	result := s.db.First(&customer, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 计算信用评估结果
	creditEvaluation := map[string]interface{}{
		"customer_id":      customer.ID,
		"customer_name":    customer.Name,
		"credit_limit":     customer.CreditLimit,
		"credit_balance":   customer.CreditBalance,
		"available_credit": customer.CreditLimit - customer.CreditBalance, // 假设模型中已添加 CreditBalance 字段，若未添加需同步修改模型
		"evaluation_date":  time.Now(),
	}

	return creditEvaluation, nil
}

// 销售报价管理方法
func (s *salesService) GetQuotationList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要从数据库读取销售报价数据
	// 由于模型中可能没有定义销售报价表，暂时返回空列表
	return []map[string]interface{}{}, nil
}

func (s *salesService) GetQuotationDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要从数据库读取销售报价详情
	// 由于模型中可能没有定义销售报价表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) CreateQuotation(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要创建销售报价
	// 由于模型中可能没有定义销售报价表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) UpdateQuotation(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要更新销售报价
	// 由于模型中可能没有定义销售报价表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) DeleteQuotation(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 实际应用中，这里需要删除销售报价
	// 由于模型中可能没有定义销售报价表，暂时返回nil
	return nil
}

func (s *salesService) ApproveQuotation(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 实际应用中，这里需要批准销售报价
	// 由于模型中可能没有定义销售报价表，暂时返回nil
	return nil
}

func (s *salesService) GenerateOrderFromQuotation(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要从销售报价生成销售订单
	// 由于模型中可能没有定义销售报价表和销售订单表，暂时返回空map
	return map[string]interface{}{}, nil
}

// 销售订单管理方法
func (s *salesService) GetOrderList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要从数据库读取销售订单数据
	// 由于模型中可能没有定义销售订单表，暂时返回空列表
	return []map[string]interface{}{}, nil
}

func (s *salesService) GetOrderDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要从数据库读取销售订单详情
	// 由于模型中可能没有定义销售订单表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) CreateOrder(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要创建销售订单
	// 由于模型中可能没有定义销售订单表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) UpdateOrder(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要更新销售订单
	// 由于模型中可能没有定义销售订单表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) DeleteOrder(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 实际应用中，这里需要删除销售订单
	// 由于模型中可能没有定义销售订单表，暂时返回nil
	return nil
}

func (s *salesService) ApproveOrder(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 实际应用中，这里需要批准销售订单
	// 由于模型中可能没有定义销售订单表，暂时返回nil
	return nil
}

func (s *salesService) CancelOrder(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 实际应用中，这里需要取消销售订单
	// 由于模型中可能没有定义销售订单表，暂时返回nil
	return nil
}

func (s *salesService) GenerateDeliveryFromOrder(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要从销售订单生成销售发货单
	// 由于模型中可能没有定义销售订单表和销售发货单表，暂时返回空map
	return map[string]interface{}{}, nil
}

// 销售发货管理方法
func (s *salesService) GetDeliveryList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要从数据库读取销售发货单数据
	// 由于模型中可能没有定义销售发货单表，暂时返回空列表
	return []map[string]interface{}{}, nil
}

func (s *salesService) GetDeliveryDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要从数据库读取销售发货单详情
	// 由于模型中可能没有定义销售发货单表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) CreateDelivery(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要创建销售发货单
	// 由于模型中可能没有定义销售发货单表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) UpdateDelivery(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要更新销售发货单
	// 由于模型中可能没有定义销售发货单表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) DeleteDelivery(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 实际应用中，这里需要删除销售发货单
	// 由于模型中可能没有定义销售发货单表，暂时返回nil
	return nil
}

// 销售发票管理方法
func (s *salesService) GetInvoiceList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要从数据库读取销售发票数据
	// 由于模型中可能没有定义销售发票表，暂时返回空列表
	return []map[string]interface{}{}, nil
}

func (s *salesService) GetInvoiceDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要从数据库读取销售发票详情
	// 由于模型中可能没有定义销售发票表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) CreateInvoice(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要创建销售发票
	// 由于模型中可能没有定义销售发票表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) UpdateInvoice(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要更新销售发票
	// 由于模型中可能没有定义销售发票表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) DeleteInvoice(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 实际应用中，这里需要删除销售发票
	// 由于模型中可能没有定义销售发票表，暂时返回nil
	return nil
}

func (s *salesService) ReceiveInvoicePayment(id string, req map[string]interface{}) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 实际应用中，这里需要处理销售发票收款
	// 由于模型中可能没有定义销售发票表，暂时返回nil
	return nil
}

// 销售退货管理方法
func (s *salesService) GetReturnList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要从数据库读取销售退货单数据
	// 由于模型中可能没有定义销售退货单表，暂时返回空列表
	return []map[string]interface{}{}, nil
}

func (s *salesService) GetReturnDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要从数据库读取销售退货单详情
	// 由于模型中可能没有定义销售退货单表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) CreateReturn(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要创建销售退货单
	// 由于模型中可能没有定义销售退货单表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) UpdateReturn(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要更新销售退货单
	// 由于模型中可能没有定义销售退货单表，暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) DeleteReturn(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 实际应用中，这里需要删除销售退货单
	// 由于模型中可能没有定义销售退货单表，暂时返回nil
	return nil
}

// 销售报表管理方法
func (s *salesService) GetOrderExecutionReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要基于数据库数据生成订单执行报表
	// 暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) GetCustomerAnalysisReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要基于数据库数据生成客户分析报表
	// 暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) GetProductTrendReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要基于数据库数据生成产品趋势报表
	// 暂时返回空map
	return map[string]interface{}{}, nil
}

func (s *salesService) ExportSalesReport(req map[string]interface{}) ([]byte, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要基于数据库数据导出销售报表
	// 暂时返回空字节数组
	return []byte{}, nil
}
