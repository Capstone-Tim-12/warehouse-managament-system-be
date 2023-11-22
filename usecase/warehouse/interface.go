package warehouse

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/warehouse/model"
)

type WarehouseUsecase interface {
	CreateWarehouse(ctx context.Context, req model.WarehouseDataRequest, userId string) (err error)
	GetWarehouse(ctx context.Context, id string) (resp *model.WarehouseDataResponse, err error)
	GetAllWarehouse(ctx context.Context) (resp []*model.WarehouseIdResponse, err error)
	UpdateWarehouseDetails(ctx context.Context, req model.WarehouseDataRequest, userId string, id string) (err error)
}
