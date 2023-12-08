package model

import "time"

type TransactionDetailUser struct {
	WarehouseName    string           `json:"warehouseName"`
	Username         string           `json:"username"`
	IsVerifyIdentity bool             `json:"isVerifyIdentity"`
	Address          string           `json:"address"`
	DistricName      string           `json:"districName"`
	RegencyName      string           `json:"regencyName"`
	Duration         int              `json:"duration"`
	PaymentScheme    string           `json:"paymentScheme"`
	EntryDate        time.Time        `json:"entryDate"`
	OutDate          time.Time        `json:"outDate"`
	TotalPayment     float64          `json:"totalPayment"`
	Instalment       []InstalmentList `json:"instalment"`
}

type InstalmentList struct {
	InstalmentId int        `json:"instalmentId"`
	PaymentName  string     `json:"paymentName"`
	DueDate      time.Time  `json:"dueDate"`
	PaymentTime  *time.Time `json:"paymentTime"`
	Nominal      float64    `json:"nominal"`
	Status       string     `json:"status"`
}

const (
	PaymentActive = "Aktif"
	PaymentFinish = "Selesai"
)
