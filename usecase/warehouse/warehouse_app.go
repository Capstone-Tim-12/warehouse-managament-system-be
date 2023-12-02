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
			Id:                list[i].ID,
			Name:              list[i].Name,
			DistrictName:      list[i].District.Name,
			RegencyName:       list[i].District.Regency.Name,
			ProvinceName:      list[i].District.Regency.Province.Name,
			SurfaceArea:       list[i].SurfaceArea,
			BuildingArea:      list[i].BuildingArea,
			WeeklyPrice:       math.Ceil(list[i].Price / 52),
			MonthlyPrice:      math.Ceil(list[i].Price / 12),
			AnnualPrice:       list[i].Price,
			Distance:          distance,
			Status:            string(list[i].Status),
			WarehouseTypeId:   list[i].WarehouseType.ID,
			WarehouseTypeName: list[i].WarehouseType.Name,
			Image:             image,
		})
	}

	return
}

func (s *defaultWarehouse) GetMywarehouse(ctx context.Context, userId int, status model.TrxStatus, param paginate.Pagination) (resp []model.MyWarehoyseResponse, count int64, err error) {
	trxData, count, err := s.paymentRepo.GetListTransactionByUserIdAndStatus(ctx, userId, model.GetStatusTrx[status], param)
	if err != nil {
		fmt.Println("error getting tlist warehouse data:", err.Error())
		err = errors.New(http.StatusInternalServerError, "error getting tlist warehouse data")
		return
	}

	for i := 0; i < len(trxData); i++ {
		var image string
		if len(trxData[i].Warehouse.WarehouseImg) != 0 {
			image = trxData[i].Warehouse.WarehouseImg[0].Image
		}
		resp = append(resp, model.MyWarehoyseResponse{
			TransactionId:     trxData[i].ID,
			TransactionStatus: string(trxData[i].Status),
			WarehouseId:       trxData[i].WarehouseID,
			WarehouseName:     trxData[i].Warehouse.Name,
			WatehouseRegency:  trxData[i].Warehouse.District.Regency.Name,
			SurfaceArea:       trxData[i].Warehouse.SurfaceArea,
			BuildingArea:      trxData[i].Warehouse.BuildingArea,
			WarehouseImage:    image,
		})
	}

	return
}

func (s *defaultWarehouse) GetWarehouseInfo(ctx context.Context, warehouseId string) (resp model.WarehouseInfoResponse, err error) {
	warehouseData, err := s.warehouseRepo.FindWarehouseById(ctx, warehouseId)
	if err != nil {
		fmt.Println("failed find warehouse")
		err = errors.New(http.StatusInternalServerError, "failed find warehouse")
		return
	}
	
	var image string
	if len(warehouseData.WarehouseImg) != 0 {
		image = warehouseData.WarehouseImg[0].Image
	}

	resp = model.WarehouseInfoResponse{
		Id:                warehouseData.ID,
		Name:              warehouseData.Name,
		DistrictName:      warehouseData.District.Name,
		RegencyName:       warehouseData.District.Regency.Name,
		ProvinceName:      warehouseData.District.Regency.Province.Name,
		AnnualPrice:       warehouseData.Price,
		Image:             image,
	}

	return
}
