package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/regiondb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/userdb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/byrpt"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/costrans"
	customError "github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

type defaultUser struct {
	regionRepo regiondb.RegionRepository
	userRepo   userdb.UserRepository
}

func NewUserUsecase(regionRepo regiondb.RegionRepository, userRepo userdb.UserRepository) *defaultUser {
	return &defaultUser{
		regionRepo: regionRepo,
		userRepo:   userRepo,
	}
}

func (s *defaultUser) GetAllProvince(ctx context.Context) (resp []model.RegionResponse, err error) {
	data, err := s.regionRepo.FindAllProvince(ctx)
	if err != nil {
		err = errors.New("failed to get all province")
		return
	}
	for i := 0; i < len(data); i++ {
		resp = append(resp, model.RegionResponse{
			Id:   data[i].ID,
			Name: data[i].Name,
		})
	}
	return
}

func (s *defaultUser) GetRegencyByProvinceId(ctx context.Context, id string) (resp []model.RegionResponse, err error) {
	_, err = s.regionRepo.GetProvinceById(ctx, id)
	if err != nil {
		err = errors.New("province not found")
		return
	}

	data, err := s.regionRepo.FindRegencyByProvinceId(ctx, id)
	if err != nil {
		err = errors.New("failed to get data regency")
		return
	}

	for i := 0; i < len(data); i++ {
		resp = append(resp, model.RegionResponse{
			Id:   data[i].ID,
			Name: data[i].Name,
		})
	}
	return
}

func (s *defaultUser) GetDistricByRegencyId(ctx context.Context, id string) (resp []model.RegionResponse, err error) {
	_, err = s.regionRepo.GetRegencyById(ctx, id)
	if err != nil {
		err = errors.New("regency not found")
		return
	}

	data, err := s.regionRepo.FindDistrictByRegencyId(ctx, id)
	if err != nil {
		err = errors.New("failed to get data distric")
		return
	}

	for i := 0; i < len(data); i++ {
		resp = append(resp, model.RegionResponse{
			Id:   data[i].ID,
			Name: data[i].Name,
		})
	}
	return
}

func (s *defaultUser) RegisterData(ctx context.Context, req model.RegisterDataRequest) (err error) {
	userData, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		err = customError.ErrNotFound
		fmt.Println("Error getting Email", err)
		return
	}

	_, err = s.regionRepo.GetProvinceById(ctx, req.ProvinceID)
	if err != nil {
		err = customError.ErrNotFound
		fmt.Println("Error getting province id", err)
		return
	}
	_, err = s.regionRepo.GetRegencyById(ctx, req.RegencyID)
	if err != nil {
		err = customError.ErrNotFound
		fmt.Println("Error getting regency id", err)
		return
	}
	_, err = s.regionRepo.GetDistrictById(ctx, req.DistrictID)
	if err != nil {
		err = customError.ErrNotFound
		fmt.Println("Error getting regency id", err)
		return
	}

	createUserData := userdb.UserDetail{
		NIK:          req.NIK,
		Address:      req.Address,
		Country:      "Indonesia",
		FullName:     req.FullName,
		Gender:       req.Gender,
		PlaceOfBirth: req.PlaceBirth,
		DateBirth:    req.DateBirth,
		Work:         req.Work,
		Citizenship:  req.Citizenship,
		UserID:       userData.ID,
		ProvinceID:   req.ProvinceID,
		RegencyID:    req.RegencyID,
		DistrictID:   req.DistrictID,
	}

	err = s.userRepo.Create(ctx, &createUserData)
	if err != nil {
		err = errors.New("internal error create user data")
		fmt.Println("Internal error create user data")
		return
	}
	return
}

func (s *defaultUser) UserRegister(ctx context.Context, req model.RegisterUserRequest) (resp model.RegionResponse, err error) {
	userdata, _ := s.userRepo.GetUserByEmail(ctx, req.Email)
	if userdata.Email != "" {
		err = errors.New("email already exists")
		fmt.Println("email already exists")
		return
	}
	passwordByrpt := byrpt.HashSHA256(costrans.KeyPassword, req.Password)
	createUser := userdb.User{
		Username: req.Username,
		Password: passwordByrpt,
		Email:    req.Email,
	}
	err = s.userRepo.CreateUser(ctx, &createUser)
	if err != nil {
		err = errors.New("failed create data user")
		fmt.Println("failed create data user")
		return
	}
	return
}
