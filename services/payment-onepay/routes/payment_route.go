package routes

import (
	"github.com/TranXuanPhong25/ecom/services/payment-onepay/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterPaymentRoutes(e *echo.Echo, ctrl controllers.IPaymentController) {
	e.POST("/api/payments/create", ctrl.CreatePayment)
	e.GET("/api/payments/return", ctrl.HandleReturn)
	e.GET("/api/payments/ipn", ctrl.HandleIPN)
	e.POST("/api/payments/ipn", ctrl.HandleIPN)
	e.GET("/api/payments/:txn_ref/status", ctrl.GetPaymentStatus)
	e.GET("/api/payments/:txn_ref/query", ctrl.QueryTransaction)
}
