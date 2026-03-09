package handler

import (
	"github.com/labstack/echo/v4"
)

func RegisterHealthRoute(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "healthy"})
	})
}
