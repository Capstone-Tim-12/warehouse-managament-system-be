package regiondb

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
)

type RegionRepository interface {
	FindAllProvince(ctx context.Context) (resp []entity.Province, err error)
	FindRegencyByProvinceId(ctx context.Context, id string)(resp []entity.Regency, err error)
	FindDistrictByRegencyId(ctx context.Context, id string) (resp []entity.District, err error)
	FindVillageByDistrictId(ctx context.Context, id string) (resp []entity.Village, err error)
	GetProvinceById(ctx context.Context, id string) (resp *entity.Province, err error)
	GetRegencyById(ctx context.Context, id string) (resp *entity.Regency, err error)
	GetDistrictById(ctx context.Context, id string) (resp *entity.District, err error) 
	GetVillageById(ctx context.Context, id string) (resp *entity.Village, err error)
}