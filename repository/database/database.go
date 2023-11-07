package database

import (
	"fmt"
	"os"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/capacity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/city"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/province"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/users"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/warehouse"
	warehousepicture "github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/warehousePicture"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDB() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}

func InitMigrate() {
	DB.AutoMigrate(
		users.User{},
		city.City{},
		capacity.Capacity{},
		province.Province{},
		warehouse.Warehouse{},
		warehousepicture.WarehousePicture{},
	)
}
