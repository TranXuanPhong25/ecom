package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"image-upload/routes"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.ImageUploadRoutes(e)
	e.Logger.Fatal(e.Start(":8201"))
}
