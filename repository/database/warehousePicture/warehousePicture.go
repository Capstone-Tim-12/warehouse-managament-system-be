package warehousepicture

import (
	"time"

	"gorm.io/gorm"
)

type WarehousePicture struct {
	ID          uint   `gorm:"primaryKey"`
	Picture     string `gorm:"type:varchar(255)"`
	WarehouseID uint
	CreatedAt   time.Time `gorm:"type:datetime(3)"`
	UpdatedAt   time.Time `gorm:"type:datetime(3)"`
	DeletedAt   gorm.DeletedAt
}
