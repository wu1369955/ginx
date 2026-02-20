package models

import (
	"time"

	"gorm.io/gorm"
)

// InventoryWarehouse 仓库表模型
type InventoryWarehouse struct {
	ID            string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Code          string         `json:"code" gorm:"unique;not null;type:varchar(20)"`
	Name          string         `json:"name" gorm:"not null;type:varchar(100)"`
	Type          string         `json:"type" gorm:"not null;type:varchar(20)"`
	Address       string         `json:"address" gorm:"type:varchar(255)"`
	Region        string         `json:"region" gorm:"type:varchar(50)"`
	Contact       string         `json:"contact" gorm:"type:varchar(50)"`
	Phone         string         `json:"phone" gorm:"type:varchar(20)"`
	Description   string         `json:"description" gorm:"type:text"`
	Capacity      int            `json:"capacity" gorm:"type:int;default:0"`
	UsedCapacity  int            `json:"used_capacity" gorm:"type:int;default:0"`
	Status        string         `json:"status" gorm:"type:varchar(20);default:'active'"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy     string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Locations    []InventoryLocation    `json:"locations,omitempty" gorm:"foreignKey:WarehouseID"`
	OnHandItems  []InventoryOnHand     `json:"on_hand_items,omitempty" gorm:"foreignKey:WarehouseID"`
	Transactions []InventoryTransaction `json:"transactions,omitempty" gorm:"foreignKey:WarehouseID"`
	Counts       []InventoryCount      `json:"counts,omitempty" gorm:"foreignKey:WarehouseID"`
}

// TableName 指定表名
func (InventoryWarehouse) TableName() string {
	return "inventory_warehouses"
}

// InventoryLocation 库位表模型
type InventoryLocation struct {
	ID            string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	WarehouseID   string         `json:"warehouse_id" gorm:"not null;type:varchar(36)"`
	Code          string         `json:"code" gorm:"not null;type:varchar(20)"`
	Name          string         `json:"name" gorm:"not null;type:varchar(100)"`
	Type          string         `json:"type" gorm:"type:varchar(20)"`
	Capacity      float64        `json:"capacity" gorm:"type:decimal(18,4)"`
	UsedCapacity  float64        `json:"used_capacity" gorm:"type:decimal(18,4);default:0"`
	Status        string         `json:"status" gorm:"type:varchar(20);default:'available'"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy     string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Warehouse    InventoryWarehouse    `json:"warehouse,omitempty" gorm:"foreignKey:WarehouseID"`
	OnHandItems  []InventoryOnHand     `json:"on_hand_items,omitempty" gorm:"foreignKey:LocationID"`
	Transactions []InventoryTransaction `json:"transactions,omitempty" gorm:"foreignKey:LocationID"`
}

// TableName 指定表名
func (InventoryLocation) TableName() string {
	return "inventory_locations"
}

