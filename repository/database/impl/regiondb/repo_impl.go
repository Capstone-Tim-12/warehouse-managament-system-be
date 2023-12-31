package regiondb

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"gorm.io/gorm"
)

type defaultRepo struct {
	db *gorm.DB
}

func NewRegionRepository(db *gorm.DB) RegionRepository {
	return &defaultRepo{db: db}
}

func (r *defaultRepo) FindAllProvince(ctx context.Context) (resp []entity.Province, err error){
	err = r.db.WithContext(ctx).Find(&resp).Error
	return
}

func (r *defaultRepo) FindRegencyByProvinceId(ctx context.Context, id string)(resp []entity.Regency, err error) {
	err = r.db.WithContext(ctx).Find(&resp, "province_id = ?", id).Error
	return
}

func (r *defaultRepo) FindDistrictByRegencyId(ctx context.Context, id string) (resp []entity.District, err error) {
	err = r.db.WithContext(ctx).Find(&resp, "regency_id = ?", id).Error
	return
}

func (r *defaultRepo) FindVillageByDistrictId(ctx context.Context, id string) (resp []entity.Village, err error) {
	err = r.db.WithContext(ctx).Find(&resp, "district_id = ?", id).Error
	return
}

func (r *defaultRepo) GetProvinceById(ctx context.Context, id string) (resp *entity.Province, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "id = ?", id).Error
	return
}

func (r *defaultRepo) GetRegencyById(ctx context.Context, id string) (resp *entity.Regency, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "id = ?", id).Error
	return
}

func (r *defaultRepo) GetDistrictById(ctx context.Context, id string) (resp *entity.District, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "id = ?", id).Error
	return
}

func (r *defaultRepo) GetVillageById(ctx context.Context, id string) (resp *entity.Village, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "id = ?", id).Error
	return
}