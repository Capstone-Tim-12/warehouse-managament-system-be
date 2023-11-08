package regency

import "github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/district"

type Regency struct {
	ID         uint `gorm:"primaryKey"`
	ProvinceID uint
	Name       string `gorm:"type:varchar(255)"`
	Districts  []district.District
}
