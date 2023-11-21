package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/response"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
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
		return
	}
	return response.NewSuccessResponse(c, data)
}

func (h *UserHandler) GetRegencyByProvinceId(c echo.Context) (err error) {
	ctx := c.Request().Context()
	provinceId := c.Param("provinceId")
	data, err := h.userUsecase.GetRegencyByProvinceId(ctx, provinceId)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, data)
}

func (h *UserHandler) GetDistricByRegencyId(c echo.Context) (err error) {
	ctx := c.Request().Context()
	regencyId := c.Param("regencyId")
	data, err := h.userUsecase.GetDistricByRegencyId(ctx, regencyId)
	if err != nil {
		return
	}

	return response.NewSuccessResponse(c, data)
}

func (h *UserHandler) RegisterUserData(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.RegisterDataRequest
	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invalid request")
		fmt.Println("error bind register user data: ", err)
		return
	}

	if req.NIK == "" {
		err = errors.New(http.StatusBadRequest, "nik is empty")
		fmt.Println("nik is empty ", err)
		return
	}

	if !strings.Contains(req.Email, "@" ) {
		err = errors.New(http.StatusBadRequest, "format email is invalid")
		fmt.Println("email not valid")
		return
	}

	err = h.userUsecase.RegisterData(ctx, req)
	if err != nil {
		return
	}

	return response.NewSuccessResponse(c, nil)
}

func (h *UserHandler) RegisterUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.RegisterUserRequest
	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invalid request")
		fmt.Println("error bind register user data: ", err.Error())
		return
	}
	if req.Username == "" {
		err = errors.New(http.StatusBadRequest, "username must be filled in")
		fmt.Println("username is empty")
		return
	}
	if !strings.Contains(req.Email, "@" ) {
		err = errors.New(http.StatusBadRequest, "format email is invalid")
		fmt.Println("email not valid")
		return
	}
	registerResponse, err := h.userUsecase.UserRegister(ctx, req)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, registerResponse)
}

func (h *UserHandler) ResendUserOTP(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.OtpRequest

	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invaid request")
		fmt.Println("error bind  data: ", err)
		return
	}
	
	if req.Email == "" {
		err = errors.New(http.StatusBadRequest, "email is empty")
		fmt.Println("email is empty ", err)
		return
	}

	err = h.userUsecase.ResendOtp(ctx, req)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, nil)
}

func (h *UserHandler) LoginUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.LoginRequest

	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invaid request")
		fmt.Println("error bind  data: ", err)
		return
	}
	if req.Email == "" {
		err = errors.New(http.StatusBadRequest, "email is empty")
		fmt.Println("Email is empty ", err)
		return
	}

	if req.Password == "" {
		err = errors.New(http.StatusBadRequest, "password is empty")
		fmt.Println("password is empty ", err)
		return
	}

	userResponse, err := h.userUsecase.Login(ctx, req)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, userResponse)
}

func (h *UserHandler) VerificationOtpUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.VerificationUserRequest
	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invaid request")
		fmt.Println("error bind register user data: ", err)
		return
	}
	if req.Email == "" {
		err = errors.New(http.StatusBadRequest, "email is empty")
		fmt.Println("email is empty ", err)
		return
	}
	data, err := h.userUsecase.VerificationUser(ctx, req)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, data)
}

func (h *UserHandler) ResetPassword(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.ResetPasswordRequest

	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invaid request")
		fmt.Println("error bind  data: ", err)
		return
	}
	
	if req.Email == "" {
		err = errors.New(http.StatusBadRequest, "email is empty")
		fmt.Println("email is empty ", err)
		return
	}

	err = h.userUsecase.ResetPassword(ctx, req)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, nil)
}

func (h *UserHandler) GetProfile(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)
	data, err := h.userUsecase.GetProfile(ctx, cast.ToString(clamsData.UserId))
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, data)
}

