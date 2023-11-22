package warehousedb

import (
	"context"

	"gorm.io/gorm"
)

type defaultRepo struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) WarehouseRepository {
	return &defaultRepo{db: db}
}

func (r *defaultRepo) CreateDetail(ctx context.Context, tx *gorm.DB, req *UserDetail) (err error) {
	err = tx.WithContext(ctx).Create(req).Error
	return
}
