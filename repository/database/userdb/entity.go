package userdb

import "time"

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	ID               int `gorm:"primarykey"`
	Username         string
	Email            string
	IsVerifyAccount  bool
	IsVerifyIdentity bool
	Password         string
	Role             Role
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
	UserDetail       UserDetail
}

type UserDetail struct {
	ID           int `gorm:"primarykey"`
	Address      string
	Country      string
	Photo        string
	NIK          string
	FullName     string
	Gender       string
	PlaceOfBirth string
	DateBirth    time.Time
	Work         string
	Citizenship  string
	UserID       int
	ProvinceID   string
	RegencyID    string
	DistrictID   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
