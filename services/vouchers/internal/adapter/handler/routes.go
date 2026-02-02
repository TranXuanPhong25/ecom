package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RegisterRoutes registers all routes for the voucher service
func RegisterRoutes(e *echo.Echo, handler *VoucherHandler) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Health check
	e.GET("/health", handler.HealthCheck)

	// API routes
	api := e.Group("/api/vouchers")
	{
		api.POST("", handler.CreateVoucher)
		api.GET("", handler.ListVouchers)
		api.GET("/code/:code", handler.GetVoucherByCode)
		api.GET("/:id", handler.GetVoucherByID)
		api.PUT("/:id", handler.UpdateVoucher)
		api.DELETE("/:id", handler.DeleteVoucher)
	}
}
