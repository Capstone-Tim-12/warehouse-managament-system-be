package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RecoverMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				// Terjadi panic, tangani dan kirim respons dengan pesan error
				c.String(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		return next(c)
	}
}
