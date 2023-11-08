package regiondb

import "context"

type RegionRepository interface {
	FindAllProvince(ctx context.Context) (resp []Province, err error)
	FindRegencyByProvinceId(ctx context.Context, id string)(resp []Regency, err error)
	FindDistrictByRegencyId(ctx context.Context, id string) (resp []District, err error)
	FindVillageByDistrictId(ctx context.Context, id string) (resp []Village, err error)
	GetProvinceById(ctx context.Context, id string) (resp *Province, err error)
	GetRegencyById(ctx context.Context, id string) (resp *Regency, err error)
	GetDistrictById(ctx context.Context, id string) (resp *District, err error) 
	GetVillageById(ctx context.Context, id string) (resp *Village, err error)
}