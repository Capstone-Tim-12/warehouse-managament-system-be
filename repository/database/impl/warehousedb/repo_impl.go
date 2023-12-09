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
		Preload("District.Regency.Province").
		Preload("WarehouseImg").
		Preload("WarehouseType").
		Take(&resp, "id = ?", id).Error
	return
}

func (r *defaultRepo) FindWarehouseByIdOnly(ctx context.Context, id string) (resp *entity.Warehouse, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "id = ?", id).Error
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

		if param.Status != "" {
			condision.Where("status = ?", param.Status)
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
	err = r.db.WithContext(ctx).
		Preload("District.Regency.Province").
		Preload("WarehouseImg").
		Preload("WarehouseType").
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

func (r *defaultRepo) GetListWarehouseType(ctx context.Context) (resp []entity.WarehouseType, err error) {
	err = r.db.WithContext(ctx).Find(&resp).Error
	return
}

func (s *defaultRepo) AddFavorit(ctx context.Context, req *entity.Favorit) (err error) {
	err = s.db.WithContext(ctx).Create(req).Error
	return
}

func (s *defaultRepo) FindFavoritById(ctx context.Context, waehouseId int) (resp *entity.Favorit, err error) {
	err = s.db.WithContext(ctx).Take(&resp, "id = ?", waehouseId).Error
	return
}

func (s *defaultRepo) FindFavoritByWarehouseIdAndUserId(ctx context.Context, waehouseId, userId int) (resp *entity.Favorit, err error) {
	err = s.db.WithContext(ctx).Take(&resp, "warehouse_id = ? AND user_id = ?", waehouseId, userId).Error
	return
}

func (s *defaultRepo) DeleteFavorite(ctx context.Context, userId, warehouseId int) (err error) {
	err = s.db.WithContext(ctx).Delete(&entity.Favorit{}, "user_id = ? AND warehouse_id = ?", userId, warehouseId).Error
	return
}

func (s *defaultRepo) FindListFavoriteByUserId(ctx context.Context, userId int, param paginate.Pagination) (resp []entity.Favorit, count int64, err error) {
	query := func (db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userId)
	}

	err = s.db.WithContext(ctx).Model(&entity.Favorit{}).Scopes(query).Count(&count).Error
	if err != nil {
		return
	}

	err = s.db.WithContext(ctx).Scopes(paginate.Paginate(param.Page, param.Limit)).
			Preload("Warehouse.District.Regency.Province").
			Preload("Warehouse.WarehouseType").
			Preload("Warehouse.WarehouseImg").
			Preload("User").Scopes(query).Find(&resp).Error
	return
}

func (s *defaultRepo) GetTotalWarehouseByStatus(ctx context.Context, status entity.WarehouseStatus) (total int64, err error) {
	err = s.db.WithContext(ctx).Model(&entity.Warehouse{}).Where("status = ?", status).Count(&total).Error
	return
}
