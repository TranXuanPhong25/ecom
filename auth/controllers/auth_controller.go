package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Lấy tất cả reviews
func LoginWithEmailAndPassword(c echo.Context) error {
	
	return c.JSON(http.StatusOK, "")
}

// Tạo review
func RegisterWithEmailAndPassword(c echo.Context) error {

	return c.JSON(http.StatusOK, "")

}
