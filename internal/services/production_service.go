package services

import (
	"errors"
	"time"

	"github.com/wu136995/ginx/internal/database"
	"github.com/wu136995/ginx/internal/models"
	"gorm.io/gorm"
)

// ProductionService 生产服务接口
type ProductionService interface {
	// 生产订单管理
	GetProductionOrderList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetProductionOrderDetail(id string) (map[string]interface{}, error)
	CreateProductionOrder(req map[string]interface{}) (map[string]interface{}, error)
	UpdateProductionOrder(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeleteProductionOrder(id string) error
	SubmitProductionOrder(id string) error
	ApproveProductionOrder(id string) error
	RejectProductionOrder(id string) error
	StartProductionOrder(id string) error
	CompleteProductionOrder(id string) error
	CancelProductionOrder(id string) error

	// 生产工单管理
	GetProductionTicketList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetProductionTicketDetail(id string) (map[string]interface{}, error)
	CreateProductionTicket(req map[string]interface{}) (map[string]interface{}, error)
	UpdateProductionTicket(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeleteProductionTicket(id string) error
	StartProductionTicket(id string) error
	CompleteProductionTicket(id string) error
	CancelProductionTicket(id string) error

	// 工艺路线管理
	GetRoutingList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetRoutingDetail(id string) (map[string]interface{}, error)
	CreateRouting(req map[string]interface{}) (map[string]interface{}, error)
	UpdateRouting(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeleteRouting(id string) error

	// 工作中心管理
	GetWorkCenterList(req map[string]interface{}) ([]map[string]interface{}, error)
	GetWorkCenterDetail(id string) (map[string]interface{}, error)
	CreateWorkCenter(req map[string]interface{}) (map[string]interface{}, error)
	UpdateWorkCenter(id string, req map[string]interface{}) (map[string]interface{}, error)
	DeleteWorkCenter(id string) error
	GetWorkCenterCapacity(id string) (map[string]interface{}, error)
	GetWorkCenterSchedule(id string, req map[string]interface{}) (map[string]interface{}, error)

	// 物料需求计划管理
	RunMRP(req map[string]interface{}) (map[string]interface{}, error)
	GetMRPResults(req map[string]interface{}) ([]map[string]interface{}, error)
	GetMRPSuggestions(req map[string]interface{}) ([]map[string]interface{}, error)

	// 生产报表管理
	GetProductionPlanReport(req map[string]interface{}) (map[string]interface{}, error)
	GetProductionExecutionReport(req map[string]interface{}) (map[string]interface{}, error)
	GetWorkCenterLoadReport(req map[string]interface{}) (map[string]interface{}, error)
	GetMaterialConsumptionReport(req map[string]interface{}) (map[string]interface{}, error)
	GetProductionCostReport(req map[string]interface{}) (map[string]interface{}, error)
	ExportProductionReport(req map[string]interface{}) ([]byte, error)
}

// productionService 生产服务实现
type productionService struct {
	db *gorm.DB
}

// NewProductionService 创建生产服务实例
func NewProductionService() ProductionService {
	return &productionService{
		db: database.GetDB(),
	}
}

// 生产订单管理方法
func (s *productionService) GetProductionOrderList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取生产订单数据
	var productionOrders []models.ProductionOrder
	result := s.db.Find(&productionOrders)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	orderList := make([]map[string]interface{}, len(productionOrders))
	for i, order := range productionOrders {
		orderList[i] = map[string]interface{}{
			"id":           order.ID,
			"order_no":     order.OrderNo,
			"product_name": order.ProductName,
			"quantity":     order.Quantity,
			"status":       order.Status,
			"priority":     order.Priority,
			"start_date":   order.StartDate,
			"end_date":     order.EndDate,
			"created_at":   order.CreatedAt,
			"created_by":   order.CreatedBy,
			"updated_at":   order.UpdatedAt,
			"updated_by":   order.UpdatedBy,
		}
	}

	return orderList, nil
}

func (s *productionService) GetProductionOrderDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取生产订单详情
	var productionOrder models.ProductionOrder
	result := s.db.First(&productionOrder, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	orderDetail := map[string]interface{}{
		"id":           productionOrder.ID,
		"order_no":     productionOrder.OrderNo,
		"product_name": productionOrder.ProductName,
		"quantity":     productionOrder.Quantity,
		"status":       productionOrder.Status,
		"priority":     productionOrder.Priority,
		"start_date":   productionOrder.StartDate,
		"end_date":     productionOrder.EndDate,
		"created_at":   productionOrder.CreatedAt,
		"created_by":   productionOrder.CreatedBy,
		"updated_at":   productionOrder.UpdatedAt,
		"updated_by":   productionOrder.UpdatedBy,
	}

	return orderDetail, nil
}

func (s *productionService) CreateProductionOrder(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从请求中获取数据
	productionOrder := models.ProductionOrder{
		ID:          req["id"].(string),
		OrderNo:     req["order_no"].(string),
		ProductName: req["product_name"].(string),
		Quantity:    req["quantity"].(int),
		Status:      req["status"].(string),
		Priority:    req["priority"].(string),
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(7 * 24 * time.Hour),
		CreatedAt:   time.Now(),
		CreatedBy:   req["created_by"].(string),
		UpdatedAt:   time.Now(),
		UpdatedBy:   req["created_by"].(string),
	}

	// 保存到数据库
	result := s.db.Create(&productionOrder)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	createdOrder := map[string]interface{}{
		"id":           productionOrder.ID,
		"order_no":     productionOrder.OrderNo,
		"product_name": productionOrder.ProductName,
		"quantity":     productionOrder.Quantity,
		"status":       productionOrder.Status,
		"priority":     productionOrder.Priority,
		"start_date":   productionOrder.StartDate,
		"end_date":     productionOrder.EndDate,
		"created_at":   productionOrder.CreatedAt,
		"created_by":   productionOrder.CreatedBy,
		"updated_at":   productionOrder.UpdatedAt,
		"updated_by":   productionOrder.UpdatedBy,
	}

	return createdOrder, nil
}

func (s *productionService) UpdateProductionOrder(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取生产订单
	var productionOrder models.ProductionOrder
	result := s.db.First(&productionOrder, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	productionOrder.ProductName = req["product_name"].(string)
	productionOrder.Quantity = req["quantity"].(int)
	productionOrder.Status = req["status"].(string)
	productionOrder.Priority = req["priority"].(string)
	productionOrder.StartDate = time.Now()
	productionOrder.EndDate = time.Now().Add(7 * 24 * time.Hour)
	productionOrder.UpdatedAt = time.Now()
	productionOrder.UpdatedBy = req["updated_by"].(string)

	// 保存到数据库
	result = s.db.Save(&productionOrder)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	updatedOrder := map[string]interface{}{
		"id":           productionOrder.ID,
		"order_no":     productionOrder.OrderNo,
		"product_name": productionOrder.ProductName,
		"quantity":     productionOrder.Quantity,
		"status":       productionOrder.Status,
		"priority":     productionOrder.Priority,
		"start_date":   productionOrder.StartDate,
		"end_date":     productionOrder.EndDate,
		"created_at":   productionOrder.CreatedAt,
		"created_by":   productionOrder.CreatedBy,
		"updated_at":   productionOrder.UpdatedAt,
		"updated_by":   productionOrder.UpdatedBy,
	}

	return updatedOrder, nil
}

func (s *productionService) DeleteProductionOrder(id string) error {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return errors.New("database connection is nil")
	}

	// 从数据库删除生产订单
	result := s.db.Delete(&models.ProductionOrder{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *productionService) SubmitProductionOrder(id string) error {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return errors.New("database connection is nil")
	}

	// 从数据库读取生产订单
	var productionOrder models.ProductionOrder
	result := s.db.First(&productionOrder, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为submitted
	productionOrder.Status = "submitted"
	productionOrder.UpdatedAt = time.Now()
	productionOrder.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&productionOrder)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *productionService) ApproveProductionOrder(id string) error {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return errors.New("database connection is nil")
	}

	// 从数据库读取生产订单
	var productionOrder models.ProductionOrder
	result := s.db.First(&productionOrder, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为approved
	productionOrder.Status = "approved"
	productionOrder.UpdatedAt = time.Now()
	productionOrder.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&productionOrder)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *productionService) RejectProductionOrder(id string) error {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return errors.New("database connection is nil")
	}

	// 从数据库读取生产订单
	var productionOrder models.ProductionOrder
	result := s.db.First(&productionOrder, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为rejected
	productionOrder.Status = "rejected"
	productionOrder.UpdatedAt = time.Now()
	productionOrder.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&productionOrder)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *productionService) StartProductionOrder(id string) error {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return errors.New("database connection is nil")
	}

	// 从数据库读取生产订单
	var productionOrder models.ProductionOrder
	result := s.db.First(&productionOrder, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为in_progress
	productionOrder.Status = "in_progress"
	productionOrder.UpdatedAt = time.Now()
	productionOrder.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&productionOrder)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *productionService) CompleteProductionOrder(id string) error {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return errors.New("database connection is nil")
	}

	// 从数据库读取生产订单
	var productionOrder models.ProductionOrder
	result := s.db.First(&productionOrder, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为completed
	productionOrder.Status = "completed"
	productionOrder.UpdatedAt = time.Now()
	productionOrder.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&productionOrder)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *productionService) CancelProductionOrder(id string) error {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return errors.New("database connection is nil")
	}

	// 从数据库读取生产订单
	var productionOrder models.ProductionOrder
	result := s.db.First(&productionOrder, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为cancelled
	productionOrder.Status = "cancelled"
	productionOrder.UpdatedAt = time.Now()
	productionOrder.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&productionOrder)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 生产工单管理方法
