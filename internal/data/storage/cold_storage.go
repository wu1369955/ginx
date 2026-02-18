package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 文件数据实现
type FileData struct {
	ID          string                 `json:"id"`
	CreatedAt   int64                  `json:"created_at"`
	AccessCount int                    `json:"access_count"`
	Data        map[string]interface{} `json:"data"`
	lastAccessed int64                 `json:"last_accessed"`
}

// GetID 获取数据ID
func (d *FileData) GetID() string {
	return d.ID
}

// GetCreatedAt 获取创建时间戳
func (d *FileData) GetCreatedAt() int64 {
	return d.CreatedAt
}

// GetAccessCount 获取访问次数
func (d *FileData) GetAccessCount() int {
	return d.AccessCount
}

// IncrementAccessCount 增加访问次数
func (d *FileData) IncrementAccessCount() {
	d.AccessCount++
	d.lastAccessed = time.Now().Unix()
}

// FileStorage 文件存储实现
type FileStorage struct {
	baseDir string
	mu      sync.RWMutex
}

// NewFileStorage 创建文件存储实例
func NewFileStorage(baseDir string) (*FileStorage, error) {
	// 确保目录存在
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create base directory: %w", err)
	}

	return &FileStorage{
		baseDir: baseDir,
	}, nil
}

// Get 获取数据
func (s *FileStorage) Get(ctx context.Context, id string) (Data, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	filePath := filepath.Join(s.baseDir, id+".json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var fileData FileData
	if err := json.Unmarshal(data, &fileData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	// 增加访问次数
	fileData.IncrementAccessCount()

	// 写回文件
	s.mu.RUnlock()
	s.mu.Lock()
	updatedData, err := json.MarshalIndent(fileData, "", "  ")
	if err != nil {
		s.mu.Unlock()
		s.mu.RLock()
		return &fileData, fmt.Errorf("failed to marshal data: %w", err)
	}

	if err := os.WriteFile(filePath, updatedData, 0644); err != nil {
		s.mu.Unlock()
		s.mu.RLock()
		return &fileData, fmt.Errorf("failed to write file: %w", err)
	}
	s.mu.Unlock()
	s.mu.RLock()

	return &fileData, nil
}

// Save 保存数据
func (s *FileStorage) Save(ctx context.Context, data Data) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 检查数据类型
	fileData, ok := data.(*FileData)
	if !ok {
		// 转换为FileData
		fileData = &FileData{
			ID:          data.GetID(),
			CreatedAt:   data.GetCreatedAt(),
			AccessCount: data.GetAccessCount(),
			Data:        make(map[string]interface{}),
			lastAccessed: time.Now().Unix(),
		}
	}

	filePath := filepath.Join(s.baseDir, fileData.ID+".json")
	dataBytes, err := json.MarshalIndent(fileData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if err := os.WriteFile(filePath, dataBytes, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// Delete 删除数据
func (s *FileStorage) Delete(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	filePath := filepath.Join(s.baseDir, id+".json")
	if err := os.Remove(filePath); err != nil {
		if os.IsNotExist(err) {
			return ErrNotFound
		}
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

// List 列出所有数据
func (s *FileStorage) List(ctx context.Context) ([]Data, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	files, err := os.ReadDir(s.baseDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var result []Data
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			id := filepath.Base(file.Name())[:len(filepath.Base(file.Name()))-5] // 移除 .json 后缀
			data, err := s.Get(ctx, id)
			if err == nil {
				result = append(result, data)
			}
		}
	}

	return result, nil
}

// FindColdData 查找冷数据
func (s *FileStorage) FindColdData(ctx context.Context, threshold int64) ([]Data, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	files, err := os.ReadDir(s.baseDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var result []Data
	now := time.Now().Unix()

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(s.baseDir, file.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				continue
			}

			var fileData FileData
			if err := json.Unmarshal(data, &fileData); err != nil {
				continue
			}

			// 检查最后访问时间是否超过阈值
			if now-fileData.lastAccessed > threshold {
				result = append(result, &fileData)
			}
		}
	}

	return result, nil
}

// FindHotData 查找热数据
func (s *FileStorage) FindHotData(ctx context.Context, threshold int64) ([]Data, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	files, err := os.ReadDir(s.baseDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var result []Data
	now := time.Now().Unix()

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(s.baseDir, file.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				continue
			}

			var fileData FileData
			if err := json.Unmarshal(data, &fileData); err != nil {
				continue
			}

			// 检查最后访问时间是否在阈值内
			if now-fileData.lastAccessed <= threshold {
				result = append(result, &fileData)
			}
		}
	}

	return result, nil
}

// 确保FileStorage实现了ColdStorage接口
var _ ColdStorage = (*FileStorage)(nil)
