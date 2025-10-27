package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HealthCheck(c echo.Context) error {
	healthcheckStruct := struct {
		Status bool `json:"status"`
	}{
		Status: true,
	}
	return c.JSON(http.StatusOK, healthcheckStruct)
}
