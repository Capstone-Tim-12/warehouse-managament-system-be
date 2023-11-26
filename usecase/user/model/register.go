package model

import "time"

type RegisterDataRequest struct {
	NIK         string    `json:"nik" validate:"required,min=16"`
	Address     string    `json:"address" validate:"required"`
	Email       string    `json:"email" validate:"email"`
	FullName    string    `json:"fullName" validate:"required"`
	Gender      string    `json:"gender" validate:"oneof=M F"`
	PlaceBirth  string    `json:"placeBirth" validate:"required"`
	DateBirth   time.Time `json:"dateBirth" validate:"required"`
	Work        string    `json:"work" validate:"required"`
	Citizenship string    `json:"citizenship" validate:"required"`
	DistrictID  string    `json:"districtId" validate:"number"`
}

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"min=6"`
}

type RegisterUserResponse struct {
	Email string `json:"email"`
}
