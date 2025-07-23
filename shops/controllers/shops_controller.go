package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateShop(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully registered",
	})
}

func GetShops(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
