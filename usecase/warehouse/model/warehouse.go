package model

type WarehouseDataRequest struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	ProvinceID   string   `json:"provinceId"`
	RegencyID    string   `json:"regencyId"`
	DistrictID   string   `json:"DistrictId"`
	Address      string   `json:"address"`
	SurfaceArea  float64  `json:"surfaceArea"`
	BuildingArea float64  `json:"buildingArea"`
	Owner        string   `json:"owner"`
	PhoneNumber  string   `json:"phoneNumber"`
	Longitude    float64  `json:"longitude"`
	Latitude     float64  `json:"latitude"`
	Image        []string `json:"image"`
}

type WarehouseDataResponse struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	ProvinceID   string   `json:"provinceId"`
	ProvinceName string   `json:"provinceName"`
	RegencyID    string   `json:"regencyId"`
	RegencyName  string   `json:"regencyName"`
	DistrictID   string   `json:"districtId"`
	DistrictName string   `json:"districtName"`
	Address      string   `json:"address"`
	SurfaceArea  float64  `json:"surfaceArea"`
	BuildingArea float64  `json:"buildingArea"`
	Owner        string   `json:"owner"`
	PhoneNumber  string   `json:"phoneNumber"`
	Longitude    float64  `json:"longitude"`
	Latitude     float64  `json:"latitude"`
	Status       string   `json:"status"`
	Image        []string `json:"image"`
}
