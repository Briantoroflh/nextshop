package entities

import (
	"time"
)

type OrderItems struct {
	ID        int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID   int64      `json:"order_id" gorm:"not null;index"`
	Order     Orders     `json:"order,omitempty" gorm:"foreignKey:OrderID;references:ID;constraint:OnDelete:CASCADE"`
	ProductID *int64     `json:"product_id" gorm:"index"`
	Product   *Products  `json:"product,omitempty" gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:SET NULL"`
	Quantity  int        `json:"quantity" gorm:"not null"`
	UnitPrice float64    `json:"unit_price" gorm:"type:decimal(12,2);not null"`
	Subtotal  float64    `json:"subtotal" gorm:"type:decimal(12,2);not null"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
