package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/handler"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/handler/router"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/paymentdb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/regiondb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/userdb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/warehousedb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/http/chatgbt"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/http/core"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/warehouse"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()
	db := database.InitDB()

	local, _ := time.LoadLocation("Asia/Jakarta")
	sch := gocron.NewScheduler(local)
	sch.StartAsync()

	regionRepo := regiondb.NewRegionRepository(db)
	userRepo := userdb.NewUserRepository(db)
	warehouseRepo := warehousedb.NewWarehouseRepository(db)
	paymentRepo := paymentdb.NewPaymentRepository(db)

	coreWrapper := core.NewWrapper()
	chatWrapper := chatgbt.NewWrapper()

	userUsecase := user.NewUserUsecase(regionRepo, userRepo, coreWrapper, chatWrapper)
	warehouseUsecase := warehouse.NewWarehouseUsecase(warehouseRepo, userRepo, regionRepo, coreWrapper, paymentRepo)
	paymentUsecase := payment.NewPaymentUsecase(regionRepo, userRepo, coreWrapper, warehouseRepo, paymentRepo)

	pingHandler := handler.NewPingHandler()
	userHandler := handler.NewUserHandler(userUsecase)
	warehouseHandler := handler.NewWarehouseHandler(warehouseUsecase)
	paymentHandler := handler.NewPaymentHandler(paymentUsecase)

	route := router.Router{
		PingHandler:      pingHandler,
		UserHandler:      userHandler,
		WarehouseHandler: warehouseHandler,
		PaymentHandler:   paymentHandler,
	}

	e := echo.New()
	route.SetupRouter(e).Validate()

	err := e.Start(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
	if err != nil {
		panic("Start failed")
	}
}
