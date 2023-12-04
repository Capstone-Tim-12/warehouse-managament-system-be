package model

import "time"

type PaymentRequest struct {
	PaymentMethodId int    `json:"paymentMethodId"`
	InstalmentId    int    `json:"instalmentId"`
	Data            string `json:"data"`
}

type PaymentResponse struct {
	PaymentInfo string `json:"paymentInfo"`
}

type VaDataRequest struct {
	BankCode string `json:"bankCode"`
}

type VaDataResponse struct {
	XpaymentId           string    `json:"xPaymentId"`
	VirtualAccountName   string    `json:"virtualAccountName"`
	VirtualAccountNumber string    `json:"virtualAccountNumber"`
	BankCode             string    `json:"bankCode"`
	Nominal              int       `json:"nominal"`
	ExpiredAt            time.Time `json:"expiredAt"`
}

type VaCallbackRequest struct {
	AccountNumber            string        `json:"account_number"`
	Amount                   int64         `json:"amount"`
	BankCode                 string        `json:"bank_code"`
	CallbackVirtualAccountID string        `json:"callback_virtual_account_id"`
	Currency                 string        `json:"currency"`
	ExternalID               string        `json:"external_id"`
	ID                       string        `json:"id"`
	MerchantCode             string        `json:"merchant_code"`
	OwnerID                  string        `json:"owner_id"`
	PaymentDetail            PaymentDetail `json:"payment_detail"`
	PaymentID                string        `json:"payment_id"`
	SenderName               string        `json:"sender_name"`
	TransactionTimestamp     *time.Time    `json:"transaction_timestamp"`
}

type PaymentDetail struct {
	Reference string `json:"reference"`
	Remark    string `json:"remark"`
}

