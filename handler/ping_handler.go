package handler

import (
	"time"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/response"
	"github.com/labstack/echo/v4"
)

type PingHandler struct {}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) PingTestConnection(c echo.Context) error {
	data := map[string]interface{}{
		"time": time.Now().UTC().Format(time.ANSIC),
		"service": "status server is running",
	}
	return response.NewSuccessResponse(c, data)
}
