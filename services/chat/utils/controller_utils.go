package utils

import (
	"net/http"
	"strconv"

	"github.com/TranXuanPhong25/ecom/services/chat/validators"
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

func ParseIntQuery(c echo.Context, key string, defaultVal int) int {
	val := c.QueryParam(key)
	if val == "" {
		return defaultVal
	}
	n, err := strconv.Atoi(val)
	if err != nil || n < 0 {
		return defaultVal
	}
	return n
}
