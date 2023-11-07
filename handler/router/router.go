package router

import (
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/handler"
	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
	e.GET("/ping", handler.PingTestConnection)
}
