package payment

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/paymentdb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/regiondb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/userdb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/warehousedb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/http/core"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/generate"
)

type defaultPayment struct {
	regionRepo    regiondb.RegionRepository
	userRepo      userdb.UserRepository
	coreWrapper   core.CoreWrapper
	warehouseRepo warehousedb.WarehouseRepository
	paymentRepo   paymentdb.PaymentRepository
}

func NewPaymentUsecase(regionRepo regiondb.RegionRepository,
	userRepo userdb.UserRepository,
	coreWrapper core.CoreWrapper,
	warehouseRepo warehousedb.WarehouseRepository,
	paymentRepo paymentdb.PaymentRepository) PaymentUsecase {
	return &defaultPayment{
		regionRepo:    regionRepo,
		userRepo:      userRepo,
		coreWrapper:   coreWrapper,
		warehouseRepo: warehouseRepo,
		paymentRepo:   paymentRepo,
	}
}

func (s *defaultPayment) generateBodyEmailPaymentNotif(ctx context.Context, req model.NotifPayment) (resp string, err error) {
	local, _ := time.LoadLocation("Asia/Jakarta")
	configPath := "./template/payment_notif.html"
	tempByte, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("error reading template file: ", err.Error())
		return
	}
	emailBody := string(tempByte)
	emailBody = strings.ReplaceAll(emailBody, "[Nama Pelanggan]", req.Username)
	emailBody = strings.ReplaceAll(emailBody, "[Nomor Pesanan]", req.Xpayment)
	emailBody = strings.ReplaceAll(emailBody, "[Payment Method]", req.PaymentMethod)
	emailBody = strings.ReplaceAll(emailBody, "[Va Name]", req.VaName)
	emailBody = strings.ReplaceAll(emailBody, "[Virtual Account]", req.VaNumber)
	emailBody = strings.ReplaceAll(emailBody, "[Bank Code]", req.BankCode)
	emailBody = strings.ReplaceAll(emailBody, "[Nominal]", "Rp " + generate.FormatRupiah(req.Nominal))
	emailBody = strings.ReplaceAll(emailBody, "[Batas Pembayaran]", req.Expired.In(local).Format("02 January 2006 15:04"))

	resp = emailBody
	return
}

func (s *defaultPayment) generatePaymentSuccessBody(req core.CheckPaymentData, transactionDate time.Time) (emailBody string, err error) {
	local, _ := time.LoadLocation("Asia/Jakarta")
	configPath := "./template/payment_success.html"
	tempByte, err := os.ReadFile(configPath)
	if err != nil {
		return
	}
	emailBody = string(tempByte)
	emailBody = strings.ReplaceAll(emailBody, "[Nomor Pesanan]", req.ExternalID)
	emailBody = strings.ReplaceAll(emailBody, "[Tanggal Transaksi]", transactionDate.In(local).Format("02 January 2006 15:04"))
	emailBody = strings.ReplaceAll(emailBody, "[Total Pembayaran]", "Rp " + generate.FormatRupiah(int(req.Amount)))
	emailBody = strings.ReplaceAll(emailBody, "[Tanggal Pembayaran]", req.TransactionTimestamp.In(local).Format("02 January 2006 15:04"))
	emailBody = strings.ReplaceAll(emailBody, "[Nomor Virtual Account]", req.AccountNumber)
	emailBody = strings.ReplaceAll(emailBody, "[Nama Bank]", req.BankCode)
	emailBody = strings.ReplaceAll(emailBody, "[Tanggal Notifikasi]", time.Now().In(local).Format("02 January 2006"))

	return
}

