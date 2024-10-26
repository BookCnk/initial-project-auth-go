package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserAuth struct {
	ID        string         `json:"id" gorm:"type:varchar(255);primarykey"`
	UserID    string         `json:"user_id" gorm:"type:varchar(255)"`
	User      User           `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Salt      string         `json:"salt" gorm:"type:varchar(255)"`
	Hash      string         `json:"hash" gorm:"type:varchar(255)"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
