package routes

import (
	"github.com/TranXuanPhong25/ecom/carts/controllers"
	"github.com/labstack/echo/v4"
)

// Define health check route
func RegisterHealthRoute(e *echo.Echo) {
	e.GET("/health", controllers.GetHealthStatus)
}
