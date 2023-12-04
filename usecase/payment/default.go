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
