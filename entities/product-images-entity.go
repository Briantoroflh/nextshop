package entities

import (
	"time"
)

type ProductImages struct {
	ID        int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID int64      `json:"product_id" gorm:"not null;index"`
	Product   Products   `json:"product,omitempty" gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE"`
	ImageURL  string     `json:"image_url" gorm:"size:255;not null"`
	IsPrimary bool       `json:"is_primary" gorm:"default:false;not null"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
