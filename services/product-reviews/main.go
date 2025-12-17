package main

import (
	"github.com/TranXuanPhong25/ecom/services/product-review/database"
	"github.com/TranXuanPhong25/ecom/services/product-review/models"
	"github.com/TranXuanPhong25/ecom/services/product-review/routes"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database.ConnectDB()
	err := database.DB.AutoMigrate(&models.Review{})
	if err != nil {
		return
	} // Tự tạo bảng nếu chưa có

	routes.ReviewRoutes(e) // Add các route
	e.Logger.Fatal(e.Start(":8080"))
}
