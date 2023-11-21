package utils

import (
	"encoding/json"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetClamsJwt(c echo.Context) model.JwtCustomClaims {
	userJwt := c.Get("user")
	data := userJwt.(jwt.MapClaims)

	var datas = map[string]interface{}{
		"userId":   data["userId"],
		"userRole": data["role"],
	}
	dataByte, _ := json.Marshal(datas)
	var resp model.JwtCustomClaims
	json.Unmarshal(dataByte, &resp)

	return resp
}
