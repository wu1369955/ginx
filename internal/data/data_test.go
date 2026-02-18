package data

import (
	"context"
	"testing"
	"time"

	"github.com/wu136995/ginx/internal/data/storage"
)

func TestDataAccessor(t *testing.T) {
	// 创建热存储
	hotStorage := storage.NewMemoryStorage()

	// 创建冷存储
	coldStorage, err := storage.NewFileStorage("./test_cold_data")
	if err != nil {
		t.Fatalf("创建冷存储失败: %v", err)
	}

	// 创建数据访问器
	accessor := NewAccessor(hotStorage, coldStorage)

	// 测试保存数据
	t.Run("SaveData", func(t *testing.T) {
		ctx := context.Background()
		data := &storage.MemoryData{
			ID:          "test-data-1",
			CreatedAt:   time.Now().Unix(),
			AccessCount: 0,
			Data: map[string]interface{}{
				"key": "value",
			},
		}

		err := accessor.Save(ctx, data)
		if err != nil {
			t.Fatalf("保存数据失败: %v", err)
		}
	})

	// 测试获取数据
	t.Run("GetData", func(t *testing.T) {
		ctx := context.Background()
		data, err := accessor.Get(ctx, "test-data-1")
		if err != nil {
			t.Fatalf("获取数据失败: %v", err)
		}

		if data.GetID() != "test-data-1" {
			t.Errorf("数据ID不匹配，期望: test-data-1, 实际: %s", data.GetID())
		}
	})

	// 测试列出所有数据
	t.Run("ListData", func(t *testing.T) {
		ctx := context.Background()
		dataList, err := accessor.List(ctx)
		if err != nil {
			t.Fatalf("列出数据失败: %v", err)
		}

		if len(dataList) == 0 {
			t.Error("数据列表为空")
		}
	})

	// 测试删除数据
	t.Run("DeleteData", func(t *testing.T) {
		ctx := context.Background()
		err := accessor.Delete(ctx, "test-data-1")
		if err != nil {
			t.Fatalf("删除数据失败: %v", err)
		}

		// 验证数据已删除
		_, err = accessor.Get(ctx, "test-data-1")
		if err == nil {
			t.Error("数据应该已删除，但仍能获取到")
		}
	})
}
