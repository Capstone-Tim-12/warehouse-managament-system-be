package user

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
)

type UserUsecase interface {
	GetAllProvince(ctx context.Context) (resp []model.RegionResponse, err error) 
	GetRegencyByProvinceId(ctx context.Context, id string) (resp []model.RegionResponse, err error)
	GetDistricByRegencyId(ctx context.Context, id string) (resp []model.RegionResponse, err error)
}