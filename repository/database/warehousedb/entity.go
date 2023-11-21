package warehousedb

import (
	"time"

	"gorm.io/gorm"
)

type WarehouseStatus string

const (
	Available    WarehouseStatus = "tersedia"
	NotAvailable WarehouseStatus = "tidak tersedia"
)

type PaymentScheme struct {
	ID        int `gorm:"primarykey"`
	Scheme    string
	Ration    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Warehouse []Warehouse
}

type Warehouse struct {
	ID              int `gorm:"primarykey"`
	Name            string
	Longitude       float64
	Latitude        float64
	LocationName    string
	BuildingArea    float64
	SurfaceArea     float64
	Owner           string
	PhoneNumber     string
	Price           float64
	Description     string
	PaymentSchemeID int
	Status          WarehouseStatus
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	WareHouseImg    []WareHouseImg
}

type WareHouseImg struct {
	ID          int `gorm:"primarykey"`
	Image       string
	WarehouseID int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
