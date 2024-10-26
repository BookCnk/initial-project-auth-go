package entity

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID          string         `json:"id" gorm:"type:varchar(255);primarykey"`
	Name        string         `json:"name" gorm:"type:varchar(255);"`
	Permissions []Permission   `json:"permissions" gorm:"many2many:role_permissions"`
	IsActive    bool           `json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
