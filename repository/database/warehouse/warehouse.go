package warehouse

import (
	"time"

	warehousepicture "github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/warehousePicture"
	"gorm.io/gorm"
)

type Warehouse struct {
	ID                uint   `gorm:"primaryKey"`
	Name              string `gorm:"type:varchar(255)"`
	ProvinceID        uint
	RegencyID         uint
	Wide              string `gorm:"type:varchar(200)"`
	Owner             string `gorm:"type:varchar(100)"`
	PhoneNumber       string `gorm:"type:varchar(15)"`
	Price             int64
	Status            string    `gorm:"type:enum('tesedia', 'disewa', 'dalam pemeliharaan')"`
	Description       string    `gorm:"type:text"`
	CreatedAt         time.Time `gorm:"type:datetime(3)"`
	UpdatedAt         time.Time `gorm:"type:datetime(3)"`
	DeletedAt         gorm.DeletedAt
	WarehousePictures []warehousepicture.WarehousePicture
}
