package province

import (
	"time"

	"gorm.io/gorm"
)

type Province struct {
	ID        uint           `gorm:"primary_key"`
	Name      string         `gorm:"type:varchar(255)"`
	CreatedAt time.Time      `gorm:"type:datetime(3)"`
	UpdatedAt time.Time      `gorm:"type:datetime(3)"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(3)"`
}
