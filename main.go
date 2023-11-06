package main

import (
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/handler/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	router.SetupRouter(e)
	e.Start(":8080")
}
