package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"product-review/database"
	"product-review/models"
	"product-review/routes"
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
	e.Logger.Fatal(e.Start(":8200"))
}
