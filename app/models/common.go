package models

import (
	"time"

	"gorm.io/gorm"
)

// 自增ID主键
type ID struct {
	ID int `json:"id" gorm:"primaryKey"`
}

// 创建、更新时间
type Timestamps struct {
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

// 软删除
type SoftDeletes struct {
	DeleteAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
