package entities

import (
	"time"
)

type Reviews struct {
	ID        int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID int64      `json:"product_id" gorm:"not null;index"`
	Product   Products   `json:"product,omitempty" gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE"`
	UserID    int64      `json:"user_id" gorm:"not null;index"`
	User      Users      `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Rating    int        `json:"rating" gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Comment   *string    `json:"comment" gorm:"type:text"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
