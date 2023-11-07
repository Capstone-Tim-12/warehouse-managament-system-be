package main

import (
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/handler/router"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/env"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/labstack/echo/v4"
)

func init() {
	env.InitEnvironment()
	database.InitDB()
	database.InitMigrate()
}

func main() {
	e := echo.New()
	e.Use(errors.RecoverMiddleware)

	router.SetupRouter(e)

	e.Start(":8080")
}
