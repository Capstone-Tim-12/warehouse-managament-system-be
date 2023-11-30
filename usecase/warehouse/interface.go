package warehouse

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/warehouse/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
)

type WarehouseUsecase interface {
	CreateWarehouse(ctx context.Context, req model.WarehouseDataRequest) (err error)
	GetWarehouse(ctx context.Context, id string) (resp *model.WarehouseDataResponse, err error)
	GetWarehouseList(ctx context.Context, param paginate.Pagination, userId int) (resp []model.WarehouseListResponse, count int64, err error)
	UpdateWarehouseDetails(ctx context.Context, req model.WarehouseDataRequest, id string) (err error)
	DeleteWarehouse(ctx context.Context, id string) (err error)
}
