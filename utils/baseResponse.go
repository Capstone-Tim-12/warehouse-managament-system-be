package utils

import (
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "Success",
		Data:    data,
	})
}

func NewErrorResponse(c echo.Context, err error) error {
	return c.JSON(errors.GetCodeError(err), BaseResponse{
		Status:  false,
		Message: err.Error(),
		Data:    nil,
	})
}
