package router

import (
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/handler"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/handler/middleware"
	// "github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
)

type Router struct {
	PingHandler      *handler.PingHandler
	UserHandler      *handler.UserHandler
	WarehouseHandler *handler.WarehouseHandler
	PaymentHandler   *handler.PaymentHandler
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

	if r.PaymentHandler == nil {
		panic("payment handler is nil")
	}
}

// func (r *Router) SetupScheduler(sch *gocron.Scheduler) {
// 	sch.Every(1).Hour().SingletonMode().Do(h.cashOutHandler.DisbursementProcess)
// 	sch.Every(1).Day().At("06:00").Do(h.cashinHandler.BatchingCashInToCashOut)
// }

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
	e.POST("/payment/va/callback", r.PaymentHandler.VaCallback)

	sc := e.Group("")
	sc.Use(middleware.JwtMiddleware())
	sc.GET("/user/profile", r.UserHandler.GetProfile)
	sc.PUT("/user/profile/username", r.UserHandler.UpdateUsername)
	sc.PUT("/user/profile/photo", r.UserHandler.UpdatePhotoProfile)
	sc.POST("/user/upload/photo", r.UserHandler.UploadPhoto)
	sc.GET("/user/avatar", r.UserHandler.GetAvatarList)
	sc.PUT("/user/profile/email", r.UserHandler.UpdateEmail)
	sc.GET("/user/info", r.UserHandler.GetUserInfo)
	sc.POST("/user/chatbot", r.UserHandler.ChatBot)
	
	sc.GET("/dasboard/user/list", r.UserHandler.GetUserList)
	sc.GET("/dasboard/warehouse/type", r.WarehouseHandler.GetListWarehouseType)
	sc.DELETE("/dasboard/user/:userId", r.UserHandler.DeleteUser)
	sc.GET("/dasboard/user/:userId", r.UserHandler.GetUserById)
	sc.GET("/dasboard/user/:userId/transaction", r.PaymentHandler.GetListTrxUserDasboar)
	sc.GET("/dasboard/user/transaction/:transactionId", r.PaymentHandler.GetTransactionDetailDasboardUser)
	sc.PUT("/dasboard/user/setting", r.UserHandler.UpdateAdminUser)

	sc.POST("/warehouse/detail", r.WarehouseHandler.CreateWarehouseDetail)
	sc.PUT("/warehouse/detail/:warehouseId", r.WarehouseHandler.UpdateWarehouseById)
	sc.GET("/warehouse/user/list", r.WarehouseHandler.GetWarehouseList)
	sc.GET("/warehouse/detail/:warehouseId", r.WarehouseHandler.GetWarehouseById)
	sc.DELETE("/warehouse/detail/:warehouseId", r.WarehouseHandler.DeleteWarehouseById)
	sc.POST("/warehouse/photo/upload", r.WarehouseHandler.UploadPhotoWarehouse)
	sc.GET("/warehouse/submitted", r.WarehouseHandler.MywarehouseSubmitted)
	sc.GET("/warehouse/rented", r.WarehouseHandler.MywarehouseRented)
	sc.GET("/warehouse/info/:warehouseId", r.WarehouseHandler.GetWarehouseInfo)
	sc.POST("/warehouse/favorit", r.WarehouseHandler.AddFavorit)
	sc.DELETE("/warehouse/favorit/:warehouseId", r.WarehouseHandler.DeleteFavorit)
	sc.GET("/warehouse/favorit", r.WarehouseHandler.GetListFavorit)
	sc.POST("warehouse/import-data", r.WarehouseHandler.ImportDataWarehouse)

	sc.POST("/payment/user/submission", r.PaymentHandler.SubmissionWarehouse)
	sc.GET("/payment/scheme", r.PaymentHandler.GetScheme)
	sc.GET("/payment/instalment/:transactionId", r.PaymentHandler.GetListInstalment)
	sc.GET("/payment/transaction/:transactionId", r.PaymentHandler.GetTransactionInfo)
	sc.GET("/payment/method", r.PaymentHandler.GetListPaymentMethod)
	sc.GET("/payment/va/bank", r.PaymentHandler.GetBankVa)
	sc.POST("/payment/checkout", r.PaymentHandler.PaymentCheckout)

	sc.GET("/dasboard/home/trx-history", r.PaymentHandler.GetHistoryInstalmentUser)
	sc.GET("/dasboard/list/trx-history", r.PaymentHandler.GetAllTransaction)
	sc.PUT("/dasboard/transaction/approval/:transactionId", r.PaymentHandler.TransactionApproved)
	sc.PUT("/dasboard/transaction/rejected/:transactionId", r.PaymentHandler.TransactionRejected)
	sc.GET("/dasboard/transaction/detail/:transactionId", r.PaymentHandler.GetTransactionListDetail)
	sc.GET("/dasboard/transaction/warehouse/:warehouseId", r.PaymentHandler.GetTransactionByWarehouseId)
	sc.GET("/dasboard/payment/total", r.PaymentHandler.GetTotalPaymentDasboard)
	sc.GET("/dasboard/payment/statistic", r.PaymentHandler.GetStatiscticPaymentDasboard)
	sc.GET("/dasboard/payment/reasone", r.PaymentHandler.GetReasone)

	return r
}
