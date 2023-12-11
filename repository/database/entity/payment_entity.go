package entity

import (
	"time"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/generate"
	"gorm.io/gorm"
)

type TranscationStatus string
type InstalmentStatus string

const (
	Approved   TranscationStatus = "disetujui"
	Rejected   TranscationStatus = "ditolak"
	Submission TranscationStatus = "butuh persetujuan"
)

const (
	Paid    InstalmentStatus = "dibayar"
	Unpaid  InstalmentStatus = "belum dibayar"
	Waiting InstalmentStatus = "menunggu pembayaran"
	Failed  InstalmentStatus = "gagal bayar"
)

type PaymentScheme struct {
	ID          int `gorm:"primarykey"`
	Scheme      string
	Ration      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Transaction []Transaction
}

type Transaction struct {
	ID              string `gorm:"size:20;primarykey"`
	DateEntry       time.Time
	DateOut         time.Time
	UserID          int
	User            User `gorm:"foreignKey:UserID"`
	WarehouseID     int
	Warehouse       Warehouse `gorm:"foreignKey:WarehouseID"`
	PaymentSchemeID int
	PaymentScheme   PaymentScheme `gorm:"foreignKey:PaymentSchemeID"`
	Duration        int
	Status          TranscationStatus
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Instalment      []Instalment
}

type Instalment struct {
	ID                int `gorm:"primarykey"`
	TransactionID     string
	Transaction       Transaction `gorm:"foreignKey:TransactionID"`
	Nominal           float64
	DueDate           time.Time
	Status            InstalmentStatus
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	OngoingInstalment *OngoingInstalment
	TerminateContract *TerminateContract
}

type OngoingInstalment struct {
	ID              int `gorm:"primarykey"`
	InstalmentID    int
	Instalment      Instalment `gorm:"foreignKey:InstalmentID"`
	PaymentMethodID int
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID"`
	XPayment        string
	AccountNumber   string
	AdminFee        float64
	TotalPayment    float64
	BankCode        string
	PaymentTime     *time.Time
	Expired         time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type PaymentMethod struct {
	ID                int `gorm:"primarykey"`
	Name              string
	Image             string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	OngoingInstalment []OngoingInstalment
}

type TerminateContract struct {
	ID                        int `gorm:"primarykey"`
	InstalmentID              int
	Instalment                Instalment `gorm:"foreignKey:InstalmentID"`
	ReasonTerminateContractID int
	ReasonTerminateContract   ReasonTerminateContract `gorm:"foreignKey:ReasonTerminateContractID"`
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	DeletedAt                 gorm.DeletedAt `gorm:"index"`
}

type ReasonTerminateContract struct {
	ID        int `gorm:"primarykey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := generate.GenerateRandomString(20)
	m.ID = uuid

	return nil
}
