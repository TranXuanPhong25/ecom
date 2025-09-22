package controllers

import "github.com/labstack/echo/v4"

func GetHealthStatus(c echo.Context) error {
	return c.String(200, "Cart Service is running")
}
