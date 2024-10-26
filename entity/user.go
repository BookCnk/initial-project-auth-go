package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string         `json:"id" gorm:"type:varchar(255);primarykey"`
	Name      EncryptedField `json:"name"`
	Email     EncryptedField `json:"email"`
	RoleID    string         `json:"role_id" gorm:"type:varchar(255);"`
	Role      Role           `json:"role" gorm:"foreignKey:RoleID;references:ID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type CreateUserReq struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type SignInRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ResponseMe struct {
	Email       string   `json:"email"`
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	Role        string   `json:"role"`
}

type ResponseAllUser struct {
	Email    string `json:"email"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	RoleName string `json:"role_name"`
}

type UpdateUserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
