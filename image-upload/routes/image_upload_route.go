package routes

import (
	"github.com/labstack/echo/v4"
	"image-upload/controllers"
)

func ImageUploadRoutes(e *echo.Echo) {
	e.GET("/api/image-upload", controllers.ImageUpload)
}
