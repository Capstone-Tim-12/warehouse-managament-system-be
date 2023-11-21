package entity

type Province struct {
	ID         string `gorm:"size:12;primarykey"`
	Name       string
	Regency    []Regency
	UserDetail []UserDetail
	Warehouse  []Warehouse
}

type Regency struct {
	ID         string `gorm:"size:12;primarykey"`
	ProvinceID string
	Name       string
	District   []District
	UserDetail []UserDetail
	Warehouse  []Warehouse
}

type District struct {
	ID         string `gorm:"size:12;primarykey"`
	RegencyID  string
	Name       string
	Village    []Village
	UserDetail []UserDetail
	Warehouse  []Warehouse
}

type Village struct {
	ID         string `gorm:"size:12;primarykey"`
	DistrictID string
	Name       string
}
