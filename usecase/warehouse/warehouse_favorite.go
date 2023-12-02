package warehouse

import (
	"context"
	"fmt"
	"math"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/warehouse/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/calculate"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
	"github.com/spf13/cast"
)

func (s *defaultWarehouse) AddFavorite(ctx context.Context, userId int, req model.AddFavoritRequest) (err error) {
	_, err = s.warehouseRepo.FindWarehouseByIdOnly(ctx, cast.ToString( req.WarehouseId))
	if err != nil {
		fmt.Println("warehouse not found: ", err.Error())
		err = errors.New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	wrData, err := s.warehouseRepo.FindFavoritByWarehouseIdAndUserId(ctx, req.WarehouseId, userId)
	if wrData.WarehouseID != 0 {
		return
	}

	reqFavorit := entity.Favorit{
		UserID:      userId,
		WarehouseID: req.WarehouseId,
	}
	err = s.warehouseRepo.AddFavorit(ctx, &reqFavorit)
	if err != nil {
		fmt.Println("failed to add favorites: ", err.Error())
		err = errors.New(http.StatusInternalServerError, "failed to add favorites: ")
		return
	}

	return
}

func (s *defaultWarehouse) DeleteFavorit(ctx context.Context, favoritId int) (err error) {
	data, err := s.warehouseRepo.FindFavoritById(ctx, favoritId)
	if err != nil {
		fmt.Println("failed to find favorites: ", err.Error())
		err = errors.New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	err = s.warehouseRepo.DeleteFavorite(ctx, data.ID)
	if err != nil {
		fmt.Println("failed to delete favorites: ", err.Error())
		err = errors.New(http.StatusInternalServerError, "failed to delete favorites")
		return
	}

	return
}

func (s *defaultWarehouse) GetListFavorite(ctx context.Context, userId int, param paginate.Pagination) (resp []model.WarehouseListResponse, count int64, err error) {
	data, count, err := s.warehouseRepo.FindListFavoriteByUserId(ctx, userId, param)
	if err != nil {
		fmt.Println("failed to get list favorites: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	for i := 0; i < len(data); i++ {
		var image string
		if len(data[i].Warehouse.WarehouseImg) != 0 {
			image = data[i].Warehouse.WarehouseImg[0].Image
		}

		distance := calculate.Haversine(data[i].User.Latitude, data[i].User.Longitude, data[i].Warehouse.Latitude, data[i].Warehouse.Longitude)
		resp = append(resp, model.WarehouseListResponse{
			Id:                data[i].WarehouseID,
			FavoritId:         data[i].ID,
			Name:              data[i].Warehouse.Name,
			DistrictName:      data[i].Warehouse.District.Name,
			RegencyName:       data[i].Warehouse.District.Regency.Name,
			ProvinceName:      data[i].Warehouse.District.Regency.Province.Name,
			SurfaceArea:       data[i].Warehouse.SurfaceArea,
			BuildingArea:      data[i].Warehouse.BuildingArea,
			WeeklyPrice:       math.Ceil(data[i].Warehouse.Price / 52),
			MonthlyPrice:      math.Ceil(data[i].Warehouse.Price / 12),
			AnnualPrice:       data[i].Warehouse.Price,
			Distance:          distance,
			Status:            string(data[i].Warehouse.Status),
			WarehouseTypeId:   data[i].Warehouse.WarehouseTypeID,
			WarehouseTypeName: data[i].Warehouse.WarehouseType.Name,
			Image:             image,
		})
	}

	return
}
