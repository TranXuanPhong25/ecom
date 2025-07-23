package services

import (
	"github.com/TranXuanPhong25/ecom/shops/dtos"
	"github.com/TranXuanPhong25/ecom/shops/models"
	"github.com/TranXuanPhong25/ecom/shops/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func CreateShop(request *models.CreateShopRequest) (models.Shop, echo.HTTPError) {
	shop := &models.Shop{
		Name:        request.Name,
		Description: request.Description,
		OwnerID:     request.OwnerID,
		Location:    request.Location,
	}
	tx := repositories.DB.Create(shop)
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "unique constraint") {
			return models.Shop{}, *echo.NewHTTPError(http.StatusBadRequest, "Shop with this owner already exists")
		}
		return models.Shop{}, *echo.NewHTTPError(http.StatusInternalServerError, "Error while creating shop")
	}
	return *shop, echo.HTTPError{}

}

func GetShop(ownerId string) (*dtos.ShopDTO, *echo.HTTPError) {
	shop := &models.Shop{}
	tx := repositories.DB.First(shop, "owner_id = ?", ownerId)
	if tx.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, tx.Error.Error())
	}
	return toShopDTO(shop), nil

}

func toShopDTO(shop *models.Shop) *dtos.ShopDTO {
	return &dtos.ShopDTO{
		ID:          shop.ID,
		Name:        shop.Name,
		Description: shop.Description,
		OwnerID:     shop.OwnerID,
		Location:    shop.Location,
		Rating:      shop.Rating,
	}
}
