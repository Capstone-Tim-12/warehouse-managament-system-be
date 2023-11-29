package model

type TransactionHistoryResponse struct {
	TransactionID     int     `json:"transnactionId"`
	UserID            int     `json:"userId"`
	UserName          string  `json:"userName"`
	StatusTransaction string  `json:"statusTransaction"`
	Nominal           float64 `json:"nominal"`
}
