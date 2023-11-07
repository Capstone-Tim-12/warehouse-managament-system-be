package warehouse

import (
	"time"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/city"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/province"
	"gorm.io/gorm"
)

type Warehouse struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(255)"`
	CityID      uint   `gorm:"index"`
	ProvinceID  uint   `gorm:"index"`
	CityId      int    `gorm:"type:int(10)"`
	Wide        string `gorm:"type:varchar(200)"`
	Owner       string `gorm:"type:varchar(100)"`
	PhoneNumber string `gorm:"type:varchar(15)"`
	Price       int64
	Status      string         `gorm:"type:enum('tesedia','disewa','dalam pemeliharaan')"`
	Description string         `gorm:"type:text"`
	CreatedAt   time.Time      `gorm:"type:datetime(3)"`
	UpdatedAt   time.Time      `gorm:"type:datetime(3)"`
	DeletedAt   gorm.DeletedAt `gorm:"type:datetime(3)"`

	City     city.City         `gorm:"foreignKey:CityID"`
	Province province.Province `gorm:"foreignKey:ProvinceID"`
}
