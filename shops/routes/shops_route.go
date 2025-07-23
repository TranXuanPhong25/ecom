package routes

import (
	"github.com/TranXuanPhong25/ecom/shops/controllers"
	"github.com/labstack/echo/v4"
)

func ShopsRoute(e *echo.Echo) {
	e.GET("/api/shops", controllers.GetShops)
	e.POST("/api/shops", controllers.CreateShop)
}
