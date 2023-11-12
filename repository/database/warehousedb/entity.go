package warehousedb

import "time"

type WarehouseImg struct {
	ID          int `gorm:"primarykey"`
	Image       string
	WarehouseID int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type WarehousePaymentType struct {
	ID          int  `gorm:"primarykey"`
	Weekly      bool `gorm:"NOT NULL"`
	Mountly     bool `gorm:"NOT NULL"`
	Annual      bool `gorm:"NOT NULL"`
	WarehouseID int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type Warehouse struct {
	ID                   int `gorm:"primarykey"`
	Name                 string
	Longitude            float64
	Latitude             float64
	BuildingArea         float64
	SurfaceArea          float64
	Owner                string
	PhoneNumber          string
	Price                float64
	Description          string
	ProvinceID           string
	RegencyID            string
	DistrictID           string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            time.Time
	WareHouseImg         []WarehouseImg
	WarehousePaymentType WarehousePaymentType
}
