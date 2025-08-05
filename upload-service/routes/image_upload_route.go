package routes

import (
	"github.com/TranXuanPhong25/ecom/upload-service/controllers"
	"github.com/labstack/echo/v4"
)

func ImageUploadRoutes(e *echo.Echo) {
	e.POST("/api/upload/presigned-url/image", controllers.GeneratePresignedURLUploadImage)
}
