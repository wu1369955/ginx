package models

import (
	"time"
)

// ProductionOrder 生产订单模型
type ProductionOrder struct {
	ID           string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	OrderNo      string    `json:"order_no" gorm:"unique;not null;type:varchar(20)"`
	ProductName  string    `json:"product_name" gorm:"not null;type:varchar(100)"`
	Quantity     int       `json:"quantity" gorm:"not null"`
	Status       string    `json:"status" gorm:"not null;type:varchar(20)"` // pending, approved, in_progress, completed, cancelled
	Priority     string    `json:"priority" gorm:"not null;type:varchar(20)"` // high, medium, low
	StartDate    time.Time `json:"start_date" gorm:"not null"`
	EndDate      time.Time `json:"end_date" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null"`
	CreatedBy    string    `json:"created_by" gorm:"not null;type:varchar(50)"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"not null"`
	UpdatedBy    string    `json:"updated_by" gorm:"not null;type:varchar(50)"`
}

// TableName 指定表名
func (ProductionOrder) TableName() string {
	return "production_orders"
}

// ProductionTicket 生产工单模型
type ProductionTicket struct {
	ID             string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	TicketNo       string    `json:"ticket_no" gorm:"unique;not null;type:varchar(20)"`
	ProductionOrderID string  `json:"production_order_id" gorm:"not null;type:varchar(36)"`
	ProductName    string    `json:"product_name" gorm:"not null;type:varchar(100)"`
	Quantity       int       `json:"quantity" gorm:"not null"`
	Status         string    `json:"status" gorm:"not null;type:varchar(20)"` // pending, in_progress, completed, cancelled
	WorkCenterID   string    `json:"work_center_id" gorm:"not null;type:varchar(36)"`
	StartTime      time.Time `json:"start_time" gorm:"not null"`
	EndTime        time.Time `json:"end_time" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"not null"`
	CreatedBy      string    `json:"created_by" gorm:"not null;type:varchar(50)"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"not null"`
	UpdatedBy      string    `json:"updated_by" gorm:"not null;type:varchar(50)"`
}

// TableName 指定表名
func (ProductionTicket) TableName() string {
	return "production_tickets"
}

// Routing 工艺路线模型
type Routing struct {
	ID           string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	RoutingNo    string    `json:"routing_no" gorm:"unique;not null;type:varchar(20)"`
	ProductName  string    `json:"product_name" gorm:"not null;type:varchar(100)"`
	Description  string    `json:"description" gorm:"type:text"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null"`
	CreatedBy    string    `json:"created_by" gorm:"not null;type:varchar(50)"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"not null"`
	UpdatedBy    string    `json:"updated_by" gorm:"not null;type:varchar(50)"`
}

// TableName 指定表名
func (Routing) TableName() string {
	return "routings"
}

// WorkCenter 工作中心模型
type WorkCenter struct {
	ID           string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	WorkCenterNo string    `json:"work_center_no" gorm:"unique;not null;type:varchar(20)"`
	Name         string    `json:"name" gorm:"not null;type:varchar(100)"`
	Description  string    `json:"description" gorm:"type:text"`
	Capacity     int       `json:"capacity" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null"`
	CreatedBy    string    `json:"created_by" gorm:"not null;type:varchar(50)"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"not null"`
	UpdatedBy    string    `json:"updated_by" gorm:"not null;type:varchar(50)"`
}

// TableName 指定表名
func (WorkCenter) TableName() string {
	return "work_centers"
}
