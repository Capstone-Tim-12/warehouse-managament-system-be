package model

import "time"

type NotifPayment struct {
	Username      string    `json:"username"`
	Xpayment      string    `json:"xPayment"`
	VaNumber      string    `json:"VaNumber"`
	PaymentMethod string    `json:"paymentMethod"`
	VaName        string    `json:"vaName"`
	BankCode      string    `json:"bankCode"`
	Nominal       int       `json:"nominal"`
	Expired       time.Time `json:"expired"`
}

