package routes

import (
	"github.com/TranXuanPhong25/ecom/product-review/controllers"
	"github.com/labstack/echo/v4"
)

func ReviewRoutes(e *echo.Echo) {
	e.GET("/reviews", controllers.GetReviews)
	e.POST("/reviews", controllers.CreateReview)
}
