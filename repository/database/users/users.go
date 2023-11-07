package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               uint      `gorm:"primary_key"`
	Photo            string    `gorm:"type:varchar(255)"`
	Username         string    `gorm:"type:varchar(35)"`
	Email            string    `gorm:"type:varchar(100)"`
	NIK              string    `gorm:"type:varchar(16)"`
	FullName         string    `gorm:"type:varchar(100)"`
	Gender           string    `gorm:"type:varchar(200)"`
	PlaceOfBirth     time.Time `gorm:"type:date"`
	Work             string    `gorm:"type:varchar(100)"`
	Citizenship      string    `gorm:"type:varchar(100)"`
	IsVerifyAccount  bool
	IsVerifyIdentity bool
	Password         string         `gorm:"type:text"`
	CreatedAt        time.Time      `gorm:"type:datetime(3)"`
	UpdatedAt        time.Time      `gorm:"type:datetime(3)"`
	DeletedAt        gorm.DeletedAt `gorm:"type:datetime(3)"`
}
