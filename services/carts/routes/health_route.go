package routes

import (
	"github.com/TranXuanPhong25/ecom/services/carts/controllers"
)

// Define health check route
func RegisterHealthRoute(e *echo.Echo) {
	e.GET("/health", controllers.GetHealthStatus)
}
