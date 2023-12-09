package model

type GetTotalPayment struct {
	TotalWarehouseAvailabe     int64   `json:"totalWarehouseAvailabe"`
	TotalWarehouseNotAvailable int64   `json:"totalWarehouseNotAvailable"`
	TotalPayment               float64 `json:"totalPayment"`
}

type StatiscticPayment struct {
	Year         int     `json:"year"`
	TotalPayment float64 `json:"totalPayment"`
}
