package routes

import (
	"github.com/TranXuanPhong25/ecom/services/shops/controllers"
	"github.com/labstack/echo/v4"
)

func ShopsRoute(e *echo.Echo) {
	e.GET("/api/shops/owners/:ownerId", controllers.GetShopsByOwnerID)
	e.POST("/api/shops", controllers.CreateShop)
	e.GET("/api/shops", controllers.GetShopsByIDs)
}
