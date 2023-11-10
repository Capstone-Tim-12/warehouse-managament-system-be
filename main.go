package main

import (
	"fmt"
	"os"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/handler"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/handler/router"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/regiondb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/userdb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()
	db := database.InitDB()

	regionRepo := regiondb.NewRegionRepository(db)
	userRepo := userdb.NewUserRepository(db)

	userUsecase := user.NewUserUsecase(regionRepo, userRepo)

	pingHandler := handler.NewPingHandler()
	userHandler := handler.NewUserHandler(userUsecase)

	route := router.Router{
		PingHandler: pingHandler,
		UserHandler: userHandler,
	}

	e := echo.New()
	route.SetupRouter(e).Validate()

	err := e.Start(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
	if err != nil {
		panic("Start failed")
	}
}
