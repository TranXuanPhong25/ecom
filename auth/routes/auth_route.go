package routes

import (
	"github.com/TranXuanPhong25/ecom/auth/controllers"
	"github.com/labstack/echo/v4"
)

func AuthRoute(e *echo.Echo) {
	e.POST("/api/auth/login", controllers.LoginWithEmailAndPassword)
	e.POST("/api/auth/register", controllers.RegisterWithEmailAndPassword)
}
