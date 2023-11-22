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
	UserRegister(ctx context.Context, req model.RegisterUserRequest) (resp model.RegisterUserResponse, err error)
	ResendOtp(ctx context.Context, req model.OtpRequest) (err error)
	Login(ctx context.Context, req model.LoginRequest) (resp model.LoginResponse, err error)
	VerificationUser(ctx context.Context, req model.VerificationUserRequest) (resp model.VerificationUserResponse, err error)
	ResetPassword(ctx context.Context, req model.ResetPasswordRequest) (err error)
	UpdateUsernameProfile(ctx context.Context, userId string, req model.UpdateUsernameProfileRequest) (err error)
	GetProfile(ctx context.Context, userId string) (resp model.GetProfileResponse, err error)
}
