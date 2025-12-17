package routes

import (
	"github.com/TranXuanPhong25/ecom/services/upload-service/controllers"
)

func ImageUploadRoutes(e *echo.Echo) {
	e.POST("/api/upload/presigned-url/image", controllers.GeneratePresignedURLUploadImage)
}
