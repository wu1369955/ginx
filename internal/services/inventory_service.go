package services

import (
	"errors"
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
	if req.Code != "" {
		warehouse.Code = req.Code
	}
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
	warehouse.UpdatedBy = req.UpdatedBy

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
		CreatedAt:    location.CreatedAt,
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
		CreatedAt:   item.CreatedAt,
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
		CreatedAt:   item.CreatedAt,
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
		CreatedAt:   item.CreatedAt,
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

	// 构建库存明细
	stockDetails := make([]schemas.StockDetail, len(onHandItems))
	for i, onHand := range onHandItems {
		totalQuantity += onHand.Quantity
		totalCost += onHand.TotalCost

		// 获取仓库信息
		var warehouse models.InventoryWarehouse
		s.db.First(&warehouse, "id = ?", onHand.WarehouseID)

		// 构建库存明细
		stockDetails[i] = schemas.StockDetail{
			WarehouseID:   onHand.WarehouseID,
			WarehouseName: warehouse.Name,
			LocationID:    onHand.LocationID,
			Quantity:      onHand.Quantity,
			UnitCost:      onHand.UnitCost,
			TotalCost:     onHand.TotalCost,
		}
	}

	// 构建响应
	response := &schemas.ItemStockResponse{
		ItemID:        id,
		TotalQuantity: totalQuantity,
		TotalCost:     totalCost,
		StockDetails:  stockDetails,
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
		response[i] = schemas.TransactionResponse{
			ID:              transaction.ID,
			TransactionNo:   transaction.TransactionNo,
			ItemID:          transaction.ItemID,
			WarehouseID:     transaction.WarehouseID,
			LocationID:      transaction.LocationID,
			Type:            transaction.Type,
			Quantity:        transaction.Quantity,
			UnitCost:        transaction.UnitCost,
			TotalCost:       transaction.TotalCost,
			ReferenceType:   transaction.ReferenceType,
			ReferenceID:     transaction.ReferenceID,
			TransactionDate: transaction.TransactionDate,
			Remarks:         transaction.Remarks,
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

	// 将模型转换为响应格式
	response := &schemas.TransactionResponse{
		ID:              transaction.ID,
		TransactionNo:   transaction.TransactionNo,
		ItemID:          transaction.ItemID,
		WarehouseID:     transaction.WarehouseID,
		LocationID:      transaction.LocationID,
		Type:            transaction.Type,
		Quantity:        transaction.Quantity,
		UnitCost:        transaction.UnitCost,
		TotalCost:       transaction.TotalCost,
		ReferenceType:   transaction.ReferenceType,
		ReferenceID:     transaction.ReferenceID,
		TransactionDate: transaction.TransactionDate,
		Remarks:         transaction.Remarks,
		CreatedAt:       transaction.CreatedAt,
	}

	return response, nil
}

func (s *inventoryService) CreateTransaction(req schemas.CreateTransactionRequest) (*schemas.TransactionResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 创建交易模型
	transaction := models.InventoryTransaction{
		TransactionNo:   req.TransactionNo,
		ItemID:          req.ItemID,
		WarehouseID:     req.WarehouseID,
		LocationID:      req.LocationID,
		Type:            req.Type,
		Quantity:        req.Quantity,
		UnitCost:        req.UnitCost,
		TotalCost:       req.TotalCost,
		ReferenceType:   req.ReferenceType,
		ReferenceID:     req.ReferenceID,
		TransactionDate: req.TransactionDate,
		Remarks:         req.Remarks,
		CreatedBy:       req.CreatedBy,
	}

	// 保存到数据库
	result := s.db.Create(&transaction)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := &schemas.TransactionResponse{
		ID:              transaction.ID,
		TransactionNo:   transaction.TransactionNo,
		ItemID:          transaction.ItemID,
		WarehouseID:     transaction.WarehouseID,
		LocationID:      transaction.LocationID,
		Type:            transaction.Type,
		Quantity:        transaction.Quantity,
		UnitCost:        transaction.UnitCost,
		TotalCost:       transaction.TotalCost,
		ReferenceType:   transaction.ReferenceType,
		ReferenceID:     transaction.ReferenceID,
		TransactionDate: transaction.TransactionDate,
		Remarks:         transaction.Remarks,
		CreatedAt:       transaction.CreatedAt,
	}

	return response, nil
}

func (s *inventoryService) CreateInventoryAdjustment(req schemas.CreateInventoryAdjustmentRequest) (*schemas.TransactionResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 创建调整交易模型
	transaction := models.InventoryTransaction{
		TransactionNo:   req.TransactionNo,
		ItemID:          req.ItemID,
		WarehouseID:     req.WarehouseID,
		LocationID:      req.LocationID,
		Type:            "adjustment",
		Quantity:        req.Quantity,
		UnitCost:        req.UnitCost,
		TotalCost:       req.Quantity * req.UnitCost,
		ReferenceType:   "inventory_adjustment",
		ReferenceID:     req.ReferenceID,
		TransactionDate: req.TransactionDate,
		Remarks:         req.Remarks,
		CreatedBy:       req.CreatedBy,
	}

	// 保存到数据库
	result := s.db.Create(&transaction)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := &schemas.TransactionResponse{
		ID:              transaction.ID,
		TransactionNo:   transaction.TransactionNo,
		ItemID:          transaction.ItemID,
		WarehouseID:     transaction.WarehouseID,
		LocationID:      transaction.LocationID,
		Type:            transaction.Type,
		Quantity:        transaction.Quantity,
		UnitCost:        transaction.UnitCost,
		TotalCost:       transaction.TotalCost,
		ReferenceType:   transaction.ReferenceType,
		ReferenceID:     transaction.ReferenceID,
		TransactionDate: transaction.TransactionDate,
		Remarks:         transaction.Remarks,
		CreatedAt:       transaction.CreatedAt,
	}

	return response, nil
}

func (s *inventoryService) CreateWarehouseTransfer(req schemas.CreateWarehouseTransferRequest) (*schemas.TransactionResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 创建转出交易
	outTransaction := models.InventoryTransaction{
		TransactionNo:   req.TransactionNo,
		ItemID:          req.ItemID,
		WarehouseID:     req.FromWarehouseID,
		LocationID:      req.FromLocationID,
		Type:            "transfer_out",
		Quantity:        -req.Quantity,
		UnitCost:        req.UnitCost,
		TotalCost:       -req.Quantity * req.UnitCost,
		ReferenceType:   "warehouse_transfer",
		ReferenceID:     req.ReferenceID,
		TransactionDate: req.TransactionDate,
		Remarks:         req.Remarks,
		CreatedBy:       req.CreatedBy,
	}

	// 保存转出交易
	result := s.db.Create(&outTransaction)
	if result.Error != nil {
		return nil, result.Error
	}

	// 创建转入交易
	inTransaction := models.InventoryTransaction{
		TransactionNo:   req.TransactionNo,
		ItemID:          req.ItemID,
		WarehouseID:     req.ToWarehouseID,
		LocationID:      req.ToLocationID,
		Type:            "transfer_in",
		Quantity:        req.Quantity,
		UnitCost:        req.UnitCost,
		TotalCost:       req.Quantity * req.UnitCost,
		ReferenceType:   "warehouse_transfer",
		ReferenceID:     req.ReferenceID,
		TransactionDate: req.TransactionDate,
		Remarks:         req.Remarks,
		CreatedBy:       req.CreatedBy,
	}

	// 保存转入交易
	result = s.db.Create(&inTransaction)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := &schemas.TransactionResponse{
		ID:              inTransaction.ID,
		TransactionNo:   inTransaction.TransactionNo,
		ItemID:          inTransaction.ItemID,
		WarehouseID:     inTransaction.WarehouseID,
		LocationID:      inTransaction.LocationID,
		Type:            "transfer",
		Quantity:        inTransaction.Quantity,
		UnitCost:        inTransaction.UnitCost,
		TotalCost:       inTransaction.TotalCost,
		ReferenceType:   inTransaction.ReferenceType,
		ReferenceID:     inTransaction.ReferenceID,
		TransactionDate: inTransaction.TransactionDate,
		Remarks:         inTransaction.Remarks,
		CreatedAt:       inTransaction.CreatedAt,
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
			ID:            count.ID,
			CountNo:       count.CountNo,
			WarehouseID:   count.WarehouseID,
			CountDate:     count.CountDate,
			Status:        count.Status,
			TotalItems:    count.TotalItems,
			TotalVariance: count.TotalVariance,
			Remarks:       count.Remarks,
			CreatedAt:     count.CreatedAt,
			UpdatedAt:     count.UpdatedAt,
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
		ID:            count.ID,
		CountNo:       count.CountNo,
		WarehouseID:   count.WarehouseID,
		CountDate:     count.CountDate,
		Status:        count.Status,
		TotalItems:    count.TotalItems,
		TotalVariance: count.TotalVariance,
		Remarks:       count.Remarks,
		CreatedAt:     count.CreatedAt,
		UpdatedAt:     count.UpdatedAt,
	}

	return response, nil
}

func (s *inventoryService) CreateCount(req schemas.CreateCountRequest) (*schemas.CountResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 创建盘点模型
	count := models.InventoryCount{
		CountNo:     req.CountNo,
		WarehouseID: req.WarehouseID,
		CountDate:   req.CountDate,
		Status:      "pending",
		TotalItems:  req.TotalItems,
		Remarks:     req.Remarks,
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.CreatedBy,
	}

	// 保存到数据库
	result := s.db.Create(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := &schemas.CountResponse{
		ID:            count.ID,
		CountNo:       count.CountNo,
		WarehouseID:   count.WarehouseID,
		CountDate:     count.CountDate,
		Status:        count.Status,
		TotalItems:    count.TotalItems,
		TotalVariance: count.TotalVariance,
		Remarks:       count.Remarks,
		CreatedAt:     count.CreatedAt,
		UpdatedAt:     count.UpdatedAt,
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
	if req.CountNo != "" {
		count.CountNo = req.CountNo
	}
	if req.WarehouseID != "" {
		count.WarehouseID = req.WarehouseID
	}
	if !req.CountDate.IsZero() {
		count.CountDate = req.CountDate
	}
	if req.Status != "" {
		count.Status = req.Status
	}
	if req.TotalItems > 0 {
		count.TotalItems = req.TotalItems
	}
	if req.Remarks != "" {
		count.Remarks = req.Remarks
	}
	count.UpdatedBy = req.UpdatedBy

	// 保存到数据库
	result = s.db.Save(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	response := &schemas.CountResponse{
		ID:            count.ID,
		CountNo:       count.CountNo,
		WarehouseID:   count.WarehouseID,
		CountDate:     count.CountDate,
		Status:        count.Status,
		TotalItems:    count.TotalItems,
		TotalVariance: count.TotalVariance,
		Remarks:       count.Remarks,
		CreatedAt:     count.CreatedAt,
		UpdatedAt:     count.UpdatedAt,
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
	count.TotalVariance = req.TotalVariance
	count.UpdatedBy = req.UpdatedBy

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
	count.UpdatedBy = req.UpdatedBy

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