func (s *productionService) GetProductionTicketList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取生产工单数据
	var productionTickets []models.ProductionTicket
	result := s.db.Find(&productionTickets)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	ticketList := make([]map[string]interface{}, len(productionTickets))
	for i, ticket := range productionTickets {
		ticketList[i] = map[string]interface{}{
			"id":                  ticket.ID,
			"ticket_no":           ticket.TicketNo,
			"production_order_id": ticket.ProductionOrderID,
			"product_name":        ticket.ProductName,
			"quantity":            ticket.Quantity,
			"status":              ticket.Status,
			"work_center_id":      ticket.WorkCenterID,
			"start_time":          ticket.StartTime,
			"end_time":            ticket.EndTime,
			"created_at":          ticket.CreatedAt,
			"created_by":          ticket.CreatedBy,
			"updated_at":          ticket.UpdatedAt,
			"updated_by":          ticket.UpdatedBy,
		}
	}

	return ticketList, nil
}

func (s *productionService) GetProductionTicketDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取生产工单详情
	var productionTicket models.ProductionTicket
	result := s.db.First(&productionTicket, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	ticketDetail := map[string]interface{}{
		"id":                  productionTicket.ID,
		"ticket_no":           productionTicket.TicketNo,
		"production_order_id": productionTicket.ProductionOrderID,
		"product_name":        productionTicket.ProductName,
		"quantity":            productionTicket.Quantity,
		"status":              productionTicket.Status,
		"work_center_id":      productionTicket.WorkCenterID,
		"start_time":          productionTicket.StartTime,
		"end_time":            productionTicket.EndTime,
		"created_at":          productionTicket.CreatedAt,
		"created_by":          productionTicket.CreatedBy,
		"updated_at":          productionTicket.UpdatedAt,
		"updated_by":          productionTicket.UpdatedBy,
	}

	return ticketDetail, nil
}

func (s *productionService) CreateProductionTicket(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从请求中获取数据
	productionTicket := models.ProductionTicket{
		ID:                req["id"].(string),
		TicketNo:          req["ticket_no"].(string),
		ProductionOrderID: req["production_order_id"].(string),
		ProductName:       req["product_name"].(string),
		Quantity:          req["quantity"].(int),
		Status:            req["status"].(string),
		WorkCenterID:      req["work_center_id"].(string),
		StartTime:         time.Now(),
		EndTime:           time.Now().Add(24 * time.Hour),
		CreatedAt:         time.Now(),
		CreatedBy:         req["created_by"].(string),
		UpdatedAt:         time.Now(),
		UpdatedBy:         req["created_by"].(string),
	}

	// 保存到数据库
	result := s.db.Create(&productionTicket)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	createdTicket := map[string]interface{}{
		"id":                  productionTicket.ID,
		"ticket_no":           productionTicket.TicketNo,
		"production_order_id": productionTicket.ProductionOrderID,
		"product_name":        productionTicket.ProductName,
		"quantity":            productionTicket.Quantity,
		"status":              productionTicket.Status,
		"work_center_id":      productionTicket.WorkCenterID,
		"start_time":          productionTicket.StartTime,
		"end_time":            productionTicket.EndTime,
		"created_at":          productionTicket.CreatedAt,
		"created_by":          productionTicket.CreatedBy,
		"updated_at":          productionTicket.UpdatedAt,
		"updated_by":          productionTicket.UpdatedBy,
	}

	return createdTicket, nil
}

func (s *productionService) UpdateProductionTicket(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取生产工单
	var productionTicket models.ProductionTicket
	result := s.db.First(&productionTicket, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	productionTicket.ProductName = req["product_name"].(string)
	productionTicket.Quantity = req["quantity"].(int)
	productionTicket.Status = req["status"].(string)
	productionTicket.WorkCenterID = req["work_center_id"].(string)
	productionTicket.StartTime = time.Now()
	productionTicket.EndTime = time.Now().Add(24 * time.Hour)
	productionTicket.UpdatedAt = time.Now()
	productionTicket.UpdatedBy = req["updated_by"].(string)

	// 保存到数据库
	result = s.db.Save(&productionTicket)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	updatedTicket := map[string]interface{}{
		"id":                  productionTicket.ID,
		"ticket_no":           productionTicket.TicketNo,
		"production_order_id": productionTicket.ProductionOrderID,
		"product_name":        productionTicket.ProductName,
		"quantity":            productionTicket.Quantity,
		"status":              productionTicket.Status,
		"work_center_id":      productionTicket.WorkCenterID,
		"start_time":          productionTicket.StartTime,
		"end_time":            productionTicket.EndTime,
		"created_at":          productionTicket.CreatedAt,
		"created_by":          productionTicket.CreatedBy,
		"updated_at":          productionTicket.UpdatedAt,
		"updated_by":          productionTicket.UpdatedBy,
	}

	return updatedTicket, nil
}

func (s *productionService) DeleteProductionTicket(id string) error {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return errors.New("database connection is nil")
	}

	// 从数据库删除生产工单
	result := s.db.Delete(&models.ProductionTicket{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *productionService) StartProductionTicket(id string) error {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return errors.New("database connection is nil")
	}

	// 从数据库读取生产工单
	var productionTicket models.ProductionTicket
	result := s.db.First(&productionTicket, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为in_progress
	productionTicket.Status = "in_progress"
	productionTicket.UpdatedAt = time.Now()
	productionTicket.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&productionTicket)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *productionService) CompleteProductionTicket(id string) error {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return errors.New("database connection is nil")
	}

	// 从数据库读取生产工单
	var productionTicket models.ProductionTicket
	result := s.db.First(&productionTicket, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为completed
	productionTicket.Status = "completed"
	productionTicket.UpdatedAt = time.Now()
	productionTicket.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&productionTicket)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *productionService) CancelProductionTicket(id string) error {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回nil
		return nil
	}

	// 从数据库读取生产工单
	var productionTicket models.ProductionTicket
	result := s.db.First(&productionTicket, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为cancelled
	productionTicket.Status = "cancelled"
	productionTicket.UpdatedAt = time.Now()
	productionTicket.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&productionTicket)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 工艺路线管理方法
func (s *productionService) GetRoutingList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取工艺路线数据
	var routings []models.Routing
	result := s.db.Find(&routings)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	routingList := make([]map[string]interface{}, len(routings))
	for i, routing := range routings {
		routingList[i] = map[string]interface{}{
			"id":           routing.ID,
			"routing_no":   routing.RoutingNo,
			"product_name": routing.ProductName,
			"description":  routing.Description,
			"created_at":   routing.CreatedAt,
			"created_by":   routing.CreatedBy,
			"updated_at":   routing.UpdatedAt,
			"updated_by":   routing.UpdatedBy,
		}
	}

	return routingList, nil
}

func (s *productionService) GetRoutingDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取工艺路线详情
	var routing models.Routing
	result := s.db.First(&routing, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	routingDetail := map[string]interface{}{
		"id":           routing.ID,
		"routing_no":   routing.RoutingNo,
		"product_name": routing.ProductName,
		"description":  routing.Description,
		"created_at":   routing.CreatedAt,
		"created_by":   routing.CreatedBy,
		"updated_at":   routing.UpdatedAt,
		"updated_by":   routing.UpdatedBy,
	}

	return routingDetail, nil
}

func (s *productionService) CreateRouting(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从请求中获取数据
	routing := models.Routing{
		ID:          req["id"].(string),
		RoutingNo:   req["routing_no"].(string),
		ProductName: req["product_name"].(string),
		Description: req["description"].(string),
		CreatedAt:   time.Now(),
		CreatedBy:   req["created_by"].(string),
		UpdatedAt:   time.Now(),
		UpdatedBy:   req["created_by"].(string),
	}

	// 保存到数据库
	result := s.db.Create(&routing)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	createdRouting := map[string]interface{}{
		"id":           routing.ID,
		"routing_no":   routing.RoutingNo,
		"product_name": routing.ProductName,
		"description":  routing.Description,
		"created_at":   routing.CreatedAt,
		"created_by":   routing.CreatedBy,
		"updated_at":   routing.UpdatedAt,
		"updated_by":   routing.UpdatedBy,
	}

	return createdRouting, nil
}

func (s *productionService) UpdateRouting(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取工艺路线
	var routing models.Routing
	result := s.db.First(&routing, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	routing.ProductName = req["product_name"].(string)
	routing.Description = req["description"].(string)
	routing.UpdatedAt = time.Now()
	routing.UpdatedBy = req["updated_by"].(string)

	// 保存到数据库
	result = s.db.Save(&routing)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	updatedRouting := map[string]interface{}{
		"id":           routing.ID,
		"routing_no":   routing.RoutingNo,
		"product_name": routing.ProductName,
		"description":  routing.Description,
		"created_at":   routing.CreatedAt,
		"created_by":   routing.CreatedBy,
		"updated_at":   routing.UpdatedAt,
		"updated_by":   routing.UpdatedBy,
	}

	return updatedRouting, nil
}

func (s *productionService) DeleteRouting(id string) error {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return errors.New("database connection is nil")
	}

	// 从数据库删除工艺路线
	result := s.db.Delete(&models.Routing{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 工作中心管理方法
func (s *productionService) GetWorkCenterList(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取工作中心数据
	var workCenters []models.WorkCenter
	result := s.db.Find(&workCenters)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	workCenterList := make([]map[string]interface{}, len(workCenters))
	for i, workCenter := range workCenters {
		workCenterList[i] = map[string]interface{}{
			"id":             workCenter.ID,
			"work_center_no": workCenter.WorkCenterNo,
			"name":           workCenter.Name,
			"description":    workCenter.Description,
			"capacity":       workCenter.Capacity,
			"created_at":     workCenter.CreatedAt,
			"created_by":     workCenter.CreatedBy,
			"updated_at":     workCenter.UpdatedAt,
			"updated_by":     workCenter.UpdatedBy,
		}
	}

	return workCenterList, nil
}

func (s *productionService) GetWorkCenterDetail(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取工作中心详情
	var workCenter models.WorkCenter
	result := s.db.First(&workCenter, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	workCenterDetail := map[string]interface{}{
		"id":             workCenter.ID,
		"work_center_no": workCenter.WorkCenterNo,
		"name":           workCenter.Name,
		"description":    workCenter.Description,
		"capacity":       workCenter.Capacity,
		"created_at":     workCenter.CreatedAt,
		"created_by":     workCenter.CreatedBy,
		"updated_at":     workCenter.UpdatedAt,
		"updated_by":     workCenter.UpdatedBy,
	}

	return workCenterDetail, nil
}

func (s *productionService) CreateWorkCenter(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从请求中获取数据
	workCenter := models.WorkCenter{
		ID:           req["id"].(string),
		WorkCenterNo: req["work_center_no"].(string),
		Name:         req["name"].(string),
		Description:  req["description"].(string),
		Capacity:     req["capacity"].(int),
		CreatedAt:    time.Now(),
		CreatedBy:    req["created_by"].(string),
		UpdatedAt:    time.Now(),
		UpdatedBy:    req["created_by"].(string),
	}

	// 保存到数据库
	result := s.db.Create(&workCenter)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	createdWorkCenter := map[string]interface{}{
		"id":             workCenter.ID,
		"work_center_no": workCenter.WorkCenterNo,
		"name":           workCenter.Name,
		"description":    workCenter.Description,
		"capacity":       workCenter.Capacity,
		"created_at":     workCenter.CreatedAt,
		"created_by":     workCenter.CreatedBy,
		"updated_at":     workCenter.UpdatedAt,
		"updated_by":     workCenter.UpdatedBy,
	}

	return createdWorkCenter, nil
}

func (s *productionService) UpdateWorkCenter(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取工作中心
	var workCenter models.WorkCenter
	result := s.db.First(&workCenter, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	workCenter.Name = req["name"].(string)
	workCenter.Description = req["description"].(string)
	workCenter.Capacity = req["capacity"].(int)
	workCenter.UpdatedAt = time.Now()
	workCenter.UpdatedBy = req["updated_by"].(string)

	// 保存到数据库
	result = s.db.Save(&workCenter)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为map
	updatedWorkCenter := map[string]interface{}{
		"id":             workCenter.ID,
		"work_center_no": workCenter.WorkCenterNo,
		"name":           workCenter.Name,
		"description":    workCenter.Description,
		"capacity":       workCenter.Capacity,
		"created_at":     workCenter.CreatedAt,
		"created_by":     workCenter.CreatedBy,
		"updated_at":     workCenter.UpdatedAt,
		"updated_by":     workCenter.UpdatedBy,
	}

	return updatedWorkCenter, nil
}

func (s *productionService) DeleteWorkCenter(id string) error {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return errors.New("database connection is nil")
	}

	// 从数据库删除工作中心
	result := s.db.Delete(&models.WorkCenter{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *productionService) GetWorkCenterCapacity(id string) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取工作中心
	var workCenter models.WorkCenter
	result := s.db.First(&workCenter, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 计算工作中心容量
	capacityDetail := map[string]interface{}{
		"id":                 workCenter.ID,
		"work_center_no":     workCenter.WorkCenterNo,
		"name":               workCenter.Name,
		"total_capacity":     workCenter.Capacity,
		"used_capacity":      0, // 实际应用中需要计算已使用容量
		"available_capacity": workCenter.Capacity,
	}

	return capacityDetail, nil
}

func (s *productionService) GetWorkCenterSchedule(id string, req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取工作中心
	var workCenter models.WorkCenter
	result := s.db.First(&workCenter, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 生成工作中心调度
	scheduleDetail := map[string]interface{}{
		"id":             workCenter.ID,
		"work_center_no": workCenter.WorkCenterNo,
		"name":           workCenter.Name,
		"schedule":       []map[string]interface{}{},
	}

	return scheduleDetail, nil
}

// 物料需求计划管理方法
func (s *productionService) RunMRP(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要执行MRP计算
	// 1. 读取销售订单和预测数据
	// 2. 读取库存数据
	// 3. 读取生产订单数据
	// 4. 执行MRP计算
	// 5. 生成MRP结果和建议

	// 模拟MRP运行结果
	mrpResult := map[string]interface{}{
		"run_id":      "1",
		"run_date":    time.Now(),
		"status":      "completed",
		"total_items": 10,
		"suggestions": 5,
		"errors":      0,
	}

	return mrpResult, nil
}

func (s *productionService) GetMRPResults(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取MRP结果数据
	// 实际应用中，这里需要从数据库读取MRP结果

	// 模拟MRP结果数据
	mrpResults := []map[string]interface{}{
		{
			"id":              "1",
			"item_code":       "ITEM-001",
			"item_name":       "智能手机",
			"current_stock":   100,
			"demand":          500,
			"supply":          0,
			"net_requirement": 400,
			"suggested_qty":   400,
			"suggested_date":  time.Now().Add(7 * 24 * time.Hour),
		},
		{
			"id":              "2",
			"item_code":       "ITEM-002",
			"item_name":       "平板电脑",
			"current_stock":   50,
			"demand":          200,
			"supply":          0,
			"net_requirement": 150,
			"suggested_qty":   150,
			"suggested_date":  time.Now().Add(7 * 24 * time.Hour),
		},
	}

	return mrpResults, nil
}

func (s *productionService) GetMRPSuggestions(req map[string]interface{}) ([]map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取MRP建议数据
	// 实际应用中，这里需要从数据库读取MRP建议

	// 模拟MRP建议数据
	mrpSuggestions := []map[string]interface{}{
		{
			"id":              "1",
			"item_code":       "ITEM-001",
			"item_name":       "智能手机",
			"suggestion_type": "purchase",
			"suggested_qty":   400,
			"suggested_date":  time.Now().Add(7 * 24 * time.Hour),
			"supplier":        "供应商A",
		},
		{
			"id":              "2",
			"item_code":       "ITEM-002",
			"item_name":       "平板电脑",
			"suggestion_type": "production",
			"suggested_qty":   150,
			"suggested_date":  time.Now().Add(7 * 24 * time.Hour),
			"work_center":     "装配车间",
		},
	}

	return mrpSuggestions, nil
}

// 生产报表管理方法
func (s *productionService) GetProductionPlanReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取生产计划数据
	// 实际应用中，这里需要从数据库读取生产计划数据

	// 模拟生产计划报表数据
	report := map[string]interface{}{
		"report_type":    "production_plan",
		"report_date":    time.Now(),
		"period":         "month",
		"total_orders":   10,
		"total_quantity": 5000,
		"orders": []map[string]interface{}{
			{
				"order_no":     "PO-2026-0001",
				"product_name": "智能手机",
				"quantity":     1000,
				"start_date":   time.Now(),
				"end_date":     time.Now().Add(7 * 24 * time.Hour),
				"status":       "approved",
			},
			{
				"order_no":     "PO-2026-0002",
				"product_name": "平板电脑",
				"quantity":     500,
				"start_date":   time.Now().Add(2 * 24 * time.Hour),
				"end_date":     time.Now().Add(9 * 24 * time.Hour),
				"status":       "approved",
			},
		},
	}

	return report, nil
}

func (s *productionService) GetProductionExecutionReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取生产执行数据
	// 实际应用中，这里需要从数据库读取生产执行数据

	// 模拟生产执行报表数据
	report := map[string]interface{}{
		"report_type":       "production_execution",
		"report_date":       time.Now(),
		"period":            "month",
		"total_tickets":     20,
		"completed_tickets": 15,
		"completion_rate":   75.0,
		"tickets": []map[string]interface{}{
			{
				"ticket_no":    "PT-2026-0001",
				"product_name": "智能手机",
				"quantity":     100,
				"start_time":   time.Now().Add(-24 * time.Hour),
				"end_time":     time.Now().Add(-16 * time.Hour),
				"status":       "completed",
			},
			{
				"ticket_no":    "PT-2026-0002",
				"product_name": "平板电脑",
				"quantity":     50,
				"start_time":   time.Now().Add(-12 * time.Hour),
				"end_time":     time.Now().Add(-4 * time.Hour),
				"status":       "completed",
			},
		},
	}

	return report, nil
}

func (s *productionService) GetWorkCenterLoadReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取工作中心负载数据
	// 实际应用中，这里需要从数据库读取工作中心负载数据

	// 模拟工作中心负载报表数据
	report := map[string]interface{}{
		"report_type":        "work_center_load",
		"report_date":        time.Now(),
		"period":             "month",
		"total_work_centers": 2,
		"work_centers": []map[string]interface{}{
			{
				"work_center_no": "WC-2026-0001",
				"name":           "装配车间",
				"total_capacity": 1000,
				"used_capacity":  750,
				"load_rate":      75.0,
			},
			{
				"work_center_no": "WC-2026-0002",
				"name":           "测试车间",
				"total_capacity": 500,
				"used_capacity":  300,
				"load_rate":      60.0,
			},
		},
	}

	return report, nil
}

func (s *productionService) GetMaterialConsumptionReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取物料消耗数据
	// 实际应用中，这里需要从数据库读取物料消耗数据

	// 模拟物料消耗报表数据
	report := map[string]interface{}{
		"report_type":     "material_consumption",
		"report_date":     time.Now(),
		"period":          "month",
		"total_materials": 5,
		"materials": []map[string]interface{}{
			{
				"material_code": "MAT-001",
				"material_name": "屏幕",
				"consumption":   1000,
				"unit":          "个",
				"cost":          100000.0,
			},
			{
				"material_code": "MAT-002",
				"material_name": "电池",
				"consumption":   1000,
				"unit":          "个",
				"cost":          50000.0,
			},
		},
	}

	return report, nil
}

func (s *productionService) GetProductionCostReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取生产成本数据
	// 实际应用中，这里需要从数据库读取生产成本数据

	// 模拟生产成本报表数据
	report := map[string]interface{}{
		"report_type": "production_cost",
		"report_date": time.Now(),
		"period":      "month",
		"total_cost":  200000.0,
		"cost_breakdown": []map[string]interface{}{
			{
				"cost_type":  "material",
				"amount":     150000.0,
				"percentage": 75.0,
			},
			{
				"cost_type":  "labor",
				"amount":     30000.0,
				"percentage": 15.0,
			},
			{
				"cost_type":  "overhead",
				"amount":     20000.0,
				"percentage": 10.0,
			},
		},
	}

	return report, nil
}

func (s *productionService) ExportProductionReport(req map[string]interface{}) ([]byte, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接失败，返回错误
		return nil, errors.New("database connection is nil")
	}

	// 实际应用中，这里需要生成并导出报表
	// 1. 读取报表数据
	// 2. 生成报表文件（PDF, Excel等）
	// 3. 返回报表文件字节流

	// 模拟报表导出
	reportData := []byte("模拟报表数据")

	return reportData, nil
}

// getMockCreateProductionOrder 获取模拟创建生产订单数据
func (s *productionService) getMockCreateProductionOrder(req map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id":           req["id"],
		"order_no":     req["order_no"],
		"product_name": req["product_name"],
		"quantity":     req["quantity"],
		"status":       req["status"],
		"priority":     req["priority"],
		"start_date":   time.Now(),
		"end_date":     time.Now().Add(7 * 24 * time.Hour),
		"created_at":   time.Now(),
		"created_by":   req["created_by"],
		"updated_at":   time.Now(),
		"updated_by":   req["created_by"],
	}
}

// getMockUpdateProductionOrder 获取模拟更新生产订单数据
func (s *productionService) getMockUpdateProductionOrder(id string, req map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id":           id,
		"order_no":     "PO-2026-0001",
		"product_name": req["product_name"],
		"quantity":     req["quantity"],
		"status":       req["status"],
		"priority":     req["priority"],
		"start_date":   time.Now(),
		"end_date":     time.Now().Add(7 * 24 * time.Hour),
		"created_at":   time.Now().Add(-24 * time.Hour),
		"created_by":   "admin",
		"updated_at":   time.Now(),
		"updated_by":   req["updated_by"],
	}
}

// getMockProductionTicketList 获取模拟生产工单列表数据
func (s *productionService) getMockProductionTicketList() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"id":                  "1",
			"ticket_no":           "PT-2026-0001",
			"production_order_id": "1",
			"product_name":        "智能手机",
			"quantity":            100,
			"status":              "pending",
			"work_center_id":      "1",
			"start_time":          time.Now(),
			"end_time":            time.Now().Add(24 * time.Hour),
			"created_at":          time.Now(),
			"created_by":          "admin",
			"updated_at":          time.Now(),
			"updated_by":          "admin",
		},
		{
			"id":                  "2",
			"ticket_no":           "PT-2026-0002",
			"production_order_id": "1",
			"product_name":        "智能手机",
			"quantity":            100,
			"status":              "in_progress",
			"work_center_id":      "2",
			"start_time":          time.Now(),
			"end_time":            time.Now().Add(24 * time.Hour),
			"created_at":          time.Now(),
			"created_by":          "admin",
			"updated_at":          time.Now(),
			"updated_by":          "admin",
		},
	}
}

// getMockProductionTicketDetail 获取模拟生产工单详情数据
func (s *productionService) getMockProductionTicketDetail(id string) map[string]interface{} {
	return map[string]interface{}{
		"id":                  id,
		"ticket_no":           "PT-2026-0001",
		"production_order_id": "1",
		"product_name":        "智能手机",
		"quantity":            100,
		"status":              "pending",
		"work_center_id":      "1",
		"start_time":          time.Now(),
		"end_time":            time.Now().Add(24 * time.Hour),
		"created_at":          time.Now(),
		"created_by":          "admin",
		"updated_at":          time.Now(),
		"updated_by":          "admin",
	}
}

// getMockCreateProductionTicket 获取模拟创建生产工单数据
func (s *productionService) getMockCreateProductionTicket(req map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id":                  req["id"],
		"ticket_no":           req["ticket_no"],
		"production_order_id": req["production_order_id"],
		"product_name":        req["product_name"],
		"quantity":            req["quantity"],
		"status":              req["status"],
		"work_center_id":      req["work_center_id"],
		"start_time":          time.Now(),
		"end_time":            time.Now().Add(24 * time.Hour),
		"created_at":          time.Now(),
		"created_by":          req["created_by"],
		"updated_at":          time.Now(),
		"updated_by":          req["created_by"],
	}
}

// getMockUpdateProductionTicket 获取模拟更新生产工单数据
func (s *productionService) getMockUpdateProductionTicket(id string, req map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id":                  id,
		"ticket_no":           "PT-2026-0001",
		"production_order_id": "1",
		"product_name":        req["product_name"],
		"quantity":            req["quantity"],
		"status":              req["status"],
		"work_center_id":      req["work_center_id"],
		"start_time":          time.Now(),
		"end_time":            time.Now().Add(24 * time.Hour),
		"created_at":          time.Now().Add(-24 * time.Hour),
		"created_by":          "admin",
		"updated_at":          time.Now(),
		"updated_by":          req["updated_by"],
	}
}

