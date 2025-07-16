package routes

import (
	"github.com/TranXuanPhong25/ecom/auth/controllers"
	"github.com/labstack/echo/v4"
)

func AuthRoute(e *echo.Echo) {
	e.POST("/login", controllers.LoginWithEmailAndPassword)
	e.POST("/register", controllers.RegisterWithEmailAndPassword)
	e.GET("/me", controllers.GetCurrentUser)
}
