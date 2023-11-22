package model

type WarehouseDataRequest struct {
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	ProvinceID   int            `json:"provinceId"`
	RegencyID    int            `json:"regencyId"`
	DistrictID   int            `json:"DistrictId"`
	Address      string         `json:"address"`
	SurfaceArea  float64        `json:"surfaceArea"`
	BuildingArea float64        `json:"buildingArea"`
	Owner        string         `json:"owner"`
	PhoneNumber  string         `json:"phoneNumber"`
	Longitude    float64        `json:"longitude"`
	Latitude     float64        `json:"latitude"`
	Image        []WarehouseImg `json:"image"`
}

type WarehouseImg struct {
	Image string
}
