package handler

import (
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/response"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type PaymentHandler struct {
	paymentUsecase payment.PaymentUsecase
}

func NewPaymentHandler(paymentUsecase payment.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{paymentUsecase: paymentUsecase}
}

func (h *PaymentHandler) SubmissionWarehouse(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.SubmissionRequest
	clamsData := utils.GetClamsJwt(c)
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

	err = h.paymentUsecase.SubmissionWarehouse(ctx, clamsData.UserId, req)
	if err != nil {
		return
	}

	return response.NewSuccessResponse(c, http.StatusOK, nil)
}

func (h *PaymentHandler) GetScheme(c echo.Context) (err error) {
	ctx := c.Request().Context()

	data, err := h.paymentUsecase.GetPaymentScheme(ctx)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *PaymentHandler) GetHistoryInstalmentUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	param, _ := paginate.GetParams(c)
	data, count, err := h.paymentUsecase.GetHistoryInstalmentUser(ctx, param)
	if err != nil {
		return
	}

	clamsData := utils.GetClamsJwt(c)
	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}

	resp := response.NewResponseSuccessPagination(float64(count), param, data)
	err = c.JSON(http.StatusOK, resp)
	return
}

func (h *PaymentHandler) GetListTransactionIdUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	userId := c.Param("userId")

	clamsData := utils.GetClamsJwt(c)
	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}

	data, err := h.paymentUsecase.GetListTransactionIdUser(ctx, cast.ToInt(userId))
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

