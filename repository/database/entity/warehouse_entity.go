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
	ID           int `gorm:"primarykey"`
	Name         string
	Longitude    float64
	Latitude     float64
	ProvinceID   int
	RegencyID    int
	DistrictID   int
	Address      string
	BuildingArea float64
	SurfaceArea  float64
	Owner        string
	PhoneNumber  string
	Price        float64
	Description  string
	Status       WarehouseStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	WareHouseImg []WarehouseImg
}

type WarehouseImg struct {
	ID          int `gorm:"primarykey"`
	Image       string
	WarehouseID int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
