package controllers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthController struct {
	l *slog.Logger
}

func NewHealthController() *HealthController {
	return new(HealthController)
}

func (h *HealthController) WithLogger(logger *slog.Logger) *HealthController {
	h.l = logger
	return h
}

func (h *HealthController) HealthCheck(c echo.Context) error {
	h.l.Info("Health Check OK")
	c.Response().Header().Add("Access-Control-Allow-Origin", "*")
	return c.String(http.StatusOK, "OK")
}
