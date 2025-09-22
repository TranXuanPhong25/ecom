package routes

import (
	"github.com/TranXuanPhong25/ecom/carts/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterCartRoutes(e *echo.Echo) {
	e.GET("/api/carts", controllers.GetCart)

	e.POST("/api/carts/item", controllers.AddItemToCart)
}
