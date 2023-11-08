package capacity

import (
	"time"

	"gorm.io/gorm"
)

type Capacity struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"type:datetime(3)"`
	UpdatedAt time.Time `gorm:"type:datetime(3)"`
	DeletedAt gorm.DeletedAt
}
