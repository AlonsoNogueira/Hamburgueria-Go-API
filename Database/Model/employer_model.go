package model

import (
	"time"

	uuid "github.com/google/uuid"
)

type Employer struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name  string    `gorm:"not null"`
	Age   string    `gorm:"not null"`
	Email string    `gorm:"not null;unique"`
	Phone string    `gorm:"not null"`

	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (Employer) TableName() string {
	return "employers"
}
