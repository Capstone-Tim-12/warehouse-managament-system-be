package database

import (
	"fmt"
	"os"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/regiondb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/userdb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/warehousedb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if os.Getenv("DB_DEBUG") == "true" {
		DB = DB.Debug()
	}

	if os.Getenv("DB_MIGRATION") == "true" {
		DB.AutoMigrate(&regiondb.Province{})
		DB.AutoMigrate(&regiondb.Regency{})
		DB.AutoMigrate(&regiondb.District{})
		DB.AutoMigrate(&regiondb.Village{})
		DB.AutoMigrate(&userdb.User{})
		DB.AutoMigrate(&userdb.UserDetail{})
		DB.AutoMigrate(&warehousedb.PaymentScheme{})
		DB.AutoMigrate(&warehousedb.Warehouse{})
		DB.AutoMigrate(&warehousedb.WarehouseImg{})
	}

	return DB
}
