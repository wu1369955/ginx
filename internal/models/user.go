package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Username  string         `json:"username" gorm:"size:50;not null;unique"`
	Email     string         `json:"email" gorm:"size:100;not null;unique"`
	Password  string         `json:"password" gorm:"size:100;not null"`
	Nickname  string         `json:"nickname" gorm:"size:50"`
	Avatar    string         `json:"avatar" gorm:"size:255"`
	Role      string         `json:"role" gorm:"size:20;default:'user'"`
}
