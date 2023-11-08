package district

import "github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/village"

type District struct {
	ID        uint `gorm:"primaryKey"`
	RegencyID uint
	Name      string `gorm:"type:varchar(255)"`
	Villages  []village.Village
}
