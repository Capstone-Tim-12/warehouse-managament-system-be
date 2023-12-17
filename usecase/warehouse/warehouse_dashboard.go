package warehouse

import (
	"context"
	"encoding/csv"
	"fmt"
	"math"
	"mime/multipart"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/warehouse/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/spf13/cast"
)

func (s *defaultWarehouse) CreateWarehouse(ctx context.Context, req model.WarehouseDataRequest) (err error) {
	_, err = s.regionRepo.GetDistrictById(ctx, req.DistrictID)
	if err != nil {
		fmt.Println("Error getting regency id", err.Error())
		err = errors.New(http.StatusNotFound, "district not found")
		return
	}
	_, err = s.warehouseRepo.GetWarehouseTypeById(ctx, req.WarehouseTypeId)
	if err != nil {
		fmt.Println("error getting type id", err.Error())
		err = errors.New(http.StatusNotFound, "warehouse type not found")
		return
	}
	warehouseData := entity.Warehouse{
		Name:            req.Name,
		Description:     req.Description,
		DistrictID:      req.DistrictID,
		Address:         req.Address,
		SurfaceArea:     req.SurfaceArea,
		BuildingArea:    req.BuildingArea,
		Owner:           req.Owner,
		PhoneNumber:     req.PhoneNumber,
		Longitude:       req.Longitude,
		Latitude:        req.Latitude,
		Status:          entity.Available,
		Price:           req.Price,
		WarehouseTypeID: req.WarehouseTypeId,
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

func (s *defaultWarehouse) GetWarehouse(ctx context.Context, id int) (resp *model.WarehouseDataResponse, err error) {
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
		Name:            warehouseData.Name,
		Description:     warehouseData.Description,
		ProvinceID:      warehouseData.District.Regency.ProvinceID,
		ProvinceName:    warehouseData.District.Regency.Province.Name,
		RegencyID:       warehouseData.District.RegencyID,
		RegencyName:     warehouseData.District.Regency.Name,
		DistrictID:      warehouseData.DistrictID,
		DistrictName:    warehouseData.District.Name,
		Address:         warehouseData.Address,
		SurfaceArea:     warehouseData.SurfaceArea,
		BuildingArea:    warehouseData.BuildingArea,
		Owner:           warehouseData.Owner,
		PhoneNumber:     warehouseData.PhoneNumber,
		Longitude:       warehouseData.Longitude,
		Latitude:        warehouseData.Latitude,
		Status:          string(warehouseData.Status),
		WeeklyPrice:     math.Ceil(warehouseData.Price / 52),
		MonthlyPrice:    math.Ceil(warehouseData.Price / 12),
		AnnualPrice:     warehouseData.Price,
		WarehouseTypeId: warehouseData.WarehouseType.ID,
		WarehouseType:   warehouseData.WarehouseType.Name,
		Image:           images,
	}

	return
}

func (s *defaultWarehouse) UpdateWarehouseDetails(ctx context.Context, req model.WarehouseDataRequest, id int) (err error) {
	_, err = s.regionRepo.GetDistrictById(ctx, req.DistrictID)
	if err != nil {
		fmt.Println("Error getting regency id", err.Error())
		err = errors.New(http.StatusNotFound, "district not found")
		return
	}

	_, err = s.warehouseRepo.GetWarehouseTypeById(ctx, req.WarehouseTypeId)
	if err != nil {
		fmt.Println("error getting type id", err.Error())
		err = errors.New(http.StatusNotFound, "warehouse type not found")
		return
	}

	warehouseData, err := s.warehouseRepo.FindWarehouseByIdOnly(ctx, cast.ToString(id))
	if err != nil {
		fmt.Println("failed to get data warehouse")
		err = errors.New(http.StatusNotFound, "failed to get data warehouse")
		return
	}

	tx := s.warehouseRepo.BeginTrans(ctx)
	err = s.warehouseRepo.DeleteWarehouseImgByWarehouseId(ctx, tx, warehouseData.ID)
	if err != nil {
		tx.Rollback()
		fmt.Println("error delete data: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	
	for _, img := range req.Image {
		warehouseData.WarehouseImg = append(warehouseData.WarehouseImg, entity.WarehouseImg{
			Image:       img,
			WarehouseID: warehouseData.ID,
		})
	}

	warehouseData.Status = entity.NotAvailable
	if req.IsAvailable {
		warehouseData.Status = entity.Available
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
	warehouseData.WarehouseTypeID = req.WarehouseTypeId

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

func (s *defaultWarehouse) DeleteWarehouse(ctx context.Context, id string) (err error) {
	warehouseData, err := s.warehouseRepo.FindWarehouseByIdOnly(ctx, id)
	if err != nil {
		fmt.Println("error found warehouse: ", err.Error())
		err = errors.New(http.StatusNotFound, "warehouse not found")
		return
	}

	err = s.warehouseRepo.DeleteWarehouse(ctx, warehouseData)
	if err != nil {
		fmt.Println("error delete warehouse: ", err.Error())
		err = errors.New(http.StatusInternalServerError, "failed delete warehouse")
		return
	}
	return
}

func (s *defaultWarehouse) GetListWarehouseType(ctx context.Context) (resp []model.WarehouseTypeResponse, err error) {
	warehouseData, err := s.warehouseRepo.GetListWarehouseType(ctx)
	if err != nil {
		fmt.Println("failed find warehouse")
		err = errors.New(http.StatusInternalServerError, "failed find warehouse")
		return
	}

	for i := 0; i < len(warehouseData); i++ {
		resp = append(resp, model.WarehouseTypeResponse{
			Id:   warehouseData[i].ID,
			Name: warehouseData[i].Name,
		})
	}

	return
}

func (s *defaultWarehouse) UploadPhotoWarehouse(ctx context.Context, photo []*multipart.FileHeader) (resp model.UploadPhotoResponse, err error) {
	for i := 0; i < len(photo); i++ {
		data, errRes := s.coreWrapper.UploadImage(ctx, photo[i])
		if errRes != nil {
			fmt.Println("failed upload image: ", errRes.Error())
			err = errors.New(http.StatusInternalServerError, "failed upload photo")
			return
		}
		if len(data.Data.Images) != 0 {
			resp.Images = append(resp.Images, data.Data.Images...)
		}
	}
	return
}

// name,description,districId,address,surfaceArea,buildingArea,owner,phoneNumber,longitude,latitude,price,warehouseTypeId,status
func (s *defaultWarehouse) ImportCsvFileWarehouse(ctx context.Context, file *multipart.FileHeader) (err error) {
	src, err := file.Open()
	if err != nil {
		fmt.Println("error opening file: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	defer src.Close()
	csvRead := csv.NewReader(src)

	rows, err := csvRead.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	tx := s.warehouseRepo.BeginTrans(ctx)
	for _, row := range rows {
		districtId := row[2]
		warehouseTypeId := cast.ToInt(row[11])
		_, err = s.regionRepo.GetDistrictById(ctx, districtId)
		if err != nil {
			fmt.Println("Error getting regency id", err.Error())
			err = errors.New(http.StatusNotFound, "district not found")
			return
		}
		_, err = s.warehouseRepo.GetWarehouseTypeById(ctx, warehouseTypeId)
		if err != nil {
			fmt.Println("error getting type id", err.Error())
			err = errors.New(http.StatusNotFound, "warehouse type not found")
			return
		}
		data := entity.Warehouse{
			Name:            row[0],
			Longitude:       cast.ToFloat64(row[8]),
			Latitude:        cast.ToFloat64(row[9]),
			DistrictID:      districtId,
			Address:         row[3],
			BuildingArea:    cast.ToFloat64(row[5]),
			SurfaceArea:     cast.ToFloat64(row[4]),
			Owner:           row[6],
			PhoneNumber:     row[7],
			Price:           cast.ToFloat64(row[10]),
			Description:     row[1],
			WarehouseTypeID: warehouseTypeId,
			Status:          entity.WarehouseStatus(row[12]),
		}
		err = s.warehouseRepo.CreateDetail(ctx, tx, &data)
		if err != nil {
			tx.Rollback()
			fmt.Println("Error creating warehouse: ", err.Error())
			err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	}

	tx.Commit()
	return
}
