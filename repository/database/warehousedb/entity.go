package warehousedb

import (
	"time"

	"gorm.io/gorm"
)

type WareHouseImg struct {
	ID        int `gorm:"primarykey"`
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Warehouse struct {
	ID             int `gorm:"primarykey"`
	Name           string
	Longitude      float64
	Latitude       float64
	LocationName   string
	BuildingArea   float64
	SurfaceArea    float64
	Owner          string
	PhoneNumber    string
	Price          float64
	Description    string
	PaymentWeekly  bool
	PaymentMountly bool
	PaymentAnnual  bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	WareHouseImg   []WareHouseImg
}
