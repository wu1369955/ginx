package services

import (
	"errors"

	"github.com/wu136995/ginx/internal/models"
	"gorm.io/gorm"
)

// UserService 用户服务接口
type UserService interface {
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

// userService 用户服务实现
type userService struct {
	db *gorm.DB
}

// NewUserService 创建用户服务实例
func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

// GetUserByID 根据ID获取用户
func (s *userService) GetUserByID(id string) (*models.User, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取用户
	var user models.User
	result := s.db.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// CreateUser 创建用户
func (s *userService) CreateUser(user *models.User) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 保存到数据库
	result := s.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdateUser 更新用户
func (s *userService) UpdateUser(user *models.User) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 保存到数据库
	result := s.db.Save(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除用户
	result := s.db.Delete(&models.User{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
