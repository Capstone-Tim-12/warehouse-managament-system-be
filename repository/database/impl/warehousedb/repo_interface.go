package warehousedb

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"gorm.io/gorm"
)

type WarehouseRepository interface {
	CreateDetail(ctx context.Context, tx *gorm.DB, req *entity.Warehouse) (err error)
	CreateImg(ctx context.Context, tx *gorm.DB, req *entity.WarehouseImg) (err error)
	FindWarehouseById(ctx context.Context, id string) (resp *entity.Warehouse, err error)
	BeginTrans(ctx context.Context) *gorm.DB
}