// InventoryItem 物料表模型
type InventoryItem struct {
	ID          string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	ItemNo      string         `json:"item_no" gorm:"unique;not null;type:varchar(20)"`
	Name        string         `json:"name" gorm:"not null;type:varchar(100)"`
	Description string         `json:"description" gorm:"type:text"`
	CategoryID  string         `json:"category_id" gorm:"type:varchar(36)"`
	Unit        string         `json:"unit" gorm:"not null;type:varchar(10)"`
	Type        string         `json:"type" gorm:"not null;type:varchar(20)"`
	Status      string         `json:"status" gorm:"type:varchar(20);default:'active'"`
	CreatedBy   string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt   time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy   string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Category     *InventoryItemCategory `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	OnHandItems  []InventoryOnHand      `json:"on_hand_items,omitempty" gorm:"foreignKey:ItemID"`
	Transactions []InventoryTransaction `json:"transactions,omitempty" gorm:"foreignKey:ItemID"`
	CountItems   []InventoryCountItem   `json:"count_items,omitempty" gorm:"foreignKey:ItemID"`
}

// TableName 指定表名
func (InventoryItem) TableName() string {
	return "inventory_items"
}

// InventoryItemCategory 物料类别表模型
type InventoryItemCategory struct {
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
	Parent   *InventoryItemCategory `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children []InventoryItemCategory `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	Items    []InventoryItem         `json:"items,omitempty" gorm:"foreignKey:CategoryID"`
}

// TableName 指定表名
func (InventoryItemCategory) TableName() string {
	return "inventory_item_categories"
}

// InventoryOnHand 库存表模型
type InventoryOnHand struct {
	ID          string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	ItemID      string         `json:"item_id" gorm:"not null;type:varchar(36)"`
	WarehouseID string         `json:"warehouse_id" gorm:"not null;type:varchar(36)"`
	LocationID  string         `json:"location_id" gorm:"type:varchar(36)"`
	Quantity    float64        `json:"quantity" gorm:"not null;type:decimal(18,4)"`
	UnitCost    float64        `json:"unit_cost" gorm:"not null;type:decimal(18,2)"`
	TotalCost   float64        `json:"total_cost" gorm:"not null;type:decimal(18,2)"`
	CreatedBy   string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt   time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy   string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Item      InventoryItem      `json:"item,omitempty" gorm:"foreignKey:ItemID"`
	Warehouse InventoryWarehouse `json:"warehouse,omitempty" gorm:"foreignKey:WarehouseID"`
	Location  *InventoryLocation `json:"location,omitempty" gorm:"foreignKey:LocationID"`
}

// TableName 指定表名
func (InventoryOnHand) TableName() string {
	return "inventory_on_hand"
}

// InventoryTransaction 库存交易表模型
type InventoryTransaction struct {
	ID            string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	TransactionNo string         `json:"transaction_no" gorm:"unique;not null;type:varchar(20)"`
	ItemID        string         `json:"item_id" gorm:"not null;type:varchar(36)"`
	WarehouseID   string         `json:"warehouse_id" gorm:"not null;type:varchar(36)"`
	LocationID    string         `json:"location_id" gorm:"type:varchar(36)"`
	Type          string         `json:"type" gorm:"not null;type:varchar(20)"`
	Quantity      float64        `json:"quantity" gorm:"not null;type:decimal(18,4)"`
	UnitCost      float64        `json:"unit_cost" gorm:"not null;type:decimal(18,2)"`
	TotalCost     float64        `json:"total_cost" gorm:"not null;type:decimal(18,2)"`
	ReferenceType string         `json:"reference_type" gorm:"type:varchar(50)"`
	ReferenceID   string         `json:"reference_id" gorm:"type:varchar(36)"`
	TransactionDate time.Time    `json:"transaction_date" gorm:"not null"`
	Remarks       string         `json:"remarks" gorm:"type:text"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Item      InventoryItem      `json:"item,omitempty" gorm:"foreignKey:ItemID"`
	Warehouse InventoryWarehouse `json:"warehouse,omitempty" gorm:"foreignKey:WarehouseID"`
	Location  *InventoryLocation `json:"location,omitempty" gorm:"foreignKey:LocationID"`
}

// TableName 指定表名
func (InventoryTransaction) TableName() string {
	return "inventory_transactions"
}

// InventoryCount 库存盘点表模型
type InventoryCount struct {
	ID            string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	CountNo       string         `json:"count_no" gorm:"unique;not null;type:varchar(20)"`
	WarehouseID   string         `json:"warehouse_id" gorm:"not null;type:varchar(36)"`
	CountDate     time.Time      `json:"count_date" gorm:"not null;type:date"`
	Status        string         `json:"status" gorm:"type:varchar(20);default:'pending'"`
	TotalItems    int            `json:"total_items" gorm:"not null;type:int"`
	TotalVariance float64        `json:"total_variance" gorm:"type:decimal(18,2);default:0"`
	Remarks       string         `json:"remarks" gorm:"type:text"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy     string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Warehouse InventoryWarehouse `json:"warehouse,omitempty" gorm:"foreignKey:WarehouseID"`
	Items     []InventoryCountItem `json:"items,omitempty" gorm:"foreignKey:CountID"`
}

// TableName 指定表名
func (InventoryCount) TableName() string {
	return "inventory_counts"
}

// InventoryCountItem 库存盘点明细表模型
type InventoryCountItem struct {
	ID            string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	CountID       string         `json:"count_id" gorm:"not null;type:varchar(36)"`
	ItemID        string         `json:"item_id" gorm:"not null;type:varchar(36)"`
	LocationID    string         `json:"location_id" gorm:"type:varchar(36)"`
	SystemQuantity float64       `json:"system_quantity" gorm:"not null;type:decimal(18,4)"`
	ActualQuantity float64       `json:"actual_quantity" gorm:"not null;type:decimal(18,4)"`
	VarianceQuantity float64     `json:"variance_quantity" gorm:"not null;type:decimal(18,4)"`
	UnitCost      float64        `json:"unit_cost" gorm:"not null;type:decimal(18,2)"`
	VarianceAmount float64       `json:"variance_amount" gorm:"not null;type:decimal(18,2)"`
	Remarks       string         `json:"remarks" gorm:"type:text"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy     string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Count     InventoryCount   `json:"count,omitempty" gorm:"foreignKey:CountID"`
	Item      InventoryItem    `json:"item,omitempty" gorm:"foreignKey:ItemID"`
	Location  *InventoryLocation `json:"location,omitempty" gorm:"foreignKey:LocationID"`
}

// TableName 指定表名
func (InventoryCountItem) TableName() string {
	return "inventory_count_items"
}
