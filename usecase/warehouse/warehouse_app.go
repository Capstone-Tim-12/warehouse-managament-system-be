package warehouse

import (
	"context"
	"fmt"
	"math"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/warehouse/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/calculate"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
)

func (s *defaultWarehouse) GetWarehouseList(ctx context.Context, param paginate.Pagination, userId int) (resp []model.WarehouseListResponse, count int64, err error) {
	userData, err := s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		fmt.Println("error getting user")
		err = errors.New(http.StatusNotFound, "user not found")
		return
	}
	list, count, err := s.warehouseRepo.FindWarehouseList(ctx, param, userData.Longitude, userData.Latitude)
	if err != nil {
		fmt.Println("error getting warehouseList: ", err.Error())
		err = errors.New(http.StatusInternalServerError, "error getting warehouse list")
		return
	}

	for i := 0; i < len(list); i++ {
		var image string
		if len(list[i].WarehouseImg) != 0 {
			image = list[i].WarehouseImg[0].Image
		}

		distance := calculate.Haversine(userData.Latitude, userData.Longitude, list[i].Latitude, list[i].Longitude)
		resp = append(resp, model.WarehouseListResponse{
			Id:           list[i].ID,
			Name:         list[i].Name,
			DistrictName: list[i].District.Name,
			RegencyName:  list[i].District.Regency.Name,
			ProvinceName: list[i].District.Regency.Province.Name,
			SurfaceArea:  list[i].SurfaceArea,
			BuildingArea: list[i].BuildingArea,
			WeeklyPrice:  math.Ceil(list[i].Price / 52),
			MonthlyPrice: math.Ceil(list[i].Price / 12),
			AnnualPrice:  list[i].Price,
			Distance:     distance,
			Image:        image,
		})
	}

	return
}
