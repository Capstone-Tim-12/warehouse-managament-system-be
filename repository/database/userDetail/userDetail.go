package userdetail

import (
	"time"

	"gorm.io/gorm"
)

type UserDetail struct {
	ID           uint      `gorm:"primaryKey"`
	Photo        string    `gorm:"type:varchar(255)"`
	NIK          string    `gorm:"type:varchar(16)"`
	FullName     string    `gorm:"type:varchar(100)"`
	Gender       string    `gorm:"type:varchar(200)"`
	PlaceOfBirth time.Time `gorm:"type:date"`
	Works        string    `gorm:"type:varchar(100)"`
	Citizenship  string    `gorm:"type:varchar(100)"`
	UserID       uint
	ProvinceID   uint
	RegencyID    uint
	DistrictID   uint
	CreatedAt    time.Time `gorm:"type:datetime(3)"`
	UpdatedAt    time.Time `gorm:"type:datetime(3)"`
	DeletedAt    gorm.DeletedAt
}
