package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/TranXuanPhong25/ecom/services/auth/models"
	"github.com/TranXuanPhong25/ecom/services/auth/services"
	"github.com/TranXuanPhong25/ecom/services/auth/validators"
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
	c.Response().Header().Set("Set-Cookie", fmt.Sprintf("access_token=%s; Path=/; HttpOnly; SameSite=Lax", response.Token))
	c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
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

func GetCurrentUser(c echo.Context) error {
	userId := c.Request().Header.Get("X-User-Id")
	if userId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No user id found",
		})
	}
	user, err := services.GetCurrentUser(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &models.LoginResponse{
		User: models.UserInfo{
			UserId: user.UserId,
			Email:  user.Email,
		},
	})

}

func Logout(c echo.Context) error {
	c.Response().Header().Set("Set-Cookie", "access_token=; Path=/; HttpOnly; SameSite=Lax; Max-Age=0")
	c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully logged out",
	})
}
