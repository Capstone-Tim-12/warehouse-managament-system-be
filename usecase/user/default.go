package user

import (
	"context"
	"errors"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/regiondb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
)

type defaultUser struct {
	regionRepo regiondb.RegionRepository
}

func NewUserUsecase(regionRepo regiondb.RegionRepository) *defaultUser {
	return &defaultUser{
		regionRepo: regionRepo,
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
		err= errors.New("failed to get data regency")
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
		err= errors.New("failed to get data distric")
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
