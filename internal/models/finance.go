package models

import (
	"time"

	"gorm.io/gorm"
)

// FinanceAccount 会计科目表模型
type FinanceAccount struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Code      string         `json:"code" gorm:"unique;not null;type:varchar(20)"`
	Name      string         `json:"name" gorm:"not null;type:varchar(100)"`
	Type      string         `json:"type" gorm:"not null;type:varchar(20)"`
	Level     int            `json:"level" gorm:"not null;type:int"`
	ParentID  string         `json:"parent_id" gorm:"type:varchar(36)"`
	CreatedBy string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Parent   *FinanceAccount `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children []FinanceAccount `json:"children,omitempty" gorm:"foreignKey:ParentID"`
}

// TableName 指定表名
func (FinanceAccount) TableName() string {
	return "finance_accounts"
}
