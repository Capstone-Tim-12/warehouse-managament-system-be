package user

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
)

type UserUsecase interface {
	GetAllProvince(ctx context.Context) (resp []model.RegionResponse, err error)
	GetRegencyByProvinceId(ctx context.Context, id string) (resp []model.RegionResponse, err error)
	GetDistricByRegencyId(ctx context.Context, id string) (resp []model.RegionResponse, err error)
	RegisterData(ctx context.Context, req model.RegisterDataRequest) (err error)
	UserRegister(ctx context.Context, req model.RegisterUserRequest) (resp model.RegionResponse, err error)
	ResendOtp(ctx context.Context, req model.OtpRequest) (err error)
}
