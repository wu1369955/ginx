package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/wu136995/ginx/internal/api/schemas"
	"github.com/wu136995/ginx/internal/database"
	"github.com/wu136995/ginx/internal/models"
	"gorm.io/gorm"
)

// InventoryService 库存服务接口
type InventoryService interface {
	// 仓库管理
	GetWarehouseList(req schemas.GetWarehouseListRequest) ([]schemas.WarehouseResponse, error)
	GetWarehouseDetail(id string) (*schemas.WarehouseResponse, error)
	CreateWarehouse(req schemas.CreateWarehouseRequest) (*schemas.WarehouseResponse, error)
	UpdateWarehouse(id string, req schemas.UpdateWarehouseRequest) (*schemas.WarehouseResponse, error)
	DeleteWarehouse(id string) error
	AddWarehouseLocation(id string, req schemas.AddWarehouseLocationRequest) (*schemas.LocationResponse, error)

	// 物料管理
	GetItemList(req schemas.GetItemListRequest) ([]schemas.ItemResponse, error)
	GetItemDetail(id string) (*schemas.ItemResponse, error)
	CreateItem(req schemas.CreateItemRequest) (*schemas.ItemResponse, error)
	UpdateItem(id string, req schemas.UpdateItemRequest) (*schemas.ItemResponse, error)
	DeleteItem(id string) error
	GetItemStock(id string) (*schemas.ItemStockResponse, error)

	// 库存交易管理
	GetTransactionList(req schemas.GetTransactionListRequest) ([]schemas.TransactionResponse, error)
	GetTransactionDetail(id string) (*schemas.TransactionResponse, error)
	CreateTransaction(req schemas.CreateTransactionRequest) (*schemas.TransactionResponse, error)
	CreateInventoryAdjustment(req schemas.CreateInventoryAdjustmentRequest) (*schemas.TransactionResponse, error)
	CreateWarehouseTransfer(req schemas.CreateWarehouseTransferRequest) (*schemas.TransactionResponse, error)

	// 库存盘点管理
	GetCountList(req schemas.GetCountListRequest) ([]schemas.CountResponse, error)
	GetCountDetail(id string) (*schemas.CountResponse, error)
	CreateCount(req schemas.CreateCountRequest) (*schemas.CountResponse, error)
	UpdateCount(id string, req schemas.UpdateCountRequest) (*schemas.CountResponse, error)
	CompleteCount(id string, req schemas.CompleteCountRequest) error
	CancelCount(id string, req schemas.CancelCountRequest) error

	// 库存报表管理
	GetInventoryBalanceReport(req schemas.GetInventoryBalanceReportRequest) (*schemas.InventoryBalanceReportResponse, error)
	GetInventoryMovementReport(req schemas.GetInventoryMovementReportRequest) (*schemas.InventoryMovementReportResponse, error)
	GetInventoryAlertReport(req schemas.GetInventoryAlertReportRequest) (*schemas.InventoryAlertReportResponse, error)
	GetInventoryABCReport(req schemas.GetInventoryABCReportRequest) (*schemas.InventoryABCReportResponse, error)
	ExportInventoryReport(req schemas.ExportInventoryReportRequest) ([]byte, error)
}

// inventoryService 库存服务实现
type inventoryService struct {
	db *gorm.DB
}

// NewInventoryService 创建库存服务实例
func NewInventoryService() InventoryService {
	return &inventoryService{
		db: database.GetDB(),
	}
}

// 仓库管理方法
func (s *inventoryService) GetWarehouseList(req schemas.GetWarehouseListRequest) ([]schemas.WarehouseResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取仓库数据
	var warehouses []models.InventoryWarehouse
	result := s.db.Find(&warehouses)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := make([]schemas.WarehouseResponse, len(warehouses))
	for i, warehouse := range warehouses {
		response[i] = schemas.WarehouseResponse{
			ID:           warehouse.ID,
			Code:         warehouse.Code,
			Name:         warehouse.Name,
			Type:         warehouse.Type,
			Address:      warehouse.Address,
			Region:       warehouse.Region,
			Contact:      warehouse.Contact,
			Phone:        warehouse.Phone,
			Description:  warehouse.Description,
			Capacity:     warehouse.Capacity,
			UsedCapacity: warehouse.UsedCapacity,
			Status:       warehouse.Status,
			CreatedBy:    warehouse.CreatedBy,
			CreatedAt:    warehouse.CreatedAt,
			UpdatedBy:    warehouse.UpdatedBy,
			UpdatedAt:    warehouse.UpdatedAt,
		}
	}

	return response, nil
}

func (s *inventoryService) GetWarehouseDetail(id string) (*schemas.WarehouseResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取仓库详情
	var warehouse models.InventoryWarehouse
	result := s.db.First(&warehouse, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := &schemas.WarehouseResponse{
		ID:           warehouse.ID,
		Code:         warehouse.Code,
		Name:         warehouse.Name,
		Type:         warehouse.Type,
		Address:      warehouse.Address,
		Region:       warehouse.Region,
		Contact:      warehouse.Contact,
		Phone:        warehouse.Phone,
		Description:  warehouse.Description,
		Capacity:     warehouse.Capacity,
		UsedCapacity: warehouse.UsedCapacity,
		Status:       warehouse.Status,
		CreatedBy:    warehouse.CreatedBy,
		CreatedAt:    warehouse.CreatedAt,
		UpdatedBy:    warehouse.UpdatedBy,
		UpdatedAt:    warehouse.UpdatedAt,
	}

	return response, nil
}

func (s *inventoryService) CreateWarehouse(req schemas.CreateWarehouseRequest) (*schemas.WarehouseResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 创建仓库模型
	warehouse := models.InventoryWarehouse{
		ID:          req.ID,
		Code:        req.Code,
		Name:        req.Name,
		Type:        req.Type,
		Address:     req.Address,
		Region:      req.Region,
		Contact:     req.Contact,
		Phone:       req.Phone,
		Description: req.Description,
		Capacity:    req.Capacity,
		Status:      req.Status,
		CreatedBy:   req.CreatedBy,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		UpdatedBy:   req.CreatedBy,
	}

	// 保存到数据库
	result := s.db.Create(&warehouse)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := &schemas.WarehouseResponse{
		ID:           warehouse.ID,
		Code:         warehouse.Code,
		Name:         warehouse.Name,
		Type:         warehouse.Type,
		Address:      warehouse.Address,
		Region:       warehouse.Region,
		Contact:      warehouse.Contact,
		Phone:        warehouse.Phone,
		Description:  warehouse.Description,
		Capacity:     warehouse.Capacity,
		UsedCapacity: warehouse.UsedCapacity,
		Status:       warehouse.Status,
		CreatedBy:    warehouse.CreatedBy,
		CreatedAt:    warehouse.CreatedAt,
		UpdatedBy:    warehouse.UpdatedBy,
		UpdatedAt:    warehouse.UpdatedAt,
	}

	return response, nil
}

func (s *inventoryService) UpdateWarehouse(id string, req schemas.UpdateWarehouseRequest) (*schemas.WarehouseResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取仓库
	var warehouse models.InventoryWarehouse
	result := s.db.First(&warehouse, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	if req.Name != "" {
		warehouse.Name = req.Name
	}
	if req.Type != "" {
		warehouse.Type = req.Type
	}
	if req.Address != "" {
		warehouse.Address = req.Address
	}
	if req.Region != "" {
		warehouse.Region = req.Region
	}
	if req.Contact != "" {
		warehouse.Contact = req.Contact
	}
	if req.Phone != "" {
		warehouse.Phone = req.Phone
	}
	if req.Description != "" {
		warehouse.Description = req.Description
	}
	if req.Capacity > 0 {
		warehouse.Capacity = req.Capacity
	}
	if req.Status != "" {
		warehouse.Status = req.Status
	}

	// 保存到数据库
	result = s.db.Save(&warehouse)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := &schemas.WarehouseResponse{
		ID:           warehouse.ID,
		Code:         warehouse.Code,
		Name:         warehouse.Name,
		Type:         warehouse.Type,
		Address:      warehouse.Address,
		Region:       warehouse.Region,
		Contact:      warehouse.Contact,
		Phone:        warehouse.Phone,
		Description:  warehouse.Description,
		Capacity:     warehouse.Capacity,
		UsedCapacity: warehouse.UsedCapacity,
		Status:       warehouse.Status,
		CreatedAt:    warehouse.CreatedAt,
		UpdatedAt:    warehouse.UpdatedAt,
	}

	return response, nil
}

func (s *inventoryService) DeleteWarehouse(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除仓库
	result := s.db.Delete(&models.InventoryWarehouse{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *inventoryService) AddWarehouseLocation(id string, req schemas.AddWarehouseLocationRequest) (*schemas.LocationResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 创建库位模型
	location := models.InventoryLocation{
		WarehouseID: id,
		Code:        req.Code,
		Name:        req.Name,
		Type:        req.Type,
		Capacity:    req.Capacity,
		Status:      req.Status,
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.CreatedBy,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// 保存到数据库
	result := s.db.Create(&location)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := &schemas.LocationResponse{
		ID:           location.ID,
		WarehouseID:  location.WarehouseID,
		Code:         location.Code,
		Name:         location.Name,
		Type:         location.Type,
		Capacity:     location.Capacity,
		UsedCapacity: location.UsedCapacity,
		Status:       location.Status,
		CreatedBy:    location.CreatedBy,
		CreatedAt:    location.CreatedAt,
		UpdatedBy:    location.UpdatedBy,
		UpdatedAt:    location.UpdatedAt,
	}

	return response, nil
}

// 物料管理方法
func (s *inventoryService) GetItemList(req schemas.GetItemListRequest) ([]schemas.ItemResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取物料数据
	var items []models.InventoryItem
	result := s.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := make([]schemas.ItemResponse, len(items))
	for i, item := range items {
		response[i] = schemas.ItemResponse{
			ID:          item.ID,
			ItemNo:      item.ItemNo,
			Name:        item.Name,
			Description: item.Description,
			CategoryID:  item.CategoryID,
			Unit:        item.Unit,
			Type:        item.Type,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}
	}

	return response, nil
}

func (s *inventoryService) GetItemDetail(id string) (*schemas.ItemResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取物料详情
	var item models.InventoryItem
	result := s.db.First(&item, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := &schemas.ItemResponse{
		ID:          item.ID,
		ItemNo:      item.ItemNo,
		Name:        item.Name,
		Description: item.Description,
		CategoryID:  item.CategoryID,
		Unit:        item.Unit,
		Type:        item.Type,
		Status:      item.Status,
		CreatedBy:   item.CreatedBy,
		CreatedAt:   item.CreatedAt,
		UpdatedBy:   item.UpdatedBy,
		UpdatedAt:   item.UpdatedAt,
	}

	return response, nil
}

func (s *inventoryService) CreateItem(req schemas.CreateItemRequest) (*schemas.ItemResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 创建物料模型
	item := models.InventoryItem{
		ItemNo:      req.ItemNo,
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		Unit:        req.Unit,
		Type:        req.Type,
		Status:      req.Status,
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.CreatedBy,
	}

	// 保存到数据库
	result := s.db.Create(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := &schemas.ItemResponse{
		ID:          item.ID,
		ItemNo:      item.ItemNo,
		Name:        item.Name,
		Description: item.Description,
		CategoryID:  item.CategoryID,
		Unit:        item.Unit,
		Type:        item.Type,
		Status:      item.Status,
		CreatedBy:   item.CreatedBy,
		CreatedAt:   item.CreatedAt,
		UpdatedBy:   item.UpdatedBy,
		UpdatedAt:   item.UpdatedAt,
	}

	return response, nil
}

func (s *inventoryService) UpdateItem(id string, req schemas.UpdateItemRequest) (*schemas.ItemResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取物料
	var item models.InventoryItem
	result := s.db.First(&item, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	if req.ItemNo != "" {
		item.ItemNo = req.ItemNo
	}
	if req.Name != "" {
		item.Name = req.Name
	}
	if req.Description != "" {
		item.Description = req.Description
	}
	if req.CategoryID != "" {
		item.CategoryID = req.CategoryID
	}
	if req.Unit != "" {
		item.Unit = req.Unit
	}
	if req.Type != "" {
		item.Type = req.Type
	}
	if req.Status != "" {
		item.Status = req.Status
	}
	item.UpdatedBy = req.UpdatedBy

	// 保存到数据库
	result = s.db.Save(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := &schemas.ItemResponse{
		ID:          item.ID,
		ItemNo:      item.ItemNo,
		Name:        item.Name,
		Description: item.Description,
		CategoryID:  item.CategoryID,
		Unit:        item.Unit,
		Type:        item.Type,
		Status:      item.Status,
		CreatedBy:   item.CreatedBy,
		CreatedAt:   item.CreatedAt,
		UpdatedBy:   item.UpdatedBy,
		UpdatedAt:   item.UpdatedAt,
	}

	return response, nil
}

func (s *inventoryService) DeleteItem(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除物料
	result := s.db.Delete(&models.InventoryItem{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *inventoryService) GetItemStock(id string) (*schemas.ItemStockResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取物料库存数据
	var onHandItems []models.InventoryOnHand
	result := s.db.Where("item_id = ?", id).Find(&onHandItems)
	if result.Error != nil {
		return nil, result.Error
	}

	// 计算总库存
	var totalQuantity float64
	var totalCost float64

	// 构建仓库库存响应
	warehouseStocks := make([]schemas.WarehouseStockResponse, 0)
	warehouseMap := make(map[string]schemas.WarehouseStockResponse)

	for _, onHand := range onHandItems {
		totalQuantity += onHand.Quantity
		totalCost += onHand.TotalCost

		// 获取仓库信息
		var warehouse models.InventoryWarehouse
		s.db.First(&warehouse, "id = ?", onHand.WarehouseID)

		// 检查仓库是否已经在映射中
		if _, exists := warehouseMap[onHand.WarehouseID]; !exists {
			warehouseMap[onHand.WarehouseID] = schemas.WarehouseStockResponse{
				WarehouseId:   onHand.WarehouseID,
				WarehouseName: warehouse.Name,
				Quantity:      0,
				Locations:     make([]schemas.LocationStockResponse, 0),
			}
		}

		// 获取仓库库存
		warehouseStock := warehouseMap[onHand.WarehouseID]
		warehouseStock.Quantity += onHand.Quantity

		// 添加库位库存
		warehouseStock.Locations = append(warehouseStock.Locations, schemas.LocationStockResponse{
			LocationId: onHand.LocationID,
			Quantity:   onHand.Quantity,
		})

		// 更新映射
		warehouseMap[onHand.WarehouseID] = warehouseStock
	}

	// 将映射转换为切片
	for _, warehouseStock := range warehouseMap {
		warehouseStocks = append(warehouseStocks, warehouseStock)
	}

	// 构建响应
	response := &schemas.ItemStockResponse{
		TotalStock: totalQuantity,
		Warehouses: warehouseStocks,
	}

	return response, nil
}

// 库存交易管理方法
func (s *inventoryService) GetTransactionList(req schemas.GetTransactionListRequest) ([]schemas.TransactionResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取交易数据
	var transactions []models.InventoryTransaction
	result := s.db.Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := make([]schemas.TransactionResponse, len(transactions))
	for i, transaction := range transactions {
		// 构建交易明细
		items := []schemas.TransactionItemResponse{
			{
				ItemId:    transaction.ItemID,
				Quantity:  transaction.Quantity,
				UnitCost:  transaction.UnitCost,
				TotalCost: transaction.TotalCost,
			},
		}

		response[i] = schemas.TransactionResponse{
			ID:              transaction.ID,
			TransactionNo:   transaction.TransactionNo,
			Type:            transaction.Type,
			WarehouseId:     transaction.WarehouseID,
			TransactionDate: transaction.TransactionDate.Format("2006-01-02"),
			Remarks:         transaction.Remarks,
			TotalQuantity:   transaction.Quantity,
			Items:           items,
			CreatedBy:       transaction.CreatedBy,
			CreatedAt:       transaction.CreatedAt,
		}
	}

	return response, nil
}

func (s *inventoryService) GetTransactionDetail(id string) (*schemas.TransactionResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取交易详情
	var transaction models.InventoryTransaction
	result := s.db.First(&transaction, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 构建交易明细
	items := []schemas.TransactionItemResponse{
		{
			ItemId:    transaction.ItemID,
			Quantity:  transaction.Quantity,
			UnitCost:  transaction.UnitCost,
			TotalCost: transaction.TotalCost,
		},
	}

	// 将模型转换为响应格式
	response := &schemas.TransactionResponse{
		ID:              transaction.ID,
		TransactionNo:   transaction.TransactionNo,
		Type:            transaction.Type,
		WarehouseId:     transaction.WarehouseID,
		TransactionDate: transaction.TransactionDate.Format("2006-01-02"),
		Remarks:         transaction.Remarks,
		TotalQuantity:   transaction.Quantity,
		Items:           items,
		CreatedBy:       transaction.CreatedBy,
		CreatedAt:       transaction.CreatedAt,
	}

	return response, nil
}

func (s *inventoryService) CreateTransaction(req schemas.CreateTransactionRequest) (*schemas.TransactionResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 解析交易日期
	transactionDate, err := time.Parse("2006-01-02", req.TransactionDate)
	if err != nil {
		return nil, err
	}

	// 生成交易编号
	transactionNo := fmt.Sprintf("TRX%s", time.Now().Format("20060102030405"))

	// 构建交易明细响应
	transactionItems := make([]schemas.TransactionItemResponse, len(req.Items))
	var totalQuantity float64

	// 为每个项目创建交易记录
	for i, item := range req.Items {
		// 计算总成本
		totalCost := item.Quantity * item.UnitCost
		totalQuantity += item.Quantity

		// 创建交易模型
		transaction := models.InventoryTransaction{
			TransactionNo:   transactionNo,
			ItemID:          item.ItemId,
			WarehouseID:     req.WarehouseId,
			LocationID:      item.LocationId,
			Type:            req.Type,
			Quantity:        item.Quantity,
			UnitCost:        item.UnitCost,
			TotalCost:       totalCost,
			ReferenceType:   "", // 可以根据实际情况设置
			ReferenceID:     "", // 可以根据实际情况设置
			TransactionDate: transactionDate,
			Remarks:         req.Remarks,
			CreatedBy:       "", // 可以根据实际情况设置
		}

		// 保存到数据库
		result := s.db.Create(&transaction)
		if result.Error != nil {
			return nil, result.Error
		}

		// 构建交易明细响应
		transactionItems[i] = schemas.TransactionItemResponse{
			ItemId:    item.ItemId,
			Quantity:  item.Quantity,
			UnitCost:  item.UnitCost,
			TotalCost: totalCost,
		}
	}

	// 将模型转换为响应格式
	response := &schemas.TransactionResponse{
		ID:              transactionNo, // 使用交易编号作为响应ID
		TransactionNo:   transactionNo,
		Type:            req.Type,
		WarehouseId:     req.WarehouseId,
		TransactionDate: req.TransactionDate,
		Remarks:         req.Remarks,
		TotalQuantity:   totalQuantity,
		Items:           transactionItems,
		CreatedBy:       "", // 可以根据实际情况设置
		CreatedAt:       time.Now(),
	}

	return response, nil
}

func (s *inventoryService) CreateInventoryAdjustment(req schemas.CreateInventoryAdjustmentRequest) (*schemas.TransactionResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 解析交易日期
	transactionDate, err := time.Parse("2006-01-02", req.TransactionDate)
	if err != nil {
		return nil, err
	}

	// 生成交易编号
	transactionNo := fmt.Sprintf("ADJ%s", time.Now().Format("20060102030405"))

	// 构建交易明细响应
	transactionItems := make([]schemas.TransactionItemResponse, len(req.Items))
	var totalQuantity float64

	// 为每个项目创建调整交易记录
	for i, item := range req.Items {
		// 计算总成本
		totalCost := item.Quantity * item.UnitCost
		totalQuantity += item.Quantity

		// 创建调整交易模型
		transaction := models.InventoryTransaction{
			TransactionNo:   transactionNo,
			ItemID:          item.ItemId,
			WarehouseID:     req.WarehouseId,
			LocationID:      item.LocationId,
			Type:            "adjustment",
			Quantity:        item.Quantity,
			UnitCost:        item.UnitCost,
			TotalCost:       totalCost,
			ReferenceType:   "inventory_adjustment",
			ReferenceID:     "", // 可以根据实际情况设置
			TransactionDate: transactionDate,
			Remarks:         req.Remarks,
			CreatedBy:       "", // 可以根据实际情况设置
		}

		// 保存到数据库
		result := s.db.Create(&transaction)
		if result.Error != nil {
			return nil, result.Error
		}

		// 构建交易明细响应
		transactionItems[i] = schemas.TransactionItemResponse{
			ItemId:    item.ItemId,
			Quantity:  item.Quantity,
			UnitCost:  item.UnitCost,
			TotalCost: totalCost,
		}
	}

	// 将模型转换为响应格式
	response := &schemas.TransactionResponse{
		ID:              transactionNo, // 使用交易编号作为响应ID
		TransactionNo:   transactionNo,
		Type:            "adjustment",
		WarehouseId:     req.WarehouseId,
		TransactionDate: req.TransactionDate,
		Remarks:         req.Remarks,
		TotalQuantity:   totalQuantity,
		Items:           transactionItems,
		CreatedBy:       "", // 可以根据实际情况设置
		CreatedAt:       time.Now(),
	}

	return response, nil
}

func (s *inventoryService) CreateWarehouseTransfer(req schemas.CreateWarehouseTransferRequest) (*schemas.TransactionResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 解析交易日期
	transactionDate, err := time.Parse("2006-01-02", req.TransactionDate)
	if err != nil {
		return nil, err
	}

	// 生成交易编号
	transactionNo := fmt.Sprintf("TRF%s", time.Now().Format("20060102030405"))

	// 构建交易明细响应
	transactionItems := make([]schemas.TransactionItemResponse, len(req.Items))
	var totalQuantity float64

	// 为每个项目创建转出和转入交易记录
	for i, item := range req.Items {
		// 计算总成本
		totalCost := item.Quantity * item.UnitCost
		totalQuantity += item.Quantity

		// 创建转出交易
		outTransaction := models.InventoryTransaction{
			TransactionNo:   transactionNo,
			ItemID:          item.ItemId,
			WarehouseID:     req.FromWarehouseId,
			LocationID:      item.FromLocationId,
			Type:            "transfer_out",
			Quantity:        -item.Quantity,
			UnitCost:        item.UnitCost,
			TotalCost:       -totalCost,
			ReferenceType:   "warehouse_transfer",
			ReferenceID:     "", // 可以根据实际情况设置
			TransactionDate: transactionDate,
			Remarks:         req.Remarks,
			CreatedBy:       "", // 可以根据实际情况设置
		}

		// 保存转出交易
		result := s.db.Create(&outTransaction)
		if result.Error != nil {
			return nil, result.Error
		}

		// 创建转入交易
		inTransaction := models.InventoryTransaction{
			TransactionNo:   transactionNo,
			ItemID:          item.ItemId,
			WarehouseID:     req.ToWarehouseId,
			LocationID:      item.ToLocationId,
			Type:            "transfer_in",
			Quantity:        item.Quantity,
			UnitCost:        item.UnitCost,
			TotalCost:       totalCost,
			ReferenceType:   "warehouse_transfer",
			ReferenceID:     "", // 可以根据实际情况设置
			TransactionDate: transactionDate,
			Remarks:         req.Remarks,
			CreatedBy:       "", // 可以根据实际情况设置
		}

		// 保存转入交易
		result = s.db.Create(&inTransaction)
		if result.Error != nil {
			return nil, result.Error
		}

		// 构建交易明细响应
		transactionItems[i] = schemas.TransactionItemResponse{
			ItemId:    item.ItemId,
			Quantity:  item.Quantity,
			UnitCost:  item.UnitCost,
			TotalCost: totalCost,
		}
	}

	// 将模型转换为响应格式
	response := &schemas.TransactionResponse{
		ID:              transactionNo, // 使用交易编号作为响应ID
		TransactionNo:   transactionNo,
		Type:            "transfer",
		FromWarehouseId: req.FromWarehouseId,
		ToWarehouseId:   req.ToWarehouseId,
		TransactionDate: req.TransactionDate,
		Remarks:         req.Remarks,
		TotalQuantity:   totalQuantity,
		Items:           transactionItems,
		CreatedBy:       "", // 可以根据实际情况设置
		CreatedAt:       time.Now(),
	}

	return response, nil
}

// 库存盘点管理方法
func (s *inventoryService) GetCountList(req schemas.GetCountListRequest) ([]schemas.CountResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取盘点数据
	var counts []models.InventoryCount
	result := s.db.Find(&counts)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := make([]schemas.CountResponse, len(counts))
	for i, count := range counts {
		response[i] = schemas.CountResponse{
			ID:          count.ID,
			CountNo:     count.CountNo,
			Type:        "stock_take", // 默认为库存盘点类型
			WarehouseId: count.WarehouseID,
			StartDate:   count.CountDate.Format("2006-01-02"),
			Status:      count.Status,
			Remarks:     count.Remarks,
			CreatedBy:   count.CreatedBy,
			CreatedAt:   count.CreatedAt,
			UpdatedBy:   count.UpdatedBy,
			UpdatedAt:   count.UpdatedAt,
		}
	}

	return response, nil
}

func (s *inventoryService) GetCountDetail(id string) (*schemas.CountResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取盘点详情
	var count models.InventoryCount
	result := s.db.First(&count, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := &schemas.CountResponse{
		ID:          count.ID,
		CountNo:     count.CountNo,
		Type:        "stock_take", // 默认为库存盘点类型
		WarehouseId: count.WarehouseID,
		StartDate:   count.CountDate.Format("2006-01-02"),
		Status:      count.Status,
		Remarks:     count.Remarks,
		CreatedBy:   count.CreatedBy,
		CreatedAt:   count.CreatedAt,
		UpdatedBy:   count.UpdatedBy,
		UpdatedAt:   count.UpdatedAt,
	}

	return response, nil
}

func (s *inventoryService) CreateCount(req schemas.CreateCountRequest) (*schemas.CountResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 解析盘点日期
	countDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, err
	}

	// 生成盘点编号
	countNo := fmt.Sprintf("CNT%s", time.Now().Format("20060102030405"))

	// 创建盘点模型
	count := models.InventoryCount{
		CountNo:     countNo,
		WarehouseID: req.WarehouseId,
		CountDate:   countDate,
		Status:      "pending",
		TotalItems:  len(req.Items),
		Remarks:     req.Remarks,
		CreatedBy:   "", // 可以根据实际情况设置
		UpdatedBy:   "", // 可以根据实际情况设置
	}

	// 保存到数据库
	result := s.db.Create(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := &schemas.CountResponse{
		ID:          count.ID,
		CountNo:     count.CountNo,
		Type:        req.Type,
		WarehouseId: count.WarehouseID,
		StartDate:   count.CountDate.Format("2006-01-02"),
		Status:      count.Status,
		Remarks:     count.Remarks,
		CreatedBy:   count.CreatedBy,
		CreatedAt:   count.CreatedAt,
		UpdatedBy:   count.UpdatedBy,
		UpdatedAt:   count.UpdatedAt,
	}

	return response, nil
}

func (s *inventoryService) UpdateCount(id string, req schemas.UpdateCountRequest) (*schemas.CountResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取盘点
	var count models.InventoryCount
	result := s.db.First(&count, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	if len(req.Items) > 0 {
		count.TotalItems = len(req.Items)
	}
	count.UpdatedBy = "" // 可以根据实际情况设置

	// 保存到数据库
	result = s.db.Save(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := &schemas.CountResponse{
		ID:          count.ID,
		CountNo:     count.CountNo,
		Type:        "stock_take", // 默认为库存盘点类型
		WarehouseId: count.WarehouseID,
		StartDate:   count.CountDate.Format("2006-01-02"),
		Status:      count.Status,
		Remarks:     count.Remarks,
		CreatedBy:   count.CreatedBy,
		CreatedAt:   count.CreatedAt,
		UpdatedBy:   count.UpdatedBy,
		UpdatedAt:   count.UpdatedAt,
	}

	return response, nil
}

func (s *inventoryService) CompleteCount(id string, req schemas.CompleteCountRequest) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取盘点
	var count models.InventoryCount
	result := s.db.First(&count, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为已完成
	count.Status = "completed"
	count.UpdatedBy = "" // 可以根据实际情况设置

	// 保存到数据库
	result = s.db.Save(&count)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *inventoryService) CancelCount(id string, req schemas.CancelCountRequest) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取盘点
	var count models.InventoryCount
	result := s.db.First(&count, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为已取消
	count.Status = "cancelled"
	count.UpdatedBy = "" // 可以根据实际情况设置

	// 保存到数据库
	result = s.db.Save(&count)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 库存报表管理方法
func (s *inventoryService) GetInventoryBalanceReport(req schemas.GetInventoryBalanceReportRequest) (*schemas.InventoryBalanceReportResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 这里简单返回空响应，实际实现中应该生成库存余额报表
	return &schemas.InventoryBalanceReportResponse{}, nil
}

func (s *inventoryService) GetInventoryMovementReport(req schemas.GetInventoryMovementReportRequest) (*schemas.InventoryMovementReportResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 这里简单返回空响应，实际实现中应该生成库存移动报表
	return &schemas.InventoryMovementReportResponse{}, nil
}

func (s *inventoryService) GetInventoryAlertReport(req schemas.GetInventoryAlertReportRequest) (*schemas.InventoryAlertReportResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 这里简单返回空响应，实际实现中应该生成库存警报报表
	return &schemas.InventoryAlertReportResponse{}, nil
}

func (s *inventoryService) GetInventoryABCReport(req schemas.GetInventoryABCReportRequest) (*schemas.InventoryABCReportResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 这里简单返回空响应，实际实现中应该生成库存ABC分析报表
	return &schemas.InventoryABCReportResponse{}, nil
}

func (s *inventoryService) ExportInventoryReport(req schemas.ExportInventoryReportRequest) ([]byte, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 这里简单返回空字节数组，实际实现中应该导出库存报表
	return []byte{}, nil
}
