package warehouse

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/warehouse/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/spf13/cast"
)

func (s *defaultWarehouse) CreateWarehouse(ctx context.Context, req model.WarehouseDataRequest, userId string) (err error) {
	userData, err := s.userRepo.GetUserById(ctx, cast.ToInt(userId))
	if err != nil {
		fmt.Printf("user not found")
		err = errors.New(http.StatusNotFound, "user not found")
		return
	}

	if userData.Role != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusForbidden, "role is not admin")
		return
	}

	warehouseData := entity.Warehouse{
		Name:         req.Name,
		Description:  req.Description,
		ProvinceID:   req.ProvinceID,
		RegencyID:    req.RegencyID,
		DistrictID:   req.DistrictID,
		Address:      req.Address,
		SurfaceArea:  req.SurfaceArea,
		BuildingArea: req.BuildingArea,
		Owner:        req.Owner,
		PhoneNumber:  req.PhoneNumber,
		Longitude:    req.Longitude,
		Latitude:     req.Latitude,
	}
	tx := s.warehouseRepo.BeginTrans(ctx)
	err = s.warehouseRepo.CreateDetail(ctx, tx, &warehouseData)
	if err != nil {
		tx.Rollback()
		err = errors.New(http.StatusInternalServerError, "failed create warehouse")
		fmt.Println("failed create warehouse")
		return
	}

	var warehouseImg *entity.WarehouseImg
	for _, img := range req.Image {
		warehouseImg = &entity.WarehouseImg{
			Image:       img,
			WarehouseID: warehouseData.ID,
		}
	}

	err = s.warehouseRepo.CreateImg(ctx, tx, warehouseImg)
	if err != nil {
		tx.Rollback()
		err = errors.New(http.StatusInternalServerError, "failed create warehouse image")
		fmt.Println("failed create warehouse image")
		return
	}

	tx.Commit()
	return
}

func (s *defaultWarehouse) GetWarehouse(ctx context.Context, id string) (resp *model.WarehouseDataResponse, err error) {
	warehouseData, err := s.warehouseRepo.FindWarehouseById(ctx, id)
	if err != nil {
		fmt.Println("failed find warehouse")
		err = errors.New(http.StatusInternalServerError, "failed find warehouse")
		return
	}
	provinceData, err := s.regionRepo.GetProvinceById(ctx, cast.ToString(warehouseData.ProvinceID))
	if err != nil {
		fmt.Println("failed find province")
		err = errors.New(http.StatusInternalServerError, "failed find province")
		return
	}
	regencyData, err := s.regionRepo.GetRegencyById(ctx, cast.ToString(warehouseData.RegencyID))
	if err != nil {
		fmt.Println("failed find regency")
		err = errors.New(http.StatusInternalServerError, "failed find regency")
		return
	}
	districtData, err := s.regionRepo.GetDistrictById(ctx, cast.ToString(warehouseData.DistrictID))
	if err != nil {
		fmt.Println("failed find distric")
		err = errors.New(http.StatusInternalServerError, "failed find distric")
		return
	}

	var images []string
	for _, warehouseImage := range warehouseData.WareHouseImg {
		images = append(images, warehouseImage.Image)
	}

	resp = &model.WarehouseDataResponse{
		Name:         warehouseData.Name,
		Description:  warehouseData.Description,
		ProvinceID:   provinceData.ID,
		ProvinceName: provinceData.Name,
		RegencyID:    regencyData.ID,
		RegencyName:  regencyData.Name,
		DistrictID:   districtData.ID,
		DistrictName: districtData.Name,
		Address:      warehouseData.Address,
		SurfaceArea:  warehouseData.SurfaceArea,
		BuildingArea: warehouseData.BuildingArea,
		Owner:        warehouseData.Owner,
		PhoneNumber:  warehouseData.PhoneNumber,
		Longitude:    warehouseData.Longitude,
		Latitude:     warehouseData.Latitude,
		Status:       string(warehouseData.Status),
		Image:        images,
	}

	return
}
