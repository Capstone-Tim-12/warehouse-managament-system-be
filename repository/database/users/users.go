package users

import (
	"time"

	userdetail "github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/userDetail"
	"gorm.io/gorm"
)

type User struct {
	ID               uint   `gorm:"primaryKey"`
	Username         string `gorm:"type:varchar(35)"`
	Email            string `gorm:"type:varchar(100)"`
	IsVerifyAcount   bool
	IsVerifyIdentity bool
	Password         string    `gorm:"type:text"`
	CreatedAt        time.Time `gorm:"type:datetime(3)"`
	UpdatedAt        time.Time `gorm:"type:datetime(3)"`
	DeletedAt        gorm.DeletedAt
	UserDetails      userdetail.UserDetail
}
