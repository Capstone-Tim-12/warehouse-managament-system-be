package user

import (
	"context"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultUser) GetAllProvince(ctx context.Context) (resp []model.RegionResponse, err error) {
	data, err := s.regionRepo.FindAllProvince(ctx)
	if err != nil {
		err = errors.New(http.StatusInternalServerError, "failed to get all province")
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
		err = errors.New(http.StatusNotFound, "province not found")
		return
	}

	data, err := s.regionRepo.FindRegencyByProvinceId(ctx, id)
	if err != nil {
		err = errors.New(http.StatusNotFound, "failed to get data regency")
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
		err = errors.New(http.StatusNotFound, "regency not found")
		return
	}

	data, err := s.regionRepo.FindDistrictByRegencyId(ctx, id)
	if err != nil {
		err = errors.New(http.StatusNotFound, "failed to get data distric")
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
