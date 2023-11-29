package warehousedb

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
	"gorm.io/gorm"
)

type WarehouseRepository interface {
	CreateDetail(ctx context.Context, tx *gorm.DB, req *entity.Warehouse) (err error)
	CreateImg(ctx context.Context, tx *gorm.DB, req *entity.WarehouseImg) (err error)
	FindWarehouseById(ctx context.Context, id string) (resp *entity.Warehouse, err error)
	FindImageWarehouseById(ctx context.Context, id string) (resp *entity.WarehouseImg, err error)
	FindWarehouseList(ctx context.Context, param paginate.Pagination, long, lat float64) (resp []entity.Warehouse, count int64, err error)
	UpdateWarehouse(ctx context.Context, tx *gorm.DB, req *entity.Warehouse) (err error)
	BeginTrans(ctx context.Context) *gorm.DB
	DeleteWarehouseImgByWarehouseId(ctx context.Context, tx *gorm.DB, warehouseId int) (err error)
	GetWarehouseTypeById(ctx context.Context, id int) (resp *entity.WarehouseType, err error)
}
