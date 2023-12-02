package model

type UserInfoResponse struct {
	UserId       int    `json:"userId"`
	Username     string `json:"username"`
	FullName     string `json:"fullName"`
	Email        string `json:"email"`
	Photo        string `json:"photo"`
	RegencyName  string `json:"regencyName"`
	DistrictName string `json:"districtName"`
}
