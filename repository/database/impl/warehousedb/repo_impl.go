package warehousedb

import (
	"context"
	"fmt"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
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
	err = r.db.WithContext(ctx).
		Preload("District").
		Preload("District.Regency").
		Preload("District.Regency.Province").
		Preload("WarehouseImg").
		Take(&resp, "id = ?", id).Error
	return
}

func (r *defaultRepo) FindImageWarehouseById(ctx context.Context, id string) (resp *entity.WarehouseImg, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "warehouse_id = ?", id).Error
	return
}

func (r *defaultRepo) FindWarehouseList(ctx context.Context, param paginate.Pagination, long, lat float64) (resp []entity.Warehouse, count int64, err error) {
	query := func(condision *gorm.DB) *gorm.DB {
		if param.Search != "" {
			condision.Where("name LIKE ?", "%"+param.Search+"%")
		}
		if param.MaxSize != 0 {
			condision.Where("building_area >= ? AND building_area <= ?", param.MinSize, param.MaxSize)
		}

		if param.MaxPrice != 0 {
			condision.Where("price >= ? AND price <= ?", param.MinPrice, param.HigestPrice)
		}

		switch true {
		case param.HigestPrice:
			condision.Order("price desc")
		case param.LowerPrice:
			condision.Order("price asc")
		case param.Recomendation:
			condision.Order(fmt.Sprintf("SQRT(POW(69.1 * (latitude - %v), 2) + POW(69.1 * (%v - longitude) * COS(latitude / 57.3), 2))", lat, long))
		}

		return condision
	}
	err = r.db.WithContext(ctx).Model(&entity.Warehouse{}).Scopes(query).Count(&count).Error
	if err != nil {
		return
	}
	err = r.db.WithContext(ctx).Preload("District").Preload("District.Regency").
		Preload("District.Regency.Province").Preload("WarehouseImg").
		Scopes(paginate.Paginate(param.Page, param.Limit)).Scopes(query).Find(&resp).Error
	return
}

func (r *defaultRepo) UpdateWarehouse(ctx context.Context, tx *gorm.DB, req *entity.Warehouse) (err error) {
	err = tx.WithContext(ctx).Save(req).Error
	return
}

func (r *defaultRepo) BeginTrans(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Begin()
}

func (s *defaultRepo) DeleteWarehouseImgByWarehouseId(ctx context.Context, tx *gorm.DB, warehouseId int) (err error) {
	err = tx.WithContext(ctx).Delete(&entity.WarehouseImg{}, "warehouse_id = ?", warehouseId).Error
	return
}

func (r *defaultRepo) GetWarehouseTypeById(ctx context.Context, id int) (resp *entity.WarehouseType, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "id = ?", id).Error
	return
}

func (r *defaultRepo) DeleteWarehouse(ctx context.Context, req *entity.Warehouse) (err error) {
	err = r.db.WithContext(ctx).Delete(&req).Error
	return
}
