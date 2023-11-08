package userdb

import "time"

type User struct {
	ID               int `gorm:"primarykey"`
	Username         string
	Email            string
	IsVerifyAccount  bool
	IsVerifyIdentity bool
	Password         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}

type UserDetail struct {
	ID           int `gorm:"primarykey"`
	Photo        string
	NIK          string
	FullName     string
	Gender       string
	PlaceOfBirth time.Time
	Works        string
	Citizenship  string
	UserID       int
	ProvinceID   int
	RegencyID    int
	DistrictID   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
