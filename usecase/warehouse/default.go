package warehouse

import (
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/regiondb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/userdb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/warehousedb"
)

type defaultWarehouse struct {
	warehouseRepo warehousedb.WarehouseRepository
	userRepo      userdb.UserRepository
	regionRepo    regiondb.RegionRepository
}

func NewWarehouseUsecase(warehouseRepo warehousedb.WarehouseRepository, userRepo userdb.UserRepository, regionRepo regiondb.RegionRepository) WarehouseUsecase {
	return &defaultWarehouse{
		warehouseRepo: warehouseRepo,
		userRepo:      userRepo,
		regionRepo:    regionRepo,
	}
}
