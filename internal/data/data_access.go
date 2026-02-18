package data

import (
	"context"

	"github.com/wu136995/ginx/internal/data/storage"
)

// Accessor 数据访问器
type Accessor struct {
	hotStorage  storage.HotStorage
	coldStorage storage.ColdStorage
}

// NewAccessor 创建数据访问器实例
func NewAccessor(
	hotStorage storage.HotStorage,
	coldStorage storage.ColdStorage,
) *Accessor {
	return &Accessor{
		hotStorage:  hotStorage,
		coldStorage: coldStorage,
	}
}

// Get 获取数据
func (a *Accessor) Get(ctx context.Context, id string) (storage.Data, error) {
	// 1. 先从热存储获取
	data, err := a.hotStorage.Get(ctx, id)
	if err == nil {
		return data, nil
	}

	// 2. 热存储未找到，从冷存储获取
	data, err = a.coldStorage.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// 3. 冷数据预热到热存储
	go func() {
		_ = a.hotStorage.Save(ctx, data)
	}()

	return data, nil
}

// Save 保存数据
func (a *Accessor) Save(ctx context.Context, data storage.Data) error {
	// 新数据默认保存到热存储
	return a.hotStorage.Save(ctx, data)
}

// Delete 删除数据
func (a *Accessor) Delete(ctx context.Context, id string) error {
	// 同时从热存储和冷存储删除
	err1 := a.hotStorage.Delete(ctx, id)
	err2 := a.coldStorage.Delete(ctx, id)

	// 如果两个存储都返回错误，返回第一个错误
	if err1 != nil && err2 != nil {
		return err1
	}

	// 如果只有一个存储返回错误，返回nil（因为数据可能只存在于一个存储中）
	return nil
}

// List 列出所有数据
func (a *Accessor) List(ctx context.Context) ([]storage.Data, error) {
	// 从热存储获取数据
	hotData, err := a.hotStorage.List(ctx)
	if err != nil {
		hotData = []storage.Data{}
	}

	// 从冷存储获取数据
	coldData, err := a.coldStorage.List(ctx)
	if err != nil {
		coldData = []storage.Data{}
	}

	// 合并数据
	allData := append(hotData, coldData...)
	return allData, nil
}
