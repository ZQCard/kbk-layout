package data

import (
	"time"

	"gorm.io/gorm"
)

type BaseFields struct {
	Id        int64          `gorm:"primarykey;type:int;comment:主键id"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;comment:创建时间"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间"`
}

// Paginate 分页
func Paginate(page, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
