package router

import (
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/handler"
	"github.com/labstack/echo/v4"
)

type Router struct {
	PingHandler *handler.PingHandler
	UserHandler *handler.UserHandler
}

func (r *Router) Validate() {
	if r.PingHandler == nil {
		panic("ping handler is nil")
	}

	if r.UserHandler == nil {
		panic("user handler is nil")
	}
}

func (r *Router) SetupRouter(e *echo.Echo) *Router {
	e.GET("/ping", r.PingHandler.PingTestConnection)

	e.GET("/user/province", r.UserHandler.GetAllProvince)
	e.GET("/user/regency/:provinceId", r.UserHandler.GetRegencyByProvinceId)
	e.GET("/user/district/:regencyId", r.UserHandler.GetDistricByRegencyId)
	e.POST("/user/register-data", r.UserHandler.RegisterUserData)
	e.POST("/user/resend-otp", r.UserHandler.ResendUserOTP)

	return r
}
