package main

import (
	"github.com/TranXuanPhong25/ecom/upload-service/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.ImageUploadRoutes(e)
	e.Logger.Fatal(e.Start(":8204"))
}
