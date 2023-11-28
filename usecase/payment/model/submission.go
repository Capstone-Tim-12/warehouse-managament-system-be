package model

import "time"

type SubmissionRequest struct {
	WarehouseId     int       `json:"warehouseId"`
	PaymentSchemeId int       `json:"paymentSchemeId"`
	Duration        int       `json:"duration" validate:"required"`
	DateEntry       time.Time `json:"dateEntry" validate:"required"`
}