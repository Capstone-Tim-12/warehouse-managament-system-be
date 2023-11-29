package handler

import (
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/response"
	"github.com/labstack/echo/v4"
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

	return response.NewSuccessResponse(c, nil)
}

func (h *PaymentHandler) GetScheme(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)

	data, err := h.paymentUsecase.GetPaymentScheme(ctx, clamsData.UserId)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, data)
}
