package handler

import (
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/response"
	"github.com/labstack/echo/v4"
)

func PingTestConnection(c echo.Context) error {
	return response.NewSuccessResponse(c, "ping test")
}
