package entity

import (
	"time"

	"gorm.io/gorm"
)

type TranscationStatus string
type InstalmentStatus string

const (
	Approved   = "disetujui"
	Rejected   = "ditolak"
	Submission = "pengajuan"
)

const (
	Paid   InstalmentStatus = "dibayar"
	Unpaid InstalmentStatus = "belum dibayar"
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
	ID              int `gorm:"primarykey"`
	DateEntry       time.Time
	DateOut         time.Time
	UserID          int
	User            User `gorm:"foreignKey:UserID"`
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
	TransactionID     int
	Transaction       Transaction `gorm:"foreignKey:TransactionID"`
	Nominal           float64
	DueDate           time.Time
	Status            InstalmentStatus
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	OngoingInstalment *OngoingInstalment
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
	PaymentTime     time.Time
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
