package warehouse

import (
	"context"
	"fmt"
	"math"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/warehouse/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultWarehouse) CreateWarehouse(ctx context.Context, req model.WarehouseDataRequest) (err error) {
	_, err = s.regionRepo.GetDistrictById(ctx, req.DistrictID)
	if err != nil {
		fmt.Println("Error getting regency id", err.Error())
		err = errors.New(http.StatusNotFound, "district not found")
		return
	}
	warehouseData := entity.Warehouse{
		Name:         req.Name,
		Description:  req.Description,
		DistrictID:   req.DistrictID,
		Address:      req.Address,
		SurfaceArea:  req.SurfaceArea,
		BuildingArea: req.BuildingArea,
		Owner:        req.Owner,
		PhoneNumber:  req.PhoneNumber,
		Longitude:    req.Longitude,
		Latitude:     req.Latitude,
		Status:       entity.Available,
		Price:        req.Price,
	}

	for _, img := range req.Image {
		warehouseData.WarehouseImg = append(warehouseData.WarehouseImg, entity.WarehouseImg{
			Image:       img,
			WarehouseID: warehouseData.ID,
		})
	}

	tx := s.warehouseRepo.BeginTrans(ctx)
	err = s.warehouseRepo.CreateDetail(ctx, tx, &warehouseData)
	if err != nil {
		tx.Rollback()
		err = errors.New(http.StatusInternalServerError, err.Error())
		fmt.Println("failed create warehouse")
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

	var images []string
	for i := 0; i < len(warehouseData.WarehouseImg); i++ {
		images = append(images, warehouseData.WarehouseImg[i].Image)
	}

	resp = &model.WarehouseDataResponse{
		Name:         warehouseData.Name,
		Description:  warehouseData.Description,
		ProvinceID:   warehouseData.District.Regency.ProvinceID,
		ProvinceName: warehouseData.District.Regency.Province.Name,
		RegencyID:    warehouseData.District.RegencyID,
		RegencyName:  warehouseData.District.Regency.Name,
		DistrictID:   warehouseData.DistrictID,
		DistrictName: warehouseData.District.Name,
		Address:      warehouseData.Address,
		SurfaceArea:  warehouseData.SurfaceArea,
		BuildingArea: warehouseData.BuildingArea,
		Owner:        warehouseData.Owner,
		PhoneNumber:  warehouseData.PhoneNumber,
		Longitude:    warehouseData.Longitude,
		Latitude:     warehouseData.Latitude,
		Status:       string(warehouseData.Status),
		WeeklyPrice:  math.Ceil(warehouseData.Price / 52),
		MonthlyPrice: math.Ceil(warehouseData.Price / 12),
		AnnualPrice:  warehouseData.Price,
		Image:        images,
	}

	return
}

func (s *defaultWarehouse) UpdateWarehouseDetails(ctx context.Context, req model.WarehouseDataRequest, id string) (err error) {
	_, err = s.regionRepo.GetDistrictById(ctx, req.DistrictID)
	if err != nil {
		fmt.Println("Error getting regency id", err.Error())
		err = errors.New(http.StatusNotFound, "district not found")
		return
	}

	warehouseData, err := s.warehouseRepo.FindWarehouseById(ctx, id)
	if err != nil {
		fmt.Println("failed to get data warehouse")
		err = errors.New(http.StatusNotFound, "failed to get data warehouse")
		return
	}

	warehouseData.Name = req.Name
	warehouseData.Description = req.Description
	warehouseData.DistrictID = req.DistrictID
	warehouseData.Address = req.Address
	warehouseData.SurfaceArea = req.SurfaceArea
	warehouseData.BuildingArea = req.BuildingArea
	warehouseData.Owner = req.Owner
	warehouseData.PhoneNumber = req.PhoneNumber
	warehouseData.Longitude = req.Longitude
	warehouseData.Latitude = req.Latitude
	warehouseData.Status = entity.WarehouseStatus(req.Status)

	for _, img := range req.Image {
		warehouseData.WarehouseImg = append(warehouseData.WarehouseImg, entity.WarehouseImg{
			Image:       img,
			WarehouseID: warehouseData.ID,
		})
	}

	tx := s.warehouseRepo.BeginTrans(ctx)
	err = s.warehouseRepo.DeleteWarehouseImgByWarehouseId(ctx, tx, warehouseData.ID)
	if err != nil {
		tx.Rollback()
		fmt.Println("error delete data: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	err = s.warehouseRepo.UpdateWarehouse(ctx, tx, warehouseData)
	if err != nil {
		tx.Rollback()
		err = errors.New(http.StatusInternalServerError, "failed update warehouse data")
		fmt.Println("failed update warehouse data: ", err.Error())
		return
	}

	tx.Commit()
	return
}
