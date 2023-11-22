package warehouse

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/warehouse/model"
)

type WarehouseUsecase interface {
	CreateWarehouse(ctx context.Context, req model.WarehouseDataRequest, userId string) (err error)
	GetWarehouse(ctx context.Context, id string) (resp *model.WarehouseDataResponse, err error)
}
