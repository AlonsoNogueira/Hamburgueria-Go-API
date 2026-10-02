package model

import (
	"time"

	"github.com/google/uuid"
)

type Snacks struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"type: text"`
	Price       float64   `gorm:"not null"`

	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (Snacks) TableName() string {
	return "snacks"
}
