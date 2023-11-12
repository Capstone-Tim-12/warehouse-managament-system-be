package transactiondb

import (
	"time"
)

type Rent struct {
	ID                     int `gorm:"primarykey"`
	WarehouseID            int
	UserID                 int
	WarehousePaymentTypeID int
	EntryDate              time.Time
	OutDate                time.Time
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAt              time.Time
}
