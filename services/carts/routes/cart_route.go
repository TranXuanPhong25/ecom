package routes

import (
	"github.com/TranXuanPhong25/ecom/services/carts/controllers"
)

func RegisterCartRoutes(e *echo.Echo, controllers controllers.ICartController) {

	e.GET("/api/carts/mine", controllers.GetCart)

	e.POST("/api/carts/mine/items", controllers.AddItemToCart)
	e.PUT("/api/carts/mine/items", controllers.UpdateCartItem)
	e.DELETE("/api/carts/mine/items", controllers.DeleteItemInCart)
}
