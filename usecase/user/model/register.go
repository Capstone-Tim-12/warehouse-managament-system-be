package model

import "time"

type RegisterDataRequest struct {
	NIK         string    `json:"nik"`
	Address     string    `json:"address"`
	Email       string    `json:"email"`
	FullName    string    `json:"fullName"`
	Gender      string    `json:"gender"`
	PlaceBirth  string    `json:"placeBirth"`
	DateBirth   time.Time `json:"dateBirth"`
	Work        string    `json:"work"`
	Citizenship string    `json:"citizenship"`
	ProvinceID  string    `json:"provinceId"`
	RegencyID   string    `json:"regionId"`
	DistrictID  string    `json:"districtId"`
}

type RegisterUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	Email string `json:"email"`
}