// getMockRoutingList 获取模拟工艺路线列表数据
func (s *productionService) getMockRoutingList() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"id":           "1",
			"routing_no":   "RT-2026-0001",
			"product_name": "智能手机",
			"description":  "智能手机生产工艺路线",
			"created_at":   time.Now(),
			"created_by":   "admin",
			"updated_at":   time.Now(),
			"updated_by":   "admin",
		},
		{
			"id":           "2",
			"routing_no":   "RT-2026-0002",
			"product_name": "平板电脑",
			"description":  "平板电脑生产工艺路线",
			"created_at":   time.Now(),
			"created_by":   "admin",
			"updated_at":   time.Now(),
			"updated_by":   "admin",
		},
	}
}

// getMockRoutingDetail 获取模拟工艺路线详情数据
func (s *productionService) getMockRoutingDetail(id string) map[string]interface{} {
	return map[string]interface{}{
		"id":           id,
		"routing_no":   "RT-2026-0001",
		"product_name": "智能手机",
		"description":  "智能手机生产工艺路线",
		"created_at":   time.Now(),
		"created_by":   "admin",
		"updated_at":   time.Now(),
		"updated_by":   "admin",
	}
}

// getMockCreateRouting 获取模拟创建工艺路线数据
func (s *productionService) getMockCreateRouting(req map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id":           req["id"],
		"routing_no":   req["routing_no"],
		"product_name": req["product_name"],
		"description":  req["description"],
		"created_at":   time.Now(),
		"created_by":   req["created_by"],
		"updated_at":   time.Now(),
		"updated_by":   req["created_by"],
	}
}

// getMockUpdateRouting 获取模拟更新工艺路线数据
func (s *productionService) getMockUpdateRouting(id string, req map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id":           id,
		"routing_no":   "RT-2026-0001",
		"product_name": req["product_name"],
		"description":  req["description"],
		"created_at":   time.Now().Add(-24 * time.Hour),
		"created_by":   "admin",
		"updated_at":   time.Now(),
		"updated_by":   req["updated_by"],
	}
}

// getMockWorkCenterList 获取模拟工作中心列表数据
func (s *productionService) getMockWorkCenterList() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"id":             "1",
			"work_center_no": "WC-2026-0001",
			"name":           "装配车间",
			"description":    "智能手机装配车间",
			"capacity":       1000,
			"created_at":     time.Now(),
			"created_by":     "admin",
			"updated_at":     time.Now(),
			"updated_by":     "admin",
		},
		{
			"id":             "2",
			"work_center_no": "WC-2026-0002",
			"name":           "测试车间",
			"description":    "智能手机测试车间",
			"capacity":       500,
			"created_at":     time.Now(),
			"created_by":     "admin",
			"updated_at":     time.Now(),
			"updated_by":     "admin",
		},
	}
}

// getMockWorkCenterDetail 获取模拟工作中心详情数据
func (s *productionService) getMockWorkCenterDetail(id string) map[string]interface{} {
	return map[string]interface{}{
		"id":             id,
		"work_center_no": "WC-2026-0001",
		"name":           "装配车间",
		"description":    "智能手机装配车间",
		"capacity":       1000,
		"created_at":     time.Now(),
		"created_by":     "admin",
		"updated_at":     time.Now(),
		"updated_by":     "admin",
	}
}

// getMockCreateWorkCenter 获取模拟创建工作中心数据
func (s *productionService) getMockCreateWorkCenter(req map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id":             req["id"],
		"work_center_no": req["work_center_no"],
		"name":           req["name"],
		"description":    req["description"],
		"capacity":       req["capacity"],
		"created_at":     time.Now(),
		"created_by":     req["created_by"],
		"updated_at":     time.Now(),
		"updated_by":     req["created_by"],
	}
}

// getMockUpdateWorkCenter 获取模拟更新工作中心数据
func (s *productionService) getMockUpdateWorkCenter(id string, req map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id":             id,
		"work_center_no": "WC-2026-0001",
		"name":           req["name"],
		"description":    req["description"],
		"capacity":       req["capacity"],
		"created_at":     time.Now().Add(-24 * time.Hour),
		"created_by":     "admin",
		"updated_at":     time.Now(),
		"updated_by":     req["updated_by"],
	}
}

// getMockWorkCenterCapacity 获取模拟工作中心容量数据
func (s *productionService) getMockWorkCenterCapacity(id string) map[string]interface{} {
	return map[string]interface{}{
		"id":                 id,
		"work_center_no":     "WC-2026-0001",
		"name":               "装配车间",
		"total_capacity":     1000,
		"used_capacity":      500,
		"available_capacity": 500,
	}
}

// getMockWorkCenterSchedule 获取模拟工作中心调度数据
func (s *productionService) getMockWorkCenterSchedule(id string) map[string]interface{} {
	return map[string]interface{}{
		"id":             id,
		"work_center_no": "WC-2026-0001",
		"name":           "装配车间",
		"schedule": []map[string]interface{}{
			{
				"order_id":     "1",
				"order_no":     "PO-2026-0001",
				"product_name": "智能手机",
				"quantity":     100,
				"start_time":   time.Now(),
				"end_time":     time.Now().Add(8 * time.Hour),
			},
			{
				"order_id":     "2",
				"order_no":     "PO-2026-0002",
				"product_name": "平板电脑",
				"quantity":     50,
				"start_time":   time.Now().Add(8 * time.Hour),
				"end_time":     time.Now().Add(16 * time.Hour),
			},
		},
	}
}

