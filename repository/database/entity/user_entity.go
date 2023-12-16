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
	Longitude        float64
	Latitude         float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	UserDetail       UserDetail     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Favorit          []Favorit      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Transaction      []Transaction  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
	DistrictID   string
	District     District `gorm:"foreignKey:DistrictID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Avatar struct {
	ID        int `gorm:"primarykey"`
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Favorit struct {
	ID          int `gorm:"primarykey"`
	UserID      int
	User        User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WarehouseID int
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
