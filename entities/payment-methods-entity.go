package entities

import (
	"time"
)

type PaymentMethods struct {
	ID          int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string     `json:"name" gorm:"size:191;not null"`
	Description *string    `json:"description" gorm:"type:text"`
	IsActive    bool       `json:"is_active" gorm:"default:true;not null"`
	Orders      []Orders   `json:"orders,omitempty" gorm:"foreignKey:PaymentMethodID;references:ID"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	CreatedAt   time.Time  `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
