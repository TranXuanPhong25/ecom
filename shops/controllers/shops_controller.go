package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/TranXuanPhong25/ecom/shops/models"
	"github.com/TranXuanPhong25/ecom/shops/services"
	"github.com/TranXuanPhong25/ecom/shops/validators"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateShop(c echo.Context) error {
	req := new(models.CreateShopRequest)

	// Bind JSON request body vào struct
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

	res, err := services.CreateShop(req)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, res)
}

func GetShopsByOwnerID(c echo.Context) error {
	ownerId := c.QueryParam("owner_id")
	_, parseErr := uuid.Parse(ownerId)
	if parseErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid owner ID",
		})
	}
	shop, err := services.GetShopsByOwnerID(ownerId)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, shop)
}

func GetShopsByIDs(c echo.Context) error {
	idsParams := c.QueryParam("ids")
	if idsParams == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "IDs parameter is required",
		})
	}
	ids := []string{}
	for _, id := range strings.Split(idsParams, ",") {
		if _, err := uuid.Parse(id); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid ID format: " + id,
			})
		}
		ids = append(ids, id)
	}
	shop, err := services.GetShopsByIDs(ids)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, shop)
}
