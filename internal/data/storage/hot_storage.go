package storage

import (
	"context"
	"sync"
	"time"
)

// 内存数据实现
type MemoryData struct {
	ID           string
	CreatedAt    int64
	AccessCount  int
	Data         map[string]interface{}
	lastAccessed int64
}

// GetID 获取数据ID
func (d *MemoryData) GetID() string {
	return d.ID
}

// GetCreatedAt 获取创建时间戳
func (d *MemoryData) GetCreatedAt() int64 {
	return d.CreatedAt
}

// GetAccessCount 获取访问次数
func (d *MemoryData) GetAccessCount() int {
	return d.AccessCount
}

// IncrementAccessCount 增加访问次数
func (d *MemoryData) IncrementAccessCount() {
	d.AccessCount++
	d.lastAccessed = time.Now().Unix()
}

// MemoryStorage 内存存储实现
type MemoryStorage struct {
	data map[string]*MemoryData
	mu   sync.RWMutex
}

// NewMemoryStorage 创建内存存储实例
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]*MemoryData),
	}
}

// Get 获取数据
func (s *MemoryStorage) Get(ctx context.Context, id string) (Data, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, ok := s.data[id]
	if !ok {
		return nil, ErrNotFound
	}

	// 增加访问次数
	data.IncrementAccessCount()

	return data, nil
}

// Save 保存数据
func (s *MemoryStorage) Save(ctx context.Context, data Data) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 检查数据类型
	memData, ok := data.(*MemoryData)
	if !ok {
		// 转换为MemoryData
		memData = &MemoryData{
			ID:          data.GetID(),
			CreatedAt:   data.GetCreatedAt(),
			AccessCount: data.GetAccessCount(),
			Data:        make(map[string]interface{}),
		}
	}

	s.data[memData.ID] = memData
	return nil
}

// Delete 删除数据
func (s *MemoryStorage) Delete(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[id]; !ok {
		return ErrNotFound
	}

	delete(s.data, id)
	return nil
}

// List 列出所有数据
func (s *MemoryStorage) List(ctx context.Context) ([]Data, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []Data
	for _, data := range s.data {
		result = append(result, data)
	}

	return result, nil
}

// FindColdData 查找冷数据
func (s *MemoryStorage) FindColdData(ctx context.Context, threshold int64) ([]Data, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []Data
	now := time.Now().Unix()

	for _, data := range s.data {
		// 检查最后访问时间是否超过阈值
		if now-data.lastAccessed > threshold {
			result = append(result, data)
		}
	}

	return result, nil
}

// FindHotData 查找热数据
func (s *MemoryStorage) FindHotData(ctx context.Context, threshold int64) ([]Data, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []Data
	now := time.Now().Unix()

	for _, data := range s.data {
		// 检查最后访问时间是否在阈值内
		if now-data.lastAccessed <= threshold {
			result = append(result, data)
		}
	}

	return result, nil
}

// 确保MemoryStorage实现了HotStorage接口
var _ HotStorage = (*MemoryStorage)(nil)
