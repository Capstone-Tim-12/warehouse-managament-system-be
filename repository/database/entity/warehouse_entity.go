package entity

import (
	"time"

	"gorm.io/gorm"
)

type WarehouseStatus string

const (
	Available    WarehouseStatus = "tersedia"
	NotAvailable WarehouseStatus = "tidak tersedia"
)

type Warehouse struct {
	ID              int `gorm:"primarykey"`
	Name            string
	Longitude       float64
	Latitude        float64
	DistrictID      string
	District        District `gorm:"foreignKey:DistrictID"`
	Address         string
	BuildingArea    float64
	SurfaceArea     float64
	Owner           string
	PhoneNumber     string
	Price           float64
	Description     string
	WarehouseTypeID int
	WarehouseType   WarehouseType `gorm:"foreignKey:WarehouseTypeID"`
	Status          WarehouseStatus
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	WarehouseImg    []WarehouseImg
}

type WarehouseImg struct {
	ID          int `gorm:"primarykey"`
	Image       string
	WarehouseID int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type WarehouseType struct {
	ID        int `gorm:"primarykey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Warehouse []Warehouse
}
