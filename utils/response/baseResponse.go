package response

import (
	"net/http"

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

func ResponseError(code int, message string) BaseResponse {
	return BaseResponse{
		Status:  false,
		Message: message,
		Data:    struct{}{},
	}
}
