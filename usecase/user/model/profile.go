package model

import "time"

type UpdateUsernameProfileRequest struct {
	Username string `json:"username" validate:"required"`
}

type GetProfileResponse struct {
	Id               int       `json:"id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	IsVerifyAccount  bool      `json:"isVerifyAccount"`
	IsVerifyIdentity bool      `json:"isVerifyIdentity"`
	Address          string    `json:"address"`
	Country          string    `json:"country"`
	Photo            string    `json:"photo"`
	NIK              string    `json:"nik"`
	FullName         string    `json:"fullName"`
	Gender           string    `json:"gender"`
	PlaceOfBirth     string    `json:"placeOfBirth"`
	DateBirth        time.Time `json:"dateBirth"`
	Work             string    `json:"work"`
	Citizenship      string    `json:"citizenship"`
	ProvinceID       string    `json:"provinceID"`
	ProvinceName     string    `json:"provinceName"`
	RegencyID        string    `json:"regencyID"`
	RegencyName      string    `json:"regencyName"`
	DistrictID       string    `json:"districtID"`
	DistrictName     string    `json:"districtName"`
}

type UploadPhotoResponse struct {
	UrlImage string `json:"urlImage"`
}

type UpdatePhotoProfileRequest struct {
	UrlImage string `json:"urlImage" validate:"url,required"`
}

type GetAvatarResponse struct {
	Id    int    `json:"id"`
	Image string `json:"image"`
}
