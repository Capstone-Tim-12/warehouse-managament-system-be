package warehouse

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/userdb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/warehousedb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/warehouse/model"
)

type defaultWarehouse struct {
	warehouseRepo warehousedb.WarehouseRepository
	userRepo      userdb.UserRepository
}

func NewWarehouseUsecase(warehouseRepo warehousedb.WarehouseRepository, userRepo userdb.UserRepository) WarehouseUsecase {
	return &defaultWarehouse{
		warehouseRepo: warehouseRepo,
		userRepo:      userRepo,
	}
}

func (s *defaultWarehouse) CreateWarehouse(ctx context.Context, req model.WarehouseDataRequest) (err error) {
	return
}
