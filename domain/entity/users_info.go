package entity

import "time"

type UserInfo struct {
	NIK         string    `json:"nik"`
	FullName    string    `json:"fullName"`
	Gender      string    `json:"gender"`
	PlaceBirth  string    `json:"placeBirth"`
	DateBirth   time.Time `json:"dateBirth"`
	Work        string    `json:"work"`
	Citizenship string    `json:"citizenship"`
}
