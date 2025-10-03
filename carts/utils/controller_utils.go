package utils

import (
	"net/http"
	"strconv"

	"github.com/TranXuanPhong25/ecom/carts/validators"
	"github.com/labstack/echo/v4"
)

func ValidateRequestStructure(c echo.Context, req any) error {
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": strconv.Itoa(http.StatusBadRequest),
			"detail": "Invalid request format",
		})
	}

	if err := validators.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": strconv.Itoa(http.StatusBadRequest),
			"detail": err.Error(),
		})
	}
	return nil
}
