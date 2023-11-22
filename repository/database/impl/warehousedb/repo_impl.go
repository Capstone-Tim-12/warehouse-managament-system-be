package warehousedb

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"gorm.io/gorm"
)

type defaultRepo struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) WarehouseRepository {
	return &defaultRepo{db: db}
}

func (r *defaultRepo) CreateDetail(ctx context.Context, tx *gorm.DB, req *entity.Warehouse) (err error) {
	err = tx.WithContext(ctx).Create(req).Error
	return
}

func (r *defaultRepo) CreateImg(ctx context.Context, tx *gorm.DB, req *entity.WarehouseImg) (err error) {
	err = tx.WithContext(ctx).Create(req).Error
	return
}

func (r *defaultRepo) FindWarehouseById(ctx context.Context, id string) (resp *entity.Warehouse, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "id = ?", id).Error
	return
}

func (r *defaultRepo) FindAllWarehouse(ctx context.Context) (resp []entity.Warehouse, err error) {
	err = r.db.WithContext(ctx).Find(&resp).Error
	return
}

func (r *defaultRepo) BeginTrans(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Begin()
}
