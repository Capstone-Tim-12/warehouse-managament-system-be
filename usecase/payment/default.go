package payment

import (
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/paymentdb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/regiondb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/userdb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/warehousedb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/http/core"
)

type defaultPayment struct {
	regionRepo    regiondb.RegionRepository
	userRepo      userdb.UserRepository
	coreRepo      core.CoreWrapper
	warehouseRepo warehousedb.WarehouseRepository
	paymentRepo   paymentdb.PaymentRepository
}

func NewPaymentUsecase(regionRepo regiondb.RegionRepository,
	userRepo userdb.UserRepository,
	coreRepo core.CoreWrapper,
	warehouseRepo warehousedb.WarehouseRepository,
	paymentRepo paymentdb.PaymentRepository) PaymentUsecase {
	return &defaultPayment{
		regionRepo:    regionRepo,
		userRepo:      userRepo,
		coreRepo:      coreRepo,
		warehouseRepo: warehouseRepo,
		paymentRepo:   paymentRepo,
	}
}
