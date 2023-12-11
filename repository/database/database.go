package database

import (
	"fmt"
	"os"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
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
		DB.AutoMigrate(&entity.Province{})
		DB.AutoMigrate(&entity.Regency{})
		DB.AutoMigrate(&entity.District{})
		DB.AutoMigrate(&entity.Village{})
		DB.AutoMigrate(&entity.User{})
		DB.AutoMigrate(&entity.UserDetail{})
		DB.AutoMigrate(&entity.Avatar{})
		DB.AutoMigrate(&entity.WarehouseType{})
		DB.AutoMigrate(&entity.Warehouse{})
		DB.AutoMigrate(&entity.WarehouseImg{})
		DB.AutoMigrate(&entity.Favorit{})
		DB.AutoMigrate(&entity.PaymentScheme{})
		DB.AutoMigrate(&entity.Transaction{})
		DB.AutoMigrate(&entity.Instalment{})
		DB.AutoMigrate(&entity.PaymentMethod{})
		DB.AutoMigrate(&entity.OngoingInstalment{})
		DB.AutoMigrate(&entity.ReasonTerminateContract{})
		DB.AutoMigrate(&entity.TerminateContract{})
	}

	return DB
}
