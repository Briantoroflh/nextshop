package entities

import (
	"time"
)

type UserRole string

const (
	UserRoleBuyer  UserRole = "buyer"
	UserRoleSeller UserRole = "seller"
	UserRoleAdmin  UserRole = "admin"
)

type Users struct {
	ID           int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	FullName     *string    `json:"full_name" gorm:"size:191"`
	Email        string     `json:"email" gorm:"uniqueIndex;size:191;not null"`
	Phone        *string    `json:"phone" gorm:"size:32"`
	PasswordHash string     `json:"-" gorm:"size:191;not null"`
	Role         UserRole   `json:"role" gorm:"type:user_role;default:buyer;not null"`
	IsActive     bool       `json:"is_active" gorm:"default:true;not null"`
	LastLoginAt  *time.Time `json:"last_login_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	CreatedAt    time.Time  `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
