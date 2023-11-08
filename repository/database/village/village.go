package village

type Village struct {
	ID         uint `gorm:"primaryKey"`
	DistrictID uint
	Name       string `gorm:"type:varchar(255)"`
}
