package warehouse

import (
	"context"
	"mime/multipart"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/warehouse/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
)

type WarehouseUsecase interface {
	CreateWarehouse(ctx context.Context, req model.WarehouseDataRequest) (err error)
	GetWarehouse(ctx context.Context, id int) (resp *model.WarehouseDataResponse, err error)
	GetWarehouseList(ctx context.Context, param paginate.Pagination, userId int) (resp []model.WarehouseListResponse, count int64, err error)
	UpdateWarehouseDetails(ctx context.Context, req model.WarehouseDataRequest, id int) (err error)
	DeleteWarehouse(ctx context.Context, id string) (err error)
	GetListWarehouseType(ctx context.Context) (resp []model.WarehouseTypeResponse, err error)
	UploadPhotoWarehouse(ctx context.Context, photo []*multipart.FileHeader) (resp model.UploadPhotoResponse, err error)
	GetMywarehouse(ctx context.Context, userId int, status model.TrxStatus, param paginate.Pagination) (resp []model.MyWarehoyseResponse, count int64, err error)
	GetWarehouseInfo(ctx context.Context, warehouseId int) (resp model.WarehouseInfoResponse, err error)
	AddFavorite(ctx context.Context, userId int, req model.AddFavoritRequest) (err error)
	DeleteFavorit(ctx context.Context, userId, warehouseId int) (err error)
	GetListFavorite(ctx context.Context, userId int, param paginate.Pagination) (resp []model.WarehouseListResponse, count int64, err error)
	ImportCsvFileWarehouse(ctx context.Context, file *multipart.FileHeader) (err error)
}
