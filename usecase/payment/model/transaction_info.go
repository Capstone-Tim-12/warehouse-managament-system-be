package model

import "time"

type TransactionInfoResponse struct {
	DateEntry     time.Time `json:"dateEntry"`
	DateOut       time.Time `json:"dateOut"`
	PaymentScheme string    `json:"paymentScheme"`
	Duration      int       `json:"duration"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
}
