package handler

import (
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/warehouse"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/warehouse/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/response"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
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

	err = h.warehouseusecase.CreateWarehouse(ctx, req, cast.ToString(clamsData.UserId))
	if err != nil {
		fmt.Println("error create warehouse: ", err.Error())
		err = errors.New(http.StatusInternalServerError, "error creating Warehouse")
		return
	}
	return response.NewSuccessResponse(c, nil)
}

func (h *WarehouseHandler) GetWarehouseById(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("warehouseId")

	fmt.Println(id)
	data, err := h.warehouseusecase.GetWarehouse(ctx, id)
	if err != nil {
		fmt.Println("error get warehouse: ", err.Error())
		err = errors.New(http.StatusInternalServerError, "error get Warehouse")
		return
	}
	return response.NewSuccessResponse(c, data)
}

func (h *WarehouseHandler) GetAllWarehouse(c echo.Context) (err error) {
	ctx := c.Request().Context()
	data, err := h.warehouseusecase.GetAllWarehouse(ctx)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, data)
}
