package routes

import (
	"github.com/TranXuanPhong25/ecom/services/auth/controllers"
)

func AuthRoute(e *echo.Echo) {
	e.POST("/api/auth/login", controllers.LoginWithEmailAndPassword)
	e.POST("/api/auth/register", controllers.RegisterWithEmailAndPassword)
	e.GET("/api/auth/me", controllers.GetCurrentUser)
	e.POST("/api/auth/logout", controllers.Logout)
}
