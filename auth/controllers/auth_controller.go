package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/TranXuanPhong25/ecom/auth/models"
	"github.com/TranXuanPhong25/ecom/auth/services"
	"github.com/TranXuanPhong25/ecom/auth/validators"
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

	if err := validators.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	response, err := services.LoginWithEmailAndPassword(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	c.Response().Header().Set("Set-Cookie", fmt.Sprintf("access_token=%s; Path=/; HttpOnly; SameSite=Strict", response.Token))
	return c.JSON(http.StatusOK, response)
}

func RegisterWithEmailAndPassword(c echo.Context) error {
	req := new(models.RegisterRequest)

	// Bind JSON request body vào struct
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	if err := validators.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	err := services.RegisterWithEmailAndPassword(req.Email, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "exists") {
			return c.JSON(http.StatusConflict, map[string]string{
				"error": strings.SplitAfter(err.Error(), "desc = ")[1],
			})
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully registered",
	})
}
