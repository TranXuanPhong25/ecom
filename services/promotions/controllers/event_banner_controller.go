package controllers

import (
	"net/http"
	"strconv"

	"github.com/TranXuanPhong25/ecom/services/promotions/models"
	"github.com/TranXuanPhong25/ecom/services/promotions/services"
	"github.com/TranXuanPhong25/ecom/services/promotions/validators"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ============ EVENT BANNER CONTROLLERS ============

func CreateEventBanner(c echo.Context) error {
	req := new(models.CreateEventBannerRequest)

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

	res, err := services.CreateEventBanner(req)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusCreated, res)
}

func GetEventBannerByID(c echo.Context) error {
	id := c.Param("id")
	if _, parseErr := uuid.Parse(id); parseErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid banner ID",
		})
	}

	banner, err := services.GetEventBannerByID(id)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, banner)
}

func GetAllEventBanners(c echo.Context) error {
	activeOnly := c.QueryParam("active_only") == "true"

	banners, err := services.GetAllEventBanners(activeOnly)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, banners)
}

func GetEventBannersByType(c echo.Context) error {
	eventType := c.Param("type")
	activeOnly := c.QueryParam("active_only") == "true"

	banners, err := services.GetEventBannersByType(eventType, activeOnly)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, banners)
}

func GetEventBannersByPosition(c echo.Context) error {
	position := c.Param("position")
	activeOnly := c.QueryParam("active_only") == "true"

	banners, err := services.GetEventBannersByPosition(position, activeOnly)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, banners)
}

func UpdateEventBanner(c echo.Context) error {
	req := new(models.UpdateEventBannerRequest)

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

	res, err := services.UpdateEventBanner(req)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, res)
}

func DeleteEventBanner(c echo.Context) error {
	id := c.Param("id")
	if _, parseErr := uuid.Parse(id); parseErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid banner ID",
		})
	}

	err := services.DeleteEventBanner(id)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Event banner deleted successfully",
	})
}
