package regiondb

type Province struct {
	ID      string
	Name    string
	Regency []Regency
}

type Regency struct {
	ID         string
	ProvinceID string
	Name       string
	Province   Province `gorm:"foreignKey:ProvinceID"`
	District   []District
}

type District struct {
	ID        string
	RegencyID string
	Name      string
	Regency   Regency `gorm:"foreignKey:RegencyID"`
	Village   []Village
}

type Village struct {
	ID         string
	DistrictID string
	Name       string
	District   District `gorm:"foreignKey:DistrictID"`
}
