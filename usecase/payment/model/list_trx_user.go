package model

type ListTrxUser struct {
	TransactionId string
}

type ListAllTrxResponse struct {
	TransactionId string `json:"transactionId"`
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

type TrxListDetail struct {
	WarehouseId       int     `json:"warehouseId"`
	WarehouseName     string  `json:"warehouseName"`
	WarehousePrice    float64 `json:"warehousePrice"`
	WarehouseAdreess  string  `json:"warehouseAdreess"`
	WarehouseDistrict string  `json:"warehouseDistrict"`
	WarehouseRegency  string  `json:"warehouseRegency"`
	WarehouseProvince string  `json:"warehouseProvince"`
	WarehouseImage    string  `json:"warehouseImage"`
	Username          string  `json:"username"`
	IsVerifyIdentity  bool    `json:"isVerifyIdentity"`
	RentalDuration    int     `json:"rentalDuration"`
	PaymentScheme     string  `json:"paymentScheme"`
}