// getMockRunMRP 获取模拟MRP运行结果
func (s *productionService) getMockRunMRP() map[string]interface{} {
	return map[string]interface{}{
		"run_id":      "1",
		"run_date":    time.Now(),
		"status":      "completed",
		"total_items": 10,
		"suggestions": 5,
		"errors":      0,
	}
}

// getMockMRPResults 获取模拟MRP结果数据
func (s *productionService) getMockMRPResults() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"id":              "1",
			"item_code":       "ITEM-001",
			"item_name":       "智能手机",
			"current_stock":   100,
			"demand":          500,
			"supply":          0,
			"net_requirement": 400,
			"suggested_qty":   400,
			"suggested_date":  time.Now().Add(7 * 24 * time.Hour),
		},
		{
			"id":              "2",
			"item_code":       "ITEM-002",
			"item_name":       "平板电脑",
			"current_stock":   50,
			"demand":          200,
			"supply":          0,
			"net_requirement": 150,
			"suggested_qty":   150,
			"suggested_date":  time.Now().Add(7 * 24 * time.Hour),
		},
	}
}

// getMockMRPSuggestions 获取模拟MRP建议数据
func (s *productionService) getMockMRPSuggestions() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"id":              "1",
			"item_code":       "ITEM-001",
			"item_name":       "智能手机",
			"suggestion_type": "purchase",
			"suggested_qty":   400,
			"suggested_date":  time.Now().Add(7 * 24 * time.Hour),
			"supplier":        "供应商A",
		},
		{
			"id":              "2",
			"item_code":       "ITEM-002",
			"item_name":       "平板电脑",
			"suggestion_type": "production",
			"suggested_qty":   150,
			"suggested_date":  time.Now().Add(7 * 24 * time.Hour),
			"work_center":     "装配车间",
		},
	}
}

// getMockProductionPlanReport 获取模拟生产计划报表数据
func (s *productionService) getMockProductionPlanReport() map[string]interface{} {
	return map[string]interface{}{
		"report_type":    "production_plan",
		"report_date":    time.Now(),
		"period":         "month",
		"total_orders":   10,
		"total_quantity": 5000,
		"orders": []map[string]interface{}{
			{
				"order_no":     "PO-2026-0001",
				"product_name": "智能手机",
				"quantity":     1000,
				"start_date":   time.Now(),
				"end_date":     time.Now().Add(7 * 24 * time.Hour),
				"status":       "approved",
			},
			{
				"order_no":     "PO-2026-0002",
				"product_name": "平板电脑",
				"quantity":     500,
				"start_date":   time.Now().Add(2 * 24 * time.Hour),
				"end_date":     time.Now().Add(9 * 24 * time.Hour),
				"status":       "approved",
			},
		},
	}
}

// getMockProductionExecutionReport 获取模拟生产执行报表数据
func (s *productionService) getMockProductionExecutionReport() map[string]interface{} {
	return map[string]interface{}{
		"report_type":       "production_execution",
		"report_date":       time.Now(),
		"period":            "month",
		"total_tickets":     20,
		"completed_tickets": 15,
		"completion_rate":   75.0,
		"tickets": []map[string]interface{}{
			{
				"ticket_no":    "PT-2026-0001",
				"product_name": "智能手机",
				"quantity":     100,
				"start_time":   time.Now().Add(-24 * time.Hour),
				"end_time":     time.Now().Add(-16 * time.Hour),
				"status":       "completed",
			},
			{
				"ticket_no":    "PT-2026-0002",
				"product_name": "平板电脑",
				"quantity":     50,
				"start_time":   time.Now().Add(-12 * time.Hour),
				"end_time":     time.Now().Add(-4 * time.Hour),
				"status":       "completed",
			},
		},
	}
}

// getMockWorkCenterLoadReport 获取模拟工作中心负载报表数据
func (s *productionService) getMockWorkCenterLoadReport() map[string]interface{} {
	return map[string]interface{}{
		"report_type":        "work_center_load",
		"report_date":        time.Now(),
		"period":             "month",
		"total_work_centers": 2,
		"work_centers": []map[string]interface{}{
			{
				"work_center_no": "WC-2026-0001",
				"name":           "装配车间",
				"total_capacity": 1000,
				"used_capacity":  750,
				"load_rate":      75.0,
			},
			{
				"work_center_no": "WC-2026-0002",
				"name":           "测试车间",
				"total_capacity": 500,
				"used_capacity":  300,
				"load_rate":      60.0,
			},
		},
	}
}

// getMockMaterialConsumptionReport 获取模拟物料消耗报表数据
func (s *productionService) getMockMaterialConsumptionReport() map[string]interface{} {
	return map[string]interface{}{
		"report_type":     "material_consumption",
		"report_date":     time.Now(),
		"period":          "month",
		"total_materials": 5,
		"materials": []map[string]interface{}{
			{
				"material_code": "MAT-001",
				"material_name": "屏幕",
				"consumption":   1000,
				"unit":          "个",
				"cost":          100000.0,
			},
			{
				"material_code": "MAT-002",
				"material_name": "电池",
				"consumption":   1000,
				"unit":          "个",
				"cost":          50000.0,
			},
		},
	}
}

// getMockProductionCostReport 获取模拟生产成本报表数据
func (s *productionService) getMockProductionCostReport() map[string]interface{} {
	return map[string]interface{}{
		"report_type": "production_cost",
		"report_date": time.Now(),
		"period":      "month",
		"total_cost":  200000.0,
		"cost_breakdown": []map[string]interface{}{
			{
				"cost_type":  "material",
				"amount":     150000.0,
				"percentage": 75.0,
			},
			{
				"cost_type":  "labor",
				"amount":     30000.0,
				"percentage": 15.0,
			},
			{
				"cost_type":  "overhead",
				"amount":     20000.0,
				"percentage": 10.0,
			},
		},
	}
}
