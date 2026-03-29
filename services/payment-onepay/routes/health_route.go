package routes

import (
"github.com/TranXuanPhong25/ecom/services/payment-onepay/controllers"
"github.com/labstack/echo/v4"
)

func RegisterHealthRoute(e *echo.Echo) {
e.GET("/health", controllers.GetHealthStatus)
}
