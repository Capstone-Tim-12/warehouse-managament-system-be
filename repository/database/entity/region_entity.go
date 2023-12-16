package entity

type Province struct {
	ID         string `gorm:"size:12;primarykey"`
	Name       string
	Regency    []Regency
}

type Regency struct {
	ID         string `gorm:"size:12;primarykey"`
	ProvinceID string
	Province   Province `gorm:"foreignKey:ProvinceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name       string
	District   []District
}

type District struct {
	ID         string `gorm:"size:12;primarykey"`
	RegencyID  string
	Regency    Regency `gorm:"foreignKey:RegencyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name       string
	Village    []Village
	UserDetail []UserDetail
	Warehouse  []Warehouse
}

type Village struct {
	ID         string   `gorm:"size:12;primarykey"`
	District   District `gorm:"foreignKey:DistrictID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DistrictID string
	Name       string
}
