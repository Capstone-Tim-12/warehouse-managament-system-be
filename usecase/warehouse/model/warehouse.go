package model

import "github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"

type WarehouseDataRequest struct {
	Name            string   `json:"name" validate:"required"`
	Description     string   `json:"description" validate:"required"`
	DistrictID      string   `json:"districId" validate:"number"`
	WarehouseTypeId int      `json:"warehouseTypeId" validate:"number"`
	Address         string   `json:"address" validate:"required"`
	SurfaceArea     float64  `json:"surfaceArea" validate:"required"`
	BuildingArea    float64  `json:"buildingArea" validate:"required"`
	Owner           string   `json:"owner" validate:"required"`
	PhoneNumber     string   `json:"phoneNumber" validate:"number"`
	Longitude       float64  `json:"longitude" validate:"longitude"`
	Latitude        float64  `json:"latitude" validate:"latitude"`
	Status          string   `json:"status"`
	Price           float64  `json:"price" validate:"required"`
	Image           []string `json:"image"`
}

type WarehouseDataResponse struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	ProvinceID    string   `json:"provinceId"`
	ProvinceName  string   `json:"provinceName"`
	RegencyID     string   `json:"regencyId"`
	RegencyName   string   `json:"regencyName"`
	DistrictID    string   `json:"districtId"`
	DistrictName  string   `json:"districtName"`
	Address       string   `json:"address"`
	SurfaceArea   float64  `json:"surfaceArea"`
	BuildingArea  float64  `json:"buildingArea"`
	Owner         string   `json:"owner"`
	PhoneNumber   string   `json:"phoneNumber"`
	Longitude     float64  `json:"longitude"`
	Latitude      float64  `json:"latitude"`
	Status        string   `json:"status"`
	WeeklyPrice   float64  `json:"weeklyPrice"`
	MonthlyPrice  float64  `json:"monthlyPrice"`
	AnnualPrice   float64  `json:"annualPrice"`
	WarehouseType string   `json:"warehouseType"`
	Image         []string `json:"image"`
}

type WarehouseListResponse struct {
	Id                int     `json:"id"`
	Name              string  `json:"name"`
	DistrictName      string  `json:"districtName"`
	RegencyName       string  `json:"regencyName"`
	ProvinceName      string  `json:"provinceName"`
	SurfaceArea       float64 `json:"surfaceArea"`
	BuildingArea      float64 `json:"buildingArea"`
	WeeklyPrice       float64 `json:"weeklyPrice"`
	MonthlyPrice      float64 `json:"monthlyPrice"`
	AnnualPrice       float64 `json:"annualPrice"`
	Distance          float64 `json:"distance"`
	Status            string  `json:"status"`
	WarehouseTypeId   int     `json:"warehouseTypeId"`
	WarehouseTypeName string  `json:"warehouseTypeName"`
	Image             string  `json:"image"`
}

type WarehouseInfoResponse struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	DistrictName string  `json:"districtName"`
	RegencyName  string  `json:"regencyName"`
	ProvinceName string  `json:"provinceName"`
	AnnualPrice  float64 `json:"annualPrice"`
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

type WarehouseTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type UploadPhotoResponse struct {
	Images []string `json:"images"`
}

type MyWarehoyseResponse struct {
	TransactionId     string  `json:"transactionId"`
	TransactionStatus string  `json:"transactionStatus"`
	WarehouseId       int     `json:"warehouseId"`
	WarehouseName     string  `json:"warehouseName"`
	WatehouseRegency  string  `json:"warehouseRegency"`
	SurfaceArea       float64 `json:"surfaceArea"`
	BuildingArea      float64 `json:"buildingArea"`
	WarehouseImage    string  `json:"warehouseImage"`
}

type TrxStatus string

const (
	StatusSubmitted TrxStatus = "diajukan"
	StatusRented    TrxStatus = "disewa"
)

var GetStatusTrx = map[TrxStatus]entity.TranscationStatus{
	StatusSubmitted: entity.Submission,
	StatusRented:    entity.Approved,
}
