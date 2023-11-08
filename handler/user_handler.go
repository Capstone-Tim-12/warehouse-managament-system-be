package handler

import (
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/response"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUsecase user.UserUsecase
}

func NewUserHandler(userUsecase user.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) GetAllProvince(c echo.Context) (err error) {
	ctx := c.Request().Context()
	data, err := h.userUsecase.GetAllProvince(ctx)
	if err != nil {
		err = response.NewErrorResponse(c, err)
		return
	}
	
	return response.NewSuccessResponse(c, data)
}

func (h *UserHandler) GetRegencyByProvinceId(c echo.Context) (err error) {
	ctx := c.Request().Context()
	provinceId := c.Param("provinceId")
	data, err := h.userUsecase.GetRegencyByProvinceId(ctx, provinceId)
	if err != nil {
		err = response.NewErrorResponse(c, err)
		return
	}
	
	return response.NewSuccessResponse(c, data)
}

func (h *UserHandler) GetDistricByRegencyId(c echo.Context) (err error) {
	ctx := c.Request().Context()
	regencyId := c.Param("regencyId")
	data, err := h.userUsecase.GetDistricByRegencyId(ctx, regencyId)
	if err != nil {
		err = response.NewErrorResponse(c, err)
		return
	}
	
	return response.NewSuccessResponse(c, data)
}