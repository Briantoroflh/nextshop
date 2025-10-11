package entities

import (
	"time"
)

type Categories struct {
	ID          int64        `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string       `json:"name" gorm:"size:191;not null"`
	Slug        string       `json:"slug" gorm:"uniqueIndex;size:191;not null"`
	ParentID    *int64       `json:"parent_id" gorm:"index"`
	Parent      *Categories  `json:"parent,omitempty" gorm:"foreignKey:ParentID;references:ID;constraint:OnDelete:SET NULL"`
	Children    []Categories `json:"children,omitempty" gorm:"foreignKey:ParentID;references:ID"`
	Description *string      `json:"description" gorm:"type:text"`
	DeletedAt   *time.Time   `json:"deleted_at,omitempty" gorm:"index"`
	CreatedAt   time.Time    `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
