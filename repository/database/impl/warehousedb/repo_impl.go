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
