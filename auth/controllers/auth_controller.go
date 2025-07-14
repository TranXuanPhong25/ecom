package controllers

import (
	"github.com/TranXuanPhong25/ecom/auth/models"
	"github.com/TranXuanPhong25/ecom/auth/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginWithEmailAndPassword(c echo.Context) error {
	req := new(models.LoginRequest)

	// Bind JSON request body vào struct
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	// Validate dữ liệu
	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Email and password are required",
		})
	}

	//if !services.IsValidEmailFormat(req.Email) {
	//	return c.JSON(http.StatusBadRequest, map[string]string{
	//		"error": "Invalid email format",
	//	})
	//}
	response, err := services.LoginWithEmailAndPassword(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response)
}

func RegisterWithEmailAndPassword(c echo.Context) error {
	req := new(models.LoginRequest)

	// Bind JSON request body vào struct
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	// Validate dữ liệu
	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Email and password are required",
		})
	}
	//if !services.IsValidEmailFormat(req.Email) {
	//	return c.JSON(http.StatusBadRequest, map[string]string{
	//		"error": "Invalid email format",
	//	})
	//}
	err := services.RegisterWithEmailAndPassword(req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully registered",
	})
}
