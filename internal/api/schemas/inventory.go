package schemas

import "time"

// 仓库管理相关

// GetWarehouseListRequest 获取仓库列表请求
type GetWarehouseListRequest struct {
	Page        int    `form:"page" binding:"omitempty,min=1"`
	PageSize    int    `form:"pageSize" binding:"omitempty,min=1,max=100"`
	WarehouseNo string `form:"warehouseNo" binding:"omitempty"`
	Name        string `form:"name" binding:"omitempty"`
	Status      string `form:"status" binding:"omitempty,oneof=active inactive"`
	Type        string `form:"type" binding:"omitempty,oneof=raw_material finished_goods semi_finished tools"`
	Region      string `form:"region" binding:"omitempty"`
}

// LocationResponse 库位响应
type LocationResponse struct {
	ID           string    `json:"id"`
	WarehouseID  string    `json:"warehouse_id"`
	Code         string    `json:"code"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Capacity     float64   `json:"capacity"`
	UsedCapacity float64   `json:"usedCapacity"`
	Status       string    `json:"status"`
	Description  string    `json:"description,omitempty"`
	CreatedBy    string    `json:"createdBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedBy    string    `json:"updatedBy"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// WarehouseResponse 仓库响应
type WarehouseResponse struct {
	ID           string             `json:"id"`
	Code         string             `json:"code"`
	Name         string             `json:"name"`
	Type         string             `json:"type"`
	Address      string             `json:"address"`
	Region       string             `json:"region"`
	Contact      string             `json:"contact"`
	Phone        string             `json:"phone"`
	Description  string             `json:"description,omitempty"`
	Capacity     int                `json:"capacity,omitempty"`
	UsedCapacity int                `json:"usedCapacity,omitempty"`
	Status       string             `json:"status"`
	Locations    []LocationResponse `json:"locations,omitempty"`
	CreatedBy    string             `json:"createdBy"`
	CreatedAt    time.Time          `json:"createdAt"`
	UpdatedBy    string             `json:"updatedBy"`
	UpdatedAt    time.Time          `json:"updatedAt"`
}

// CreateWarehouseRequest 创建仓库请求
type CreateWarehouseRequest struct {
	ID          string `json:"id" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type" binding:"required,oneof=raw_material finished_goods semi_finished tools"`
	Address     string `json:"address" binding:"required"`
	Region      string `json:"region" binding:"required"`
	Contact     string `json:"contact" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
	Capacity    int    `json:"capacity" binding:"required,min=1"`
	Status      string `json:"status" binding:"omitempty,oneof=active inactive"`
	CreatedBy   string `json:"createdBy" binding:"required"`
}

// UpdateWarehouseRequest 更新仓库请求
type UpdateWarehouseRequest struct {
	Name        string `json:"name" binding:"omitempty"`
	Type        string `json:"type" binding:"omitempty,oneof=raw_material finished_goods semi_finished tools"`
	Address     string `json:"address" binding:"omitempty"`
	Region      string `json:"region" binding:"omitempty"`
	Contact     string `json:"contact" binding:"omitempty"`
	Phone       string `json:"phone" binding:"omitempty"`
	Description string `json:"description" binding:"omitempty"`
	Capacity    int    `json:"capacity" binding:"omitempty,min=1"`
	Status      string `json:"status" binding:"omitempty,oneof=active inactive"`
}

// AddWarehouseLocationRequest 添加库位请求
type AddWarehouseLocationRequest struct {
	Code        string  `json:"code" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Type        string  `json:"type" binding:"required"`
	Capacity    float64 `json:"capacity" binding:"required,min=0"`
	Status      string  `json:"status" binding:"required"`
	Description string  `json:"description" binding:"omitempty"`
	CreatedBy   string  `json:"createdBy" binding:"required"`
}

// 物料管理相关

// GetItemListRequest 获取物料列表请求
type GetItemListRequest struct {
	Page     int    `form:"page" binding:"omitempty,min=1"`
	PageSize int    `form:"pageSize" binding:"omitempty,min=1,max=100"`
	ItemNo   string `form:"itemNo" binding:"omitempty"`
	Name     string `form:"name" binding:"omitempty"`
	SKU      string `form:"sku" binding:"omitempty"`
	Category string `form:"category" binding:"omitempty"`
	Status   string `form:"status" binding:"omitempty,oneof=active inactive"`
	Unit     string `form:"unit" binding:"omitempty"`
}

// ItemResponse 物料响应
type ItemResponse struct {
	ID          string    `json:"id"`
	ItemNo      string    `json:"itemNo"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CategoryID  string    `json:"category_id"`
	Unit        string    `json:"unit"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedBy   string    `json:"updatedBy"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// CreateItemRequest 创建物料请求
type CreateItemRequest struct {
	ItemNo      string `json:"itemNo" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
	CategoryID  string `json:"category_id" binding:"required"`
	Unit        string `json:"unit" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Status      string `json:"status" binding:"required"`
	CreatedBy   string `json:"createdBy" binding:"required"`
}

// UpdateItemRequest 更新物料请求
type UpdateItemRequest struct {
	ItemNo      string `json:"itemNo" binding:"omitempty"`
	Name        string `json:"name" binding:"omitempty"`
	Description string `json:"description" binding:"omitempty"`
	CategoryID  string `json:"category_id" binding:"omitempty"`
	Unit        string `json:"unit" binding:"omitempty"`
	Type        string `json:"type" binding:"omitempty"`
	Status      string `json:"status" binding:"omitempty,oneof=active inactive"`
	UpdatedBy   string `json:"updatedBy" binding:"required"`
}

// LocationStockResponse 库位库存响应
type LocationStockResponse struct {
	LocationId   string  `json:"locationId"`
	LocationCode string  `json:"locationCode"`
	Quantity     float64 `json:"quantity"`
}

// WarehouseStockResponse 仓库库存响应
type WarehouseStockResponse struct {
	WarehouseId   string                  `json:"warehouseId"`
	WarehouseName string                  `json:"warehouseName"`
	Quantity      float64                 `json:"quantity"`
	Locations     []LocationStockResponse `json:"locations,omitempty"`
}

// ItemStockResponse 物料库存响应
type ItemStockResponse struct {
	TotalStock float64                  `json:"totalStock"`
	Warehouses []WarehouseStockResponse `json:"warehouses"`
}

// 库存交易相关

// GetTransactionListRequest 获取库存交易列表请求
type GetTransactionListRequest struct {
	Page                 int    `form:"page" binding:"omitempty,min=1"`
	PageSize             int    `form:"pageSize" binding:"omitempty,min=1,max=100"`
	TransactionNo        string `form:"transactionNo" binding:"omitempty"`
	Type                 string `form:"type" binding:"omitempty,oneof=purchase sales transfer adjustment production"`
	ItemId               string `form:"itemId" binding:"omitempty"`
	WarehouseId          string `form:"warehouseId" binding:"omitempty"`
	TransactionDateStart string `form:"transactionDateStart" binding:"omitempty,datetime=2006-01-02"`
	TransactionDateEnd   string `form:"transactionDateEnd" binding:"omitempty,datetime=2006-01-02"`
}

// TransactionItem 交易明细
type TransactionItem struct {
	ItemId         string  `json:"itemId" binding:"required"`
	Quantity       float64 `json:"quantity" binding:"required"`
	UnitCost       float64 `json:"unitCost" binding:"required,min=0"`
	LocationId     string  `json:"locationId" binding:"omitempty"`
	FromLocationId string  `json:"fromLocationId" binding:"omitempty"`
	ToLocationId   string  `json:"toLocationId" binding:"omitempty"`
}

// TransactionItemResponse 交易明细响应
type TransactionItemResponse struct {
	ID               string  `json:"id,omitempty"`
	ItemId           string  `json:"itemId"`
	ItemCode         string  `json:"itemCode,omitempty"`
	ItemName         string  `json:"itemName,omitempty"`
	Quantity         float64 `json:"quantity"`
	Unit             string  `json:"unit,omitempty"`
	UnitCost         float64 `json:"unitCost"`
	TotalCost        float64 `json:"totalCost,omitempty"`
	LocationId       string  `json:"locationId,omitempty"`
	LocationCode     string  `json:"locationCode,omitempty"`
	FromLocationId   string  `json:"fromLocationId,omitempty"`
	FromLocationCode string  `json:"fromLocationCode,omitempty"`
	ToLocationId     string  `json:"toLocationId,omitempty"`
	ToLocationCode   string  `json:"toLocationCode,omitempty"`
}

// TransactionResponse 交易响应
type TransactionResponse struct {
	ID                string                    `json:"id"`
	TransactionNo     string                    `json:"transactionNo"`
	Type              string                    `json:"type"`
	ReferenceNo       string                    `json:"referenceNo,omitempty"`
	WarehouseId       string                    `json:"warehouseId,omitempty"`
	WarehouseName     string                    `json:"warehouseName,omitempty"`
	FromWarehouseId   string                    `json:"fromWarehouseId,omitempty"`
	FromWarehouseName string                    `json:"fromWarehouseName,omitempty"`
	ToWarehouseId     string                    `json:"toWarehouseId,omitempty"`
	ToWarehouseName   string                    `json:"toWarehouseName,omitempty"`
	TransactionDate   string                    `json:"transactionDate"`
	Remarks           string                    `json:"remarks,omitempty"`
	TotalQuantity     float64                   `json:"totalQuantity,omitempty"`
	Items             []TransactionItemResponse `json:"items,omitempty"`
	CreatedBy         string                    `json:"createdBy"`
	CreatedAt         time.Time                 `json:"createdAt"`
}

// CreateTransactionRequest 创建库存交易请求
type CreateTransactionRequest struct {
	Type            string            `json:"type" binding:"required,oneof=purchase sales transfer adjustment production"`
	ReferenceNo     string            `json:"referenceNo" binding:"omitempty"`
	WarehouseId     string            `json:"warehouseId" binding:"omitempty"`
	FromWarehouseId string            `json:"fromWarehouseId" binding:"omitempty"`
	ToWarehouseId   string            `json:"toWarehouseId" binding:"omitempty"`
	TransactionDate string            `json:"transactionDate" binding:"required,datetime=2006-01-02"`
	Remarks         string            `json:"remarks" binding:"omitempty"`
	Items           []TransactionItem `json:"items" binding:"required,dive"`
}

// CreateInventoryAdjustmentRequest 创建库存调整请求
type CreateInventoryAdjustmentRequest struct {
	WarehouseId     string            `json:"warehouseId" binding:"required"`
	TransactionDate string            `json:"transactionDate" binding:"required,datetime=2006-01-02"`
	Reason          string            `json:"reason" binding:"required"`
	Remarks         string            `json:"remarks" binding:"omitempty"`
	Items           []TransactionItem `json:"items" binding:"required,dive"`
}

// CreateWarehouseTransferRequest 创建仓库间转移请求
type CreateWarehouseTransferRequest struct {
	FromWarehouseId string            `json:"fromWarehouseId" binding:"required"`
	ToWarehouseId   string            `json:"toWarehouseId" binding:"required"`
	TransactionDate string            `json:"transactionDate" binding:"required,datetime=2006-01-02"`
	Remarks         string            `json:"remarks" binding:"omitempty"`
	Items           []TransactionItem `json:"items" binding:"required,dive"`
}

// 库存盘点相关

// GetCountListRequest 获取盘点列表请求
type GetCountListRequest struct {
	Page           int    `form:"page" binding:"omitempty,min=1"`
	PageSize       int    `form:"pageSize" binding:"omitempty,min=1,max=100"`
	CountNo        string `form:"countNo" binding:"omitempty"`
	Type           string `form:"type" binding:"omitempty,oneof=cycle stock_take"`
	WarehouseId    string `form:"warehouseId" binding:"omitempty"`
	Status         string `form:"status" binding:"omitempty,oneof=draft in_progress completed cancelled"`
	StartDateStart string `form:"startDateStart" binding:"omitempty,datetime=2006-01-02"`
	StartDateEnd   string `form:"startDateEnd" binding:"omitempty,datetime=2006-01-02"`
}

// CountItem 盘点明细
type CountItem struct {
	ItemId         string  `json:"itemId" binding:"required"`
	LocationId     string  `json:"locationId" binding:"required"`
	ActualQuantity float64 `json:"actualQuantity" binding:"omitempty"`
}

// CountItemResponse 盘点明细响应
type CountItemResponse struct {
	ID               string  `json:"id,omitempty"`
	ItemId           string  `json:"itemId"`
	ItemCode         string  `json:"itemCode,omitempty"`
	ItemName         string  `json:"itemName,omitempty"`
	SystemQuantity   float64 `json:"systemQuantity"`
	ActualQuantity   float64 `json:"actualQuantity"`
	Difference       float64 `json:"difference,omitempty"`
	UnitCost         float64 `json:"unitCost,omitempty"`
	DifferenceAmount float64 `json:"differenceAmount,omitempty"`
	LocationId       string  `json:"locationId"`
	LocationCode     string  `json:"locationCode,omitempty"`
}

// CountResponse 盘点响应
type CountResponse struct {
	ID            string              `json:"id"`
	CountNo       string              `json:"countNo"`
	Type          string              `json:"type"`
	WarehouseId   string              `json:"warehouseId"`
	WarehouseName string              `json:"warehouseName,omitempty"`
	StartDate     string              `json:"startDate"`
	EndDate       string              `json:"endDate,omitempty"`
	Remarks       string              `json:"remarks,omitempty"`
	Status        string              `json:"status"`
	Items         []CountItemResponse `json:"items,omitempty"`
	CreatedBy     string              `json:"createdBy"`
	CreatedAt     time.Time           `json:"createdAt"`
	UpdatedBy     string              `json:"updatedBy"`
	UpdatedAt     time.Time           `json:"updatedAt"`
}

// CreateCountRequest 创建盘点单请求
type CreateCountRequest struct {
	Type        string      `json:"type" binding:"required,oneof=cycle stock_take"`
	WarehouseId string      `json:"warehouseId" binding:"required"`
	StartDate   string      `json:"startDate" binding:"required,datetime=2006-01-02"`
	Remarks     string      `json:"remarks" binding:"omitempty"`
	Items       []CountItem `json:"items" binding:"required,dive"`
}

// UpdateCountRequest 更新盘点结果请求
type UpdateCountRequest struct {
	EndDate string      `json:"endDate" binding:"omitempty,datetime=2006-01-02"`
	Items   []CountItem `json:"items" binding:"required,dive"`
}

// CompleteCountRequest 完成盘点请求
type CompleteCountRequest struct {
	AdjustInventory bool `json:"adjustInventory" binding:"required"`
}

// CancelCountRequest 取消盘点请求
type CancelCountRequest struct {
	Reason string `json:"reason" binding:"required"`
}

// 库存报表相关

// GetInventoryBalanceReportRequest 获取库存余额报表请求
type GetInventoryBalanceReportRequest struct {
	WarehouseId string `form:"warehouseId" binding:"omitempty"`
	ItemId      string `form:"itemId" binding:"omitempty"`
	Category    string `form:"category" binding:"omitempty"`
	AsOfDate    string `form:"asOfDate" binding:"omitempty,datetime=2006-01-02"`
}

// BalanceReportItem 余额报表项
type BalanceReportItem struct {
	ItemId        string  `json:"itemId"`
	ItemCode      string  `json:"itemCode"`
	ItemName      string  `json:"itemName"`
	WarehouseId   string  `json:"warehouseId"`
	WarehouseName string  `json:"warehouseName"`
	Quantity      float64 `json:"quantity"`
	UnitCost      float64 `json:"unitCost"`
	TotalValue    float64 `json:"totalValue"`
}

// InventoryBalanceReportResponse 库存余额报表响应
type InventoryBalanceReportResponse struct {
	AsOfDate   string              `json:"asOfDate"`
	TotalItems int                 `json:"totalItems"`
	TotalValue float64             `json:"totalValue"`
	Items      []BalanceReportItem `json:"items"`
}

// GetInventoryMovementReportRequest 获取库存变动报表请求
type GetInventoryMovementReportRequest struct {
	StartDate   string `form:"startDate" binding:"required,datetime=2006-01-02"`
	EndDate     string `form:"endDate" binding:"required,datetime=2006-01-02"`
	WarehouseId string `form:"warehouseId" binding:"omitempty"`
	ItemId      string `form:"itemId" binding:"omitempty"`
	Type        string `form:"type" binding:"omitempty"`
}

// TransactionDetail 交易明细
type TransactionDetail struct {
	TransactionNo string  `json:"transactionNo"`
	Type          string  `json:"type"`
	Date          string  `json:"date"`
	Quantity      float64 `json:"quantity"`
}

// MovementReportItem 变动报表项
type MovementReportItem struct {
	ItemId           string              `json:"itemId"`
	ItemCode         string              `json:"itemCode"`
	ItemName         string              `json:"itemName"`
	BeginningBalance float64             `json:"beginningBalance"`
	InQuantity       float64             `json:"inQuantity"`
	OutQuantity      float64             `json:"outQuantity"`
	EndingBalance    float64             `json:"endingBalance"`
	Transactions     []TransactionDetail `json:"transactions,omitempty"`
}

// InventoryMovementReportResponse 库存变动报表响应
type InventoryMovementReportResponse struct {
	Period string               `json:"period"`
	Items  []MovementReportItem `json:"items"`
}

// GetInventoryAlertReportRequest 获取库存预警报表请求
type GetInventoryAlertReportRequest struct {
	WarehouseId string `form:"warehouseId" binding:"omitempty"`
	Type        string `form:"type" binding:"omitempty,oneof=low_stock excess_stock reorder_point"`
}

// AlertReportItem 预警报表项
type AlertReportItem struct {
	ItemId        string  `json:"itemId"`
	ItemCode      string  `json:"itemCode"`
	ItemName      string  `json:"itemName"`
	WarehouseId   string  `json:"warehouseId"`
	WarehouseName string  `json:"warehouseName"`
	CurrentStock  float64 `json:"currentStock"`
	MinStock      int     `json:"minStock,omitempty"`
	MaxStock      int     `json:"maxStock,omitempty"`
	ReorderPoint  int     `json:"reorderPoint,omitempty"`
}

// InventoryAlertReportResponse 库存预警报表响应
type InventoryAlertReportResponse struct {
	LowStockItems     []AlertReportItem `json:"lowStockItems"`
	ExcessStockItems  []AlertReportItem `json:"excessStockItems"`
	ReorderPointItems []AlertReportItem `json:"reorderPointItems"`
}

// GetInventoryABCReportRequest 获取ABC分析报表请求
type GetInventoryABCReportRequest struct {
	StartDate   string `form:"startDate" binding:"required,datetime=2006-01-02"`
	EndDate     string `form:"endDate" binding:"required,datetime=2006-01-02"`
	WarehouseId string `form:"warehouseId" binding:"omitempty"`
}

// ABCReportItem ABC分析报表项
type ABCReportItem struct {
	ItemId     string  `json:"itemId"`
	ItemCode   string  `json:"itemCode"`
	ItemName   string  `json:"itemName"`
	Category   string  `json:"category"`
	Quantity   float64 `json:"quantity"`
	Value      float64 `json:"value"`
	Percentage float64 `json:"percentage"`
}

// InventoryABCReportResponse ABC分析报表响应
type InventoryABCReportResponse struct {
	Period string          `json:"period"`
	Items  []ABCReportItem `json:"items"`
}

// ExportInventoryReportRequest 导出库存报表请求
type ExportInventoryReportRequest struct {
	Type        string `form:"type" binding:"required,oneof=balance movement alert abc"`
	StartDate   string `form:"startDate" binding:"omitempty,datetime=2006-01-02"`
	EndDate     string `form:"endDate" binding:"omitempty,datetime=2006-01-02"`
	WarehouseId string `form:"warehouseId" binding:"omitempty"`
	ItemId      string `form:"itemId" binding:"omitempty"`
}
