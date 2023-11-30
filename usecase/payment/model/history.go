package model

type TransactionHistoryResponse struct {
	TransactionID     string  `json:"transnactionId"`
	InstalmentId      int     `json:"instalmentId"`
	UserID            int     `json:"userId"`
	UserName          string  `json:"userName"`
	StatusTransaction string  `json:"statusTransaction"`
	Nominal           float64 `json:"nominal"`
}
