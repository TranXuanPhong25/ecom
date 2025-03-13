package routes

import (
	"github.com/labstack/echo/v4"
	"product-review/controllers"
)

func ReviewRoutes(e *echo.Echo) {
	e.GET("/reviews", controllers.GetReviews)
	e.POST("/reviews", controllers.CreateReview)
}
