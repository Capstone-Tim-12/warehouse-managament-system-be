package model

import "time"

type ListInstalmentResponse struct {
	InstalmentId int       `json:"instalmentId"`
	DueDate      time.Time `json:"dueDate"`
	Nominal      int       `json:"nominal"`
	Status       string    `json:"status"`
}
