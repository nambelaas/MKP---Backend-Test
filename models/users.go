package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	Admin UserRole = "Admin"
	User  UserRole = "User"
)

type Users struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Email        string    `gorm:"type:text;index:unique;not null" json:"email"`
	PasswordHash string    `gorm:"type:text;not null" json:"password_hash"`
	FullName     string    `gorm:"type:text;not null" json:"full_name"`
	Role         UserRole  `gorm:"type:user_role;not null;default:'User'" json:"role"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}

type UserRegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Role     string `json:"role" binding:"omitempty,oneof=Admin User"`
}

type UserLoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	return
}

func (Users) TableName() string {
	return "tiket_bioskop.users"
}
