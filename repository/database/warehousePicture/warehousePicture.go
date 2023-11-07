package warehousepicture

import (
	"time"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/warehouse"
	"gorm.io/gorm"
)

type WarehousePicture struct {
	ID          uint           `gorm:"primary_key"`
	Picture     string         `gorm:"type:varchar(255)"`
	WarehouseID uint           `gorm:"index"`
	CreatedAt   time.Time      `gorm:"type:datetime(3)"`
	UpdatedAt   time.Time      `gorm:"type:datetime(3)"`
	DeletedAt   gorm.DeletedAt `gorm:"type:datetime(3)"`

	Warehouse warehouse.Warehouse `gorm:"foreignKey:WarehouseID"`
}
