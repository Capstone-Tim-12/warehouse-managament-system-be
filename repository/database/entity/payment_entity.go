package entity

import (
	"time"

	"gorm.io/gorm"
)

type PaymentScheme struct {
	ID        int `gorm:"primarykey"`
	Scheme    string
	Ration    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
