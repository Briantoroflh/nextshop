package entities

import (
	"time"
)

type Products struct {
	ID          int64           `json:"id" gorm:"primaryKey;autoIncrement"`
	SellerID    int64           `json:"seller_id" gorm:"not null;index"`
	Seller      Users           `json:"seller,omitempty" gorm:"foreignKey:SellerID;references:ID;constraint:OnDelete:CASCADE"`
	CategoryID  *int64          `json:"category_id" gorm:"index"`
	Category    *Categories     `json:"category,omitempty" gorm:"foreignKey:CategoryID;references:ID;constraint:OnDelete:SET NULL"`
	Name        string          `json:"name" gorm:"size:191;not null"`
	Slug        string          `json:"slug" gorm:"uniqueIndex;size:191;not null"`
	Description *string         `json:"description" gorm:"type:text"`
	Price       float64         `json:"price" gorm:"type:decimal(12,2);not null"`
	Stock       int             `json:"stock" gorm:"default:0;not null"`
	IsActive    bool            `json:"is_active" gorm:"default:true;not null"`
	Images      []ProductImages `json:"images,omitempty" gorm:"foreignKey:ProductID;references:ID"`
	OrderItems  []OrderItems    `json:"order_items,omitempty" gorm:"foreignKey:ProductID;references:ID"`
	Reviews     []Reviews       `json:"reviews,omitempty" gorm:"foreignKey:ProductID;references:ID"`
	DeletedAt   *time.Time      `json:"deleted_at,omitempty" gorm:"index"`
	CreatedAt   time.Time       `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
