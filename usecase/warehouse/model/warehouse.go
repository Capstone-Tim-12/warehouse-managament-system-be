package model

type WarehouseDataRequest struct {
	Name         string   `json:"name" validate:"required"`
	Description  string   `json:"description" validate:"required"`
	DistrictID   string   `json:"districId" validate:"number"`
	Address      string   `json:"address" validate:"required"`
	SurfaceArea  float64  `json:"surfaceArea" validate:"required"`
	BuildingArea float64  `json:"buildingArea" validate:"required"`
	Owner        string   `json:"owner" validate:"required"`
	PhoneNumber  string   `json:"phoneNumber" validate:"number"`
	Longitude    float64  `json:"longitude" validate:"longitude"`
	Latitude     float64  `json:"latitude" validate:"latitude"`
	Status       string   `json:"status"`
	Price        float64  `json:"price"`
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
	WeeklyPrice  float64  `json:"weeklyPrice"`
	MonthlyPrice float64  `json:"monthlyPrice"`
	AnnualPrice  float64  `json:"annualPrice"`
	Image        []string `json:"image"`
}

type WarehouseListResponse struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	DistrictName string  `json:"DistrictName"`
	RegencyName  string  `json:"RegencyName"`
	ProvinceName string  `json:"ProvinceName"`
	SurfaceArea  float64 `json:"surfaceArea"`
	BuildingArea float64 `json:"buildingArea"`
	WeeklyPrice  float64 `json:"weeklyPrice"`
	MonthlyPrice float64 `json:"monthlyPrice"`
	AnnualPrice  float64 `json:"annualPrice"`
	Distance     float64 `json:"distance"`
	Image        string  `json:"image"`
}

type GetWarehouseAppResponse struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	RegencyName  string  `json:"regencyName"`
	WeeklyPrice  float64 `json:"weeklyPrice"`
	MonthlyPrice float64 `json:"monthlyPrice"`
	AnnualPrice  float64 `json:"annualPrice"`
	Image        string  `json:"image"`
}
