package model

import "time"

type TransactionHistoryResponse struct {
	TransactionID     string    `json:"transnactionId"`
	InstalmentId      int       `json:"instalmentId"`
	TransactionDate   time.Time `json:"transactionDate"`
	PaymentSchemeId   int       `json:"paymentSchemeId"`
	PaymentSchemeName string    `json:"paymentSchemeName"`
	UserID            int       `json:"userId"`
	UserName          string    `json:"username"`
	Nominal           float64   `json:"nominal"`
}
