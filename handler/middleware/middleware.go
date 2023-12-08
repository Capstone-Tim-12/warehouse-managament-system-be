package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/constrans"
	customErr "github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/response"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupMiddleware(server *echo.Echo) {
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.DELETE, echo.PUT, echo.OPTIONS},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "token", echo.HeaderContentType, "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	server.HTTPErrorHandler = errorHandler
	server.Validator = &DataValidator{ValidatorData: validator.New()}
}

func errorHandler(err error, c echo.Context) {
	if c.Get("error-handled") != nil {
		return
	}

	c.Set("error-handled", true)

	code := http.StatusInternalServerError
	resp := response.ResponseError(code, "general error")

	responseCode := http.StatusInternalServerError
	if he, ok := err.(*customErr.ApplicationError); ok {
		responseCode = he.ErrorCode
		resp.Message = he.Error()
	}

	request := c.Request()
	ctx := c.Request().Context()
	c.SetRequest(request.WithContext(ctx))

	c.JSON(responseCode, resp)
}

func JwtMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, customErr.New(http.StatusUnauthorized, "signature not valid")
				}
				return []byte(constrans.JwtSecret), nil
			})

			if err != nil || !token.Valid {
				return customErr.New(http.StatusUnauthorized, "invalid token")
			}

			claims := token.Claims.(jwt.MapClaims)
			c.Set("user", claims)

			return next(c)
		}
	}
}

type DataValidator struct {
	ValidatorData *validator.Validate
}

func (cv *DataValidator) Validate(i interface{}) error {
	err := cv.ValidatorData.Struct(i)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		var message string
		for _, err := range err.(validator.ValidationErrors) {
			message = fmt.Sprintf("input %v has on the %v tag", strings.ToLower(err.Field()), err.ActualTag())
		}
		return errors.New(message)
	}
	return err
}