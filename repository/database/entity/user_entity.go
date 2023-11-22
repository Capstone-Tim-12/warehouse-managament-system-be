package entity

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	ID               int `gorm:"primarykey"`
	Username         string
	Email            string
	IsVerifyAccount  bool
	IsVerifyIdentity bool
	Password         string
	Role             Role
	Photo            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	UserDetail       UserDetail
}

type UserDetail struct {
	ID           int `gorm:"primarykey"`
	Address      string
	Country      string
	NIK          string
	FullName     string
	Gender       string
	PlaceOfBirth string
	DateBirth    time.Time
	Work         string
	Citizenship  string
	UserID       int
	ProvinceID   string
	Province     Province `gorm:"foreignKey:ProvinceID"`
	RegencyID    string
	Regency      Regency `gorm:"foreignKey:RegencyID"`
	DistrictID   string
	District     District `gorm:"foreignKey:DistrictID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
