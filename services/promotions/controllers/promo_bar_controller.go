package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TranXuanPhong25/ecom/services/promotions/models"
	"github.com/TranXuanPhong25/ecom/services/promotions/services"
	"github.com/TranXuanPhong25/ecom/services/promotions/validators"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ============ PROMO BAR CONTROLLERS ============

func CreatePromoBar(c echo.Context) error {
	req := new(models.CreatePromoBarRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": strconv.Itoa(http.StatusBadRequest),
			"detail": fmt.Sprintf("Invalid request format %v", err.Error()),
		})
	}

	if err := validators.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": strconv.Itoa(http.StatusBadRequest),
			"detail": err.Error(),
		})
	}

	res, err := services.CreatePromoBar(req)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusCreated, res)
}

func GetPromoBarByID(c echo.Context) error {
	id := c.Param("id")
	if _, parseErr := uuid.Parse(id); parseErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid promo bar ID",
		})
	}

	promoBar, err := services.GetPromoBarByID(id)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, promoBar)
}

func GetAllPromoBars(c echo.Context) error {
	activeOnly := c.QueryParam("active_only") == "true"

	promoBars, err := services.GetAllPromoBars(activeOnly)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, promoBars)
}

func UpdatePromoBar(c echo.Context) error {
	req := new(models.UpdatePromoBarRequest)

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

	res, err := services.UpdatePromoBar(req)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, res)
}

func DeletePromoBar(c echo.Context) error {
	id := c.Param("id")
	if _, parseErr := uuid.Parse(id); parseErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid promo bar ID",
		})
	}

	err := services.DeletePromoBar(id)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Promo bar deleted successfully",
	})
}

// ============ COMBINED CONTROLLERS ============

func GetActivePromotions(c echo.Context) error {
	promotions, err := services.GetActivePromotions()
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, promotions)
}
