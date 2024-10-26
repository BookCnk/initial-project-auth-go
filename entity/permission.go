package entity

import (
	"gorm.io/gorm"
	"time"
)

type Permission struct {
	ID        string         `json:"id" gorm:"type:varchar(255);primarykey"`
	Name      string         `json:"name" gorm:"type:varchar(255);"`
	IsActive  bool           `json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
