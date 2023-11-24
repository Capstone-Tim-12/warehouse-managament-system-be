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

	if req.Name == "" {
		err = errors.New(http.StatusBadRequest, "name is empty")
		fmt.Println("name is empty ", err)
		return
	}

	err = h.warehouseusecase.CreateWarehouse(ctx, req)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, nil)
}

func (h *WarehouseHandler) GetWarehouseById(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("warehouseId")

	data, err := h.warehouseusecase.GetWarehouse(ctx, id)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, data)
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

	if req.Name == "" {
		err = errors.New(http.StatusBadRequest, "name is empty")
		fmt.Println("name is empty ", err)
		return
	}

	err = h.warehouseusecase.UpdateWarehouseDetails(ctx, req, id)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, nil)
}
