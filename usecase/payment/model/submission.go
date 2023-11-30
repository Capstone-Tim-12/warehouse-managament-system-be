package model

import "time"

type SubmissionRequest struct {
	WarehouseId     int       `json:"warehouseId"`
	PaymentSchemeId int       `json:"paymentSchemeId"`
	Duration        int       `json:"duration" validate:"required"`
	DateEntry       time.Time `json:"dateEntry" validate:"required"`
}

type TransactionListDasboard struct {
	TransactionId int    `json:"transactionId"`
	WarehouseId   int    `json:"warehouseId"`
	UserId        int    `json:"userId"`
	Username      string `json:"userName"`
	ProvinceName  string `json:"provinceName"`
	RegencyName   string `json:"regencyName"`
	IsVerifyData  bool   `json:"isVerifyData"`
	Duration      string `json:"duration"`
	Status        string `json:"status"`
}
