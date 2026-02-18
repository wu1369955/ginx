package migration

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wu136995/ginx/internal/data/storage"
)

// Migrator 数据迁移服务
type Migrator struct {
	hotStorage  storage.HotStorage
	coldStorage storage.ColdStorage
	hotThreshold  int64 // 热数据阈值（秒）
	coldThreshold int64 // 冷数据阈值（秒）
	interval      time.Duration // 迁移检查间隔
}

// NewMigrator 创建数据迁移服务实例
func NewMigrator(
	hotStorage storage.HotStorage,
	coldStorage storage.ColdStorage,
	hotThreshold int64,
	coldThreshold int64,
	interval time.Duration,
) *Migrator {
	return &Migrator{
		hotStorage:  hotStorage,
		coldStorage: coldStorage,
		hotThreshold:  hotThreshold,
		coldThreshold: coldThreshold,
		interval:      interval,
	}
}

// Start 启动数据迁移服务
func (m *Migrator) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("数据迁移服务已停止")
				return
			case <-time.After(m.interval):
				if err := m.Migrate(); err != nil {
					log.Printf("数据迁移失败: %v", err)
				}
			}
		}
	}()

	log.Println("数据迁移服务已启动")
}

// Migrate 执行数据迁移
func (m *Migrator) Migrate() error {
	ctx := context.Background()

	// 1. 将冷数据从热存储迁移到冷存储
	if err := m.migrateHotToCold(ctx); err != nil {
		return fmt.Errorf("迁移热数据到冷存储失败: %w", err)
	}

	// 2. 将热数据从冷存储迁移到热存储
	if err := m.migrateColdToHot(ctx); err != nil {
		return fmt.Errorf("迁移冷数据到热存储失败: %w", err)
	}

	return nil
}

// migrateHotToCold 迁移热存储中的冷数据到冷存储
func (m *Migrator) migrateHotToCold(ctx context.Context) error {
	// 查找热存储中的冷数据
	coldData, err := m.hotStorage.FindColdData(ctx, m.coldThreshold)
	if err != nil {
		return err
	}

	for _, data := range coldData {
		// 保存到冷存储
		if err := m.coldStorage.Save(ctx, data); err != nil {
			log.Printf("保存数据到冷存储失败: %v", err)
			continue
		}

		// 从热存储删除
		if err := m.hotStorage.Delete(ctx, data.GetID()); err != nil {
			log.Printf("从热存储删除数据失败: %v", err)
			// 继续处理下一个数据
			continue
		}

		log.Printf("数据 %s 已从热存储迁移到冷存储", data.GetID())
	}

	return nil
}

// migrateColdToHot 迁移冷存储中的热数据到热存储
func (m *Migrator) migrateColdToHot(ctx context.Context) error {
	// 查找冷存储中的热数据
	hotData, err := m.coldStorage.FindHotData(ctx, m.hotThreshold)
	if err != nil {
		return err
	}

	for _, data := range hotData {
		// 保存到热存储
		if err := m.hotStorage.Save(ctx, data); err != nil {
			log.Printf("保存数据到热存储失败: %v", err)
			continue
		}

		// 从冷存储删除
		if err := m.coldStorage.Delete(ctx, data.GetID()); err != nil {
			log.Printf("从冷存储删除数据失败: %v", err)
			// 继续处理下一个数据
			continue
		}

		log.Printf("数据 %s 已从冷存储迁移到热存储", data.GetID())
	}

	return nil
}
