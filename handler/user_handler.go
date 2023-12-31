package handler

import (
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/response"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/validation"
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
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *UserHandler) GetRegencyByProvinceId(c echo.Context) (err error) {
	ctx := c.Request().Context()
	provinceId := c.Param("provinceId")
	data, err := h.userUsecase.GetRegencyByProvinceId(ctx, provinceId)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *UserHandler) GetDistricByRegencyId(c echo.Context) (err error) {
	ctx := c.Request().Context()
	regencyId := c.Param("regencyId")
	data, err := h.userUsecase.GetDistricByRegencyId(ctx, regencyId)
	if err != nil {
		return
	}

	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *UserHandler) RegisterUserData(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.RegisterDataRequest
	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invalid request")
		fmt.Println("error bind register user data: ", err.Error())
		return
	}

	err = c.Validate(req)
	if err != nil {
		fmt.Println("error validate data: ", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	err = h.userUsecase.RegisterData(ctx, req)
	if err != nil {
		return
	}

	return response.NewSuccessResponse(c, http.StatusCreated, nil)
}

func (h *UserHandler) RegisterUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.RegisterUserRequest

	longitude := c.Request().Header.Get("longitude")
	latitude := c.Request().Header.Get("latitude")

	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invalid request")
		fmt.Println("error bind register user data: ", err.Error())
		return
	}
	err = c.Validate(req)
	if err != nil {
		fmt.Println("error validate data: ", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}
	registerResponse, err := h.userUsecase.UserRegister(ctx, req, cast.ToFloat64(longitude), cast.ToFloat64(latitude))
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusCreated, registerResponse)
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

	err = c.Validate(req)
	if err != nil {
		fmt.Println("error validate data: ", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	err = h.userUsecase.ResendOtp(ctx, req)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, nil)
}

func (h *UserHandler) LoginUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.LoginRequest
	longitude := c.Request().Header.Get("longitude")
	latitude := c.Request().Header.Get("latitude")

	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invaid request")
		fmt.Println("error bind  data: ", err)
		return
	}

	err = c.Validate(req)
	if err != nil {
		fmt.Println("error validate data: ", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	userResponse, err := h.userUsecase.Login(ctx, req, cast.ToFloat64(latitude), cast.ToFloat64(longitude))
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, userResponse)
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
	err = c.Validate(req)
	if err != nil {
		fmt.Println("error validate data: ", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.userUsecase.VerificationUser(ctx, req)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
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

	err = c.Validate(req)
	if err != nil {
		fmt.Println("error validate data: ", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	err = h.userUsecase.ResetPassword(ctx, req)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, nil)
}

func (h *UserHandler) GetProfile(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)
	data, err := h.userUsecase.GetProfile(ctx, cast.ToString(clamsData.UserId))
	if err != nil {
		fmt.Println("failed to get profile", err)
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *UserHandler) UpdateUsername(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.UpdateUsernameProfileRequest
	clamsData := utils.GetClamsJwt(c)

	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invaid request")
		fmt.Println("error bind  data: ", err)
		return
	}

	err = c.Validate(req)
	if err != nil {
		fmt.Println("error validate data: ", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	err = h.userUsecase.UpdateUsernameProfile(ctx, cast.ToString(clamsData.UserId), req)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, nil)
}

func (h *UserHandler) UpdatePhotoProfile(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)

	var req model.UpdatePhotoProfileRequest
	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invaid request")
		fmt.Println("error bind  data: ", err)
		return
	}

	err = c.Validate(req)
	if err != nil {
		fmt.Println("error validate data: ", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	err = h.userUsecase.UpdatePhotoProfile(ctx, clamsData.UserId, req)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, nil)
}

func (h *UserHandler) UploadPhoto(c echo.Context) (err error) {
	ctx := c.Request().Context()

	file, err := c.FormFile("photo")
	if err != nil {
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	err = validation.ValidationImages(file.Filename, int(file.Size))
	if err != nil {
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.userUsecase.UploadPhoto(ctx, file)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *UserHandler) GetAvatarList(c echo.Context) (err error) {
	ctx := c.Request().Context()
	data, err := h.userUsecase.GetAvatarList(ctx)
	if err != nil {
		fmt.Println("failed to get profile", err)
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *UserHandler) UpdateEmail(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)

	var req model.OtpRequest
	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invaid request")
		fmt.Println("error bind  data: ", err)
		return
	}

	err = c.Validate(req)
	if err != nil {
		fmt.Println("error validate data: ", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	err = h.userUsecase.UpdateEmail(ctx, clamsData.UserId, req)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, nil)
}

func (h *UserHandler) DeleteUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)
	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}
	userId := c.Param("userId")
	err = h.userUsecase.DeleteUser(ctx, cast.ToInt(userId))
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, nil)
}

func (h *UserHandler) GetUserList(c echo.Context) (err error) {
	ctx := c.Request().Context()
	param, _ := paginate.GetParams(c)
	data, count, err := h.userUsecase.GetUserList(ctx, param)
	if err != nil {
		return
	}

	resp := response.NewResponseSuccessPagination(float64(count), param, data)
	err = c.JSON(http.StatusOK, resp)
	return
}

func (h *UserHandler) GetUserById(c echo.Context) (err error) {
	ctx := c.Request().Context()
	userId := c.Param("userId")
	clamsData := utils.GetClamsJwt(c)
	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}
	data, err := h.userUsecase.GetUserById(ctx, cast.ToInt(userId))
	if err != nil {
		fmt.Println("failed to get profile", err)
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *UserHandler) GetUserInfo(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)
	data, err := h.userUsecase.GetUserInfo(ctx, clamsData.UserId)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *UserHandler) ChatBot(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)

	var req model.ChatRequest
	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invaid request")
		fmt.Println("error bind  data: ", err)
		return
	}

	err = c.Validate(req)
	if err != nil {
		fmt.Println("error validate data: ", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.userUsecase.ChatBot(ctx, clamsData.UserId, req.Text)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *UserHandler) UpdateAdminUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.UserSettingRequest
	clamsData := utils.GetClamsJwt(c)
	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}

	err = c.Bind(&req)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invalid request")
		fmt.Println("error bind register user data: ", err.Error())
		return
	}

	err = c.Validate(req)
	if err != nil {
		fmt.Println("error validate data: ", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	err = h.userUsecase.UpdateAdminDasboarUser(ctx, clamsData.UserId, req)
	if err != nil {
		return
	}

	return response.NewSuccessResponse(c, http.StatusOK, nil)
}




