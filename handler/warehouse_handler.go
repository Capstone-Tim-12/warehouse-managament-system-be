package handler

import (
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/warehouse"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/warehouse/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/response"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/validation"
	"github.com/labstack/echo/v4"
)

type WarehouseHandler struct {
	warehouseusecase warehouse.WarehouseUsecase
}

func NewWarehouseHandler(warehouseusecase warehouse.WarehouseUsecase) *WarehouseHandler {
	return &WarehouseHandler{
		warehouseusecase: warehouseusecase,
	}
}

func (h *WarehouseHandler) CreateWarehouseDetail(c echo.Context) (err error) {
	var req model.WarehouseDataRequest
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)

	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invalid request")
		fmt.Println("error bind register warehouse data: ", err.Error())
		return
	}

	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}

	err = c.Validate(req)
	if err != nil {
		fmt.Println("error validate data: ", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	err = h.warehouseusecase.CreateWarehouse(ctx, req)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusCreated, nil)
}

func (h *WarehouseHandler) GetWarehouseById(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("warehouseId")

	data, err := h.warehouseusecase.GetWarehouse(ctx, id)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *WarehouseHandler) GetWarehouseList(c echo.Context) (err error) {
	ctx := c.Request().Context()
	param, _ := paginate.GetParams(c)
	clamsData := utils.GetClamsJwt(c)
	data, count, err := h.warehouseusecase.GetWarehouseList(ctx, param, clamsData.UserId)
	if err != nil {
		return
	}

	resp := response.NewResponseSuccessPagination(float64(count), param, data)
	err = c.JSON(http.StatusOK, resp)
	return
}

func (h *WarehouseHandler) UpdateWarehouseById(c echo.Context) (err error) {
	var req model.WarehouseDataRequest
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)
	id := c.Param("warehouseId")

	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}

	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invalid request")
		fmt.Println("error bind register warehouse data: ", err.Error())
		return
	}

	err = c.Validate(req)
	if err != nil {
		fmt.Println("error validate data: ", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	err = h.warehouseusecase.UpdateWarehouseDetails(ctx, req, id)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, nil)
}

func (h *WarehouseHandler) DeleteWarehouseById(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)
	id := c.Param("warehouseId")

	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}

	err = h.warehouseusecase.DeleteWarehouse(ctx, id)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, nil)
}

func (h *WarehouseHandler) GetListWarehouseType(c echo.Context) (err error){
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)
	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}
	data, err := h.warehouseusecase.GetListWarehouseType(ctx)
	if err != nil {
		return
	}

	return response.NewSuccessResponse(c, http.StatusOK, data)
	return
}

func (h *WarehouseHandler) UploadPhotoWarehouse(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)
	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}
	file, err := c.MultipartForm()
	if err != nil {
		fmt.Println(ctx, "error uploading product images", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	images := file.File["photos"]
	if len(images) == 0 {
		fmt.Println("image is empty")
		err = errors.New(http.StatusBadRequest, "image is empty")
		return
	}
	if len(images) > 6 {
		fmt.Println(ctx, "maximum of 6 images permitted")
		err = errors.New(http.StatusBadRequest, "maximum of 6 images permitted")
		return
	}
	for i := 0; i < len(images); i++ {
		err = validation.ValidationImages(images[i].Filename, int(images[i].Size))
		if err != nil {
			fmt.Println(ctx, "error validate file name", err.Error())
			err = errors.New(http.StatusBadRequest, err.Error())
			return
		}
	}

	data, err := h.warehouseusecase.UploadPhotoWarehouse(ctx, images)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(c, http.StatusOK, data)
}


func (h *WarehouseHandler) MywarehouseSubmitted(c echo.Context) (err error){
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)
	param, _ := paginate.GetParams(c)
	data, count, err := h.warehouseusecase.GetMywarehouse(ctx, clamsData.UserId, model.StatusSubmitted, param)
	if err != nil {
		return
	}

	resp := response.NewResponseSuccessPagination(float64(count), param, data)
	return c.JSON(http.StatusOK, resp)
}

func (h *WarehouseHandler) MywarehouseRented(c echo.Context) (err error){
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)
	param, _ := paginate.GetParams(c)
	data, count, err := h.warehouseusecase.GetMywarehouse(ctx, clamsData.UserId, model.StatusRented, param)
	if err != nil {
		return
	}

	resp := response.NewResponseSuccessPagination(float64(count), param, data)
	return c.JSON(http.StatusOK, resp)
}

func (h *WarehouseHandler) GetWarehouseInfo(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("warehouseId")

	data, err := h.warehouseusecase.GetWarehouseInfo(ctx, id)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}
