package regiondb

import (
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/userdb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/warehousedb"
	// "github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/warehousedb"
)

type Province struct {
	ID         string `gorm:"size:12;primarykey"`
	Name       string
	Regency    []Regency
	UserDetail userdb.UserDetail
	Warehouse  warehousedb.Warehouse
}

type Regency struct {
	ID         string `gorm:"size:12;primarykey"`
	ProvinceID string
	Name       string
	District   []District
	UserDetail userdb.UserDetail
	Warehouse  warehousedb.Warehouse
}

type District struct {
	ID         string `gorm:"size:12;primarykey"`
	RegencyID  string
	Name       string
	Village    []Village
	UserDetail userdb.UserDetail
	Warehouse warehousedb.Warehouse
}

type Village struct {
	ID         string `gorm:"size:12;primarykey"`
	DistrictID string
	Name       string
}
