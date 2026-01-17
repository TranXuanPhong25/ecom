package main

import (
	"github.com/TranXuanPhong25/ecom/services/upload-service/configs"
	"github.com/TranXuanPhong25/ecom/services/upload-service/routes"
	"github.com/TranXuanPhong25/ecom/services/upload-service/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configs.LoadEnv()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	services.InitMinIOClient()

	routes.ImageUploadRoutes(e)
	e.Logger.Fatal(e.Start(":" + configs.AppConfig.ServerPort))
}
