package entities

import (
	"time"
)

type ShippingMethods struct {
	ID            int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string     `json:"name" gorm:"size:191;not null"`
	Description   *string    `json:"description" gorm:"type:text"`
	Cost          float64    `json:"cost" gorm:"type:decimal(12,2);default:0.00;not null"`
	EstimatedDays *int       `json:"estimated_days"`
	IsActive      bool       `json:"is_active" gorm:"default:true;not null"`
	Orders        []Orders   `json:"orders,omitempty" gorm:"foreignKey:ShippingMethodID;references:ID"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	CreatedAt     time.Time  `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time  `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
