package router

import (
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils"
	"github.com/labstack/echo/v4"
)

func pingTestConnection(c echo.Context) error {
	return utils.NewSuccessResponse(c, "ping test")
}

func SetupRouter(e *echo.Echo) {
	e.GET("/ping", pingTestConnection)
}
