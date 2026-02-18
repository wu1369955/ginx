package storage

import (
	"context"
	"errors"
)

// 定义错误
var (
	ErrNotFound      = errors.New("data not found")
	ErrInvalidData   = errors.New("invalid data")
	ErrStorageFailed = errors.New("storage operation failed")
)

// Data 数据接口
type Data interface {
	// GetID 获取数据ID
	GetID() string
	// GetCreatedAt 获取创建时间戳
	GetCreatedAt() int64
	// GetAccessCount 获取访问次数
	GetAccessCount() int
	// IncrementAccessCount 增加访问次数
	IncrementAccessCount()
}

// Storage 存储接口
type Storage interface {
	// Get 获取数据
	Get(ctx context.Context, id string) (Data, error)
	// Save 保存数据
	Save(ctx context.Context, data Data) error
	// Delete 删除数据
	Delete(ctx context.Context, id string) error
	// List 列出所有数据
	List(ctx context.Context) ([]Data, error)
	// FindColdData 查找冷数据
	FindColdData(ctx context.Context, threshold int64) ([]Data, error)
	// FindHotData 查找热数据
	FindHotData(ctx context.Context, threshold int64) ([]Data, error)
}

// HotStorage 热存储接口
type HotStorage interface {
	Storage
}

// ColdStorage 冷存储接口
type ColdStorage interface {
	Storage
}
