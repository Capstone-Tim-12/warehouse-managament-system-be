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
	
	clamsData := utils.GetClamsJwt(c)
	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}
	
	data, count, err := h.paymentUsecase.GetHistoryInstalmentUser(ctx, param)
	if err != nil {
		return
	}
	resp := response.NewResponseSuccessPagination(float64(count), param, data)
	err = c.JSON(http.StatusOK, resp)
	return
}

func (h *PaymentHandler) GetListTrxUserDasboar(c echo.Context) (err error) {
	ctx := c.Request().Context()
	userId := c.Param("userId")
	param, _ := paginate.GetParams(c)

	clamsData := utils.GetClamsJwt(c)
	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}

	data, count, err := h.paymentUsecase.GetListTrxUserDasboar(ctx, cast.ToInt(userId), param)
	if err != nil {
		return
	}
	
	resp := response.NewResponseSuccessPagination(float64(count), param, data)
	err = c.JSON(http.StatusOK, resp)
	return
}

func (h *PaymentHandler) GetAllTransaction(c echo.Context) (err error) {
	ctx := c.Request().Context()
	param := paginate.PaginationTrx{
		Page:       cast.ToInt(c.QueryParam("page")),
		Limit:      cast.ToInt(c.QueryParam("limit")),
		Search:     c.QueryParam("search"),
		ProvinceId: cast.ToInt(c.QueryParam("provinceId")),
		Status:     c.QueryParam("status"),
	}
	data, count, err := h.paymentUsecase.GetAllTransaction(ctx, param)
	if err != nil {
		return
	}

	clamsData := utils.GetClamsJwt(c)
	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}

	resp := response.NewResponseSuccessPaginationTrx(float64(count), param, data)
	err = c.JSON(http.StatusOK, resp)
	return
}

func (h *PaymentHandler) TransactionApproved(c echo.Context) (err error) {
	ctx := c.Request().Context()
	trxId := c.Param("transactionId")

	clamsData := utils.GetClamsJwt(c)
	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}

	err = h.paymentUsecase.TransactionApproved(ctx, trxId)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, nil)
}

func (h *PaymentHandler) TransactionRejected(c echo.Context) (err error) {
	ctx := c.Request().Context()
	trxId := c.Param("transactionId")

	clamsData := utils.GetClamsJwt(c)
	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}

	err = h.paymentUsecase.TransactionRejected(ctx, trxId)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, nil)
}

func (h *PaymentHandler) GetTransactionListDetail(c echo.Context) (err error) {
	ctx := c.Request().Context()
	trxId := c.Param("transactionId")

	clamsData := utils.GetClamsJwt(c)
	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}

	data, err := h.paymentUsecase.GetTransactionListDetail(ctx, trxId)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *PaymentHandler) GetListInstalment(c echo.Context) (err error) {
	ctx := c.Request().Context()
	trxId := c.Param("transactionId")

	param, _ := paginate.GetParams(c)
	data, count, err :=h.paymentUsecase.GetListInstalmentByTrxId(ctx, trxId, param)
	if err != nil {
		return
	}

	resp := response.NewResponseSuccessPagination(float64(count), param, data)
	return c.JSON(http.StatusOK, resp)
}

func (h *PaymentHandler) GetTransactionInfo(c echo.Context) (err error) {
	ctx := c.Request().Context()
	trxId := c.Param("transactionId")

	data, err := h.paymentUsecase.GetTransactionInfo(ctx, trxId)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *PaymentHandler) GetListPaymentMethod(c echo.Context) (err error) {
	ctx := c.Request().Context()

	data, err := h.paymentUsecase.GetListPaymentMethod(ctx)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *PaymentHandler) GetBankVa(c echo.Context) (err error) {
	ctx := c.Request().Context()

	data, err := h.paymentUsecase.GetBankVa(ctx)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *PaymentHandler) PaymentCheckout(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.PaymentRequest
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

	data, err := h.paymentUsecase.PaymentCheckout(ctx, clamsData.UserId, req)
	if err != nil {
		return
	}

	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *PaymentHandler) VaCallback(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.VaCallbackRequest
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

	err = h.paymentUsecase.VaCallback(ctx, req)
	if err != nil {
		return
	}

	return response.NewSuccessResponse(c, http.StatusOK, nil)
}

func (h *PaymentHandler) GetTransactionDetailDasboardUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	trxId := c.Param("transactionId")

	clamsData := utils.GetClamsJwt(c)
	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}

	data, err := h.paymentUsecase.GetTransactionDetailDasboardUser(ctx, trxId)
	if err != nil {
		return
	}
	return response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *PaymentHandler) GetTransactionByWarehouseId(c echo.Context) (err error) {
	ctx := c.Request().Context()
	warehouseId := c.Param("warehouseId")
	param, _ := paginate.GetParams(c)

	clamsData := utils.GetClamsJwt(c)
	if clamsData.UserRole != "admin" {
		fmt.Println("role is not admin")
		err = errors.New(http.StatusUnauthorized, "role is not admin")
		return
	}

	data, count, err := h.paymentUsecase.GetListTranscationByWarehouseId(ctx, cast.ToInt(warehouseId), param)
	if err != nil {
		return
	}

	resp := response.NewResponseSuccessPagination(float64(count), param, data)
	err = c.JSON(http.StatusOK, resp)
	return
}

