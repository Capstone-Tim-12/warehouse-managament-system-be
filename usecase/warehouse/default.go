package warehouse

import (
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/regiondb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/userdb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/warehousedb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/http/core"
)

type defaultWarehouse struct {
	warehouseRepo warehousedb.WarehouseRepository
	userRepo      userdb.UserRepository
	regionRepo    regiondb.RegionRepository
	coreWrapper   core.CoreWrapper
}

func NewWarehouseUsecase(warehouseRepo warehousedb.WarehouseRepository, userRepo userdb.UserRepository, regionRepo regiondb.RegionRepository, coreWrapper core.CoreWrapper) WarehouseUsecase {
	return &defaultWarehouse{
		warehouseRepo: warehouseRepo,
		userRepo:      userRepo,
		regionRepo:    regionRepo,
		coreWrapper:   coreWrapper,
	}
}
