package response

import (
	"math"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type responsePagination struct {
	BaseResponse
	Pagination paginate.ItemPages `json:"pagination"`
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

func NewResponseSuccessPagination(totalItems float64, params paginate.Pagination, data interface{}) responsePagination {
	var totalPage float64 = 1
	if params.Limit != 0 && params.Page != 0 {
		res := totalItems / float64(params.Limit)
		totalPage = math.Ceil(res)
	}

	resp := responsePagination{
		BaseResponse: BaseResponse{
			Status:  true,
			Message: "Success",
			Data:    data,
		},
		Pagination:   paginate.ItemPages{TotalData: int64(totalItems), TotalPage: int64(totalPage), Page: int64(params.Page), Limit: int64(params.Limit)},
	}
	return resp
}
