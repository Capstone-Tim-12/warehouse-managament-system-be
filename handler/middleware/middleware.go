package middleware

import (
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupMiddleware(server *echo.Echo) {
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.OPTIONS},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "token", echo.HeaderContentType, "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	server.HTTPErrorHandler = errorHandler
}

func errorHandler(err error, c echo.Context) {
	if c.Get("error-handled") != nil {
		return
	}

	c.Set("error-handled", true)

	code := http.StatusInternalServerError
	resp := response.ResponseError(code, "general error")

	responseCode := http.StatusInternalServerError
	if he, ok := err.(*errors.ApplicationError); ok {
		responseCode = he.ErrorCode
		resp.Message = he.Error()
	}

	request := c.Request()
	ctx := c.Request().Context()
	c.SetRequest(request.WithContext(ctx))

	c.JSON(responseCode, resp)
}