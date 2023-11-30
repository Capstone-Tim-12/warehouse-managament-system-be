package model

type ListTrxUser struct {
	TransactionId string
}

type ListAllTrxResponse struct {
	UserId        int    `json:"userId"`
	Username      string `json:"username"`
	RegencyId     string `json:"regencyId"`
	RegencyName   string `json:"regencyName"`
	ProvinceId    string `json:"provinceId"`
	ProvinceName  string `json:"provinceName"`
	WarehouseName string `json:"warehouseName"`
	WarehouseId   int    `json:"warehouseId"`
	Duration      int    `json:"duration"`
	PaymentScheme string `json:"paymentScheme"`
	Status        string `json:"status"`
}
