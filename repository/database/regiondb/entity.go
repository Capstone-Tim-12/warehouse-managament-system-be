package regiondb

import "github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/userdb"

type Province struct {
	ID         string `gorm:"size:12;primarykey"`
	Name       string
	Regency    []Regency
	UserDetail []userdb.UserDetail
}

type Regency struct {
	ID         string `gorm:"size:12;primarykey"`
	ProvinceID string
	Name       string
	District   []District
	UserDetail []userdb.UserDetail
}

type District struct {
	ID         string `gorm:"size:12;primarykey"`
	RegencyID  string
	Name       string
	Village    []Village
	UserDetail []userdb.UserDetail
}

type Village struct {
	ID         string `gorm:"size:12;primarykey"`
	DistrictID string
	Name       string
}
