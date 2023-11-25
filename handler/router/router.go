package router

import (
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/handler"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/handler/middleware"
	"github.com/labstack/echo/v4"
)

type Router struct {
	PingHandler      *handler.PingHandler
	UserHandler      *handler.UserHandler
	WarehouseHandler *handler.WarehouseHandler
}

func (r *Router) Validate() {
	if r.PingHandler == nil {
		panic("ping handler is nil")
	}

	if r.UserHandler == nil {
		panic("user handler is nil")
	}

	if r.WarehouseHandler == nil {
		panic("warehouse handler is nil")
	}
}

func (r *Router) SetupRouter(e *echo.Echo) *Router {
	middleware.SetupMiddleware(e)
	e.GET("/ping", r.PingHandler.PingTestConnection)

	e.GET("/user/province", r.UserHandler.GetAllProvince)
	e.GET("/user/regency/:provinceId", r.UserHandler.GetRegencyByProvinceId)
	e.GET("/user/district/:regencyId", r.UserHandler.GetDistricByRegencyId)
	e.POST("/user/register-data", r.UserHandler.RegisterUserData)
	e.POST("/user/register", r.UserHandler.RegisterUser)
	e.POST("/user/resend-otp", r.UserHandler.ResendUserOTP)
	e.POST("/user/login", r.UserHandler.LoginUser)
	e.POST("/user/otp-verify", r.UserHandler.VerificationOtpUser)
	e.POST("/user/reset-password", r.UserHandler.ResetPassword)
	e.GET("/warehouse/detail/:warehouseId", r.WarehouseHandler.GetWarehouseById)
	e.GET("/warehouse/", r.WarehouseHandler.GetAllWarehouse)

	sc := e.Group("")
	sc.Use(middleware.JwtMiddleware())
	sc.GET("/user/profile", r.UserHandler.GetProfile)
	sc.PUT("/user/profile/username", r.UserHandler.UpdateUsername)
	sc.PUT("/user/profile/photo", r.UserHandler.UpdatePhotoProfile)
	sc.POST("/user/upload/photo", r.UserHandler.UploadPhoto)
	sc.GET("/user/avatar", r.UserHandler.GetAvatarList)
	sc.PUT("/user/profile/email", r.UserHandler.UpdateEmail)

	sc.POST("/warehouse/detail", r.WarehouseHandler.CreateWarehouseDetail)
	sc.PUT("/warehouse/detail/:warehouseId", r.WarehouseHandler.UpdateWarehouseById)
	sc.DELETE("/user/:userId", r.UserHandler.DeleteUser)

	return r
}
