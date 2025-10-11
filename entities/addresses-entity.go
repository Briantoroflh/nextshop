package entities

import (
	"time"
)

type Addresses struct {
	ID            int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID        int64      `json:"user_id" gorm:"not null;index"`
	User          Users      `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	RecipientName string     `json:"recipient_name" gorm:"size:191;not null"`
	PhoneNumber   string     `json:"phone_number" gorm:"size:32;not null"`
	StreetAddress string     `json:"street_address" gorm:"type:text;not null"`
	City          *string    `json:"city" gorm:"size:191"`
	Province      *string    `json:"province" gorm:"size:191"`
	PostalCode    *string    `json:"postal_code" gorm:"size:32"`
	IsPrimary     bool       `json:"is_primary" gorm:"default:false;not null"`
	Orders        []Orders   `json:"orders,omitempty" gorm:"foreignKey:AddressID;references:ID"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	CreatedAt     time.Time  `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time  `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
