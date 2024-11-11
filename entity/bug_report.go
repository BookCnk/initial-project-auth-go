package entity

import "time"

type (
	BugReport struct {
		BugID       uint      `json:"bug_id" gorm:"primaryKey;autoIncrement"`
		Title       string    `json:"title" gorm:"size:150;not null"`
		Description string    `json:"description" gorm:"not null"`
		CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
		UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
		ProjectID   uint      `json:"project_id"`
		UserID      string    `json:"user_id" gorm:"type:varchar(255);"`
		User        User      `json:"-" gorm:"foreignKey:UserID"`
	}
)
