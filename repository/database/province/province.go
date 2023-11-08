package province

import "github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/regency"

type Province struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255)"`
	Regencies []regency.Regency
}
