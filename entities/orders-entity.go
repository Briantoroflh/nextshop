package entities

import (
	"time"
)

type OrderStatus string
type PaymentStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

const (
	PaymentStatusPending  PaymentStatus = "pending"
	PaymentStatusSuccess  PaymentStatus = "success"
	PaymentStatusFailed   PaymentStatus = "failed"
	PaymentStatusRefunded PaymentStatus = "refunded"
)

type Orders struct {
	ID               int64            `json:"id" gorm:"primaryKey;autoIncrement"`
	BuyerID          int64            `json:"buyer_id" gorm:"not null;index"`
	Buyer            Users            `json:"buyer,omitempty" gorm:"foreignKey:BuyerID;references:ID;constraint:OnDelete:CASCADE"`
	AddressID        *int64           `json:"address_id" gorm:"index"`
	Address          *Addresses       `json:"address,omitempty" gorm:"foreignKey:AddressID;references:ID;constraint:OnDelete:SET NULL"`
	PaymentMethodID  *int64           `json:"payment_method_id" gorm:"index"`
	PaymentMethod    *PaymentMethods  `json:"payment_method,omitempty" gorm:"foreignKey:PaymentMethodID;references:ID;constraint:OnDelete:SET NULL"`
	ShippingMethodID *int64           `json:"shipping_method_id" gorm:"index"`
	ShippingMethod   *ShippingMethods `json:"shipping_method,omitempty" gorm:"foreignKey:ShippingMethodID;references:ID;constraint:OnDelete:SET NULL"`
	TotalAmount      float64          `json:"total_amount" gorm:"type:decimal(12,2);not null"`
	OrderStatus      OrderStatus      `json:"order_status" gorm:"type:order_status;default:pending;not null"`
	PaymentStatus    PaymentStatus    `json:"payment_status" gorm:"type:payment_status;default:pending;not null"`
	Note             *string          `json:"note" gorm:"type:text"`
	OrderItems       []OrderItems     `json:"order_items,omitempty" gorm:"foreignKey:OrderID;references:ID"`
	DeletedAt        *time.Time       `json:"deleted_at,omitempty" gorm:"index"`
	CreatedAt        time.Time        `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time        `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
