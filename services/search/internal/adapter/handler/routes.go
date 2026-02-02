package handler

import (
	"fmt"

	"github.com/TranXuanPhong25/ecom/services/search-service/internal/core/port/in"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetupRoutes configures all HTTP routes for the application
func SetupRoutes(e *echo.Echo, searchService in.SearchServicePort) {
	// Create handler
	handler := NewSearchHandler(searchService)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		Skipper: func(c echo.Context) bool {
			return c.Path() == "/health"
		},
		LogMethod:  true,
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogError:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			reqID, _ := c.Get("request_id").(string)
			fmt.Printf(
				"[HTTP] %s %s %d %s err=%v reqId=%s\n",
				v.Method,
				v.URI,
				v.Status,
				v.Latency,
				v.Error,
				reqID,
			)
			return nil
		},
	}))

	e.Use(middleware.Recover())

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	// Search routes
	e.GET("/api/search", handler.SearchProducts)
}
