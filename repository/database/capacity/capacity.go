package capacity

import (
	"time"

	"gorm.io/gorm"
)

type Capacity struct {
	ID        uint           `gorm:"primary_key"`
	CreatedAt time.Time      `gorm:"type:datetime(3)"`
	UpdatedAt time.Time      `gorm:"type:datetime(3)"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(3)"`
}
