package handler

import (
	"fmt"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
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

func (h *UserHandler) RegisterUserData(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.RegisterDataRequest
	err = c.Bind(&req)
	if err != nil {
		err = response.NewErrorResponse(c, err)
		fmt.Println("error bind register user data: ", err)
		return
	}

	if req.NIK == "" {
		err = response.NewErrorResponse(c, errors.ErrBadRequest)
		fmt.Println("nik is empty ", err)
		return
	}

	if req.Email == "" {
		err = response.NewErrorResponse(c, errors.ErrBadRequest)
		fmt.Println("email is empty ", err)
		return
	}

	err = h.userUsecase.RegisterData(ctx, req)
	if err != nil {
		err = response.NewErrorResponse(c, err)
		return
	}

	return response.NewSuccessResponse(c, nil)
}

func (h *UserHandler) RegisterUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.RegisterUserRequest
	err = c.Bind(&req)
	if err != nil {
		err = response.NewErrorResponse(c, err)
		fmt.Println("error bind register user data: ", err)
		return
	}
	if req.Email == "" {
		err = response.NewErrorResponse(c, errors.ErrBadRequest)
		fmt.Println("email is empty ", err)
		return
	}
	registerResponse, err := h.userUsecase.UserRegister(ctx, req)
	if err != nil {
		err = response.NewErrorResponse(c, err)
		return
	}
	return response.NewSuccessResponse(c, registerResponse)
}

func (h *UserHandler) ResendUserOTP(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.OtpRequest

	err = c.Bind(&req)

	if req.Email == "" {
		err = response.NewErrorResponse(c, errors.ErrBadRequest)
		fmt.Println("email is empty ", err)
		return
	}

	err = h.userUsecase.ResendOtp(ctx, req)
	if err != nil {
		err = response.NewErrorResponse(c, err)
		return
	}
	return response.NewSuccessResponse(c, nil)
}
