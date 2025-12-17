package controllers

func GetHealthStatus(c echo.Context) error {
	return c.String(200, "Cart Service is running")
}
