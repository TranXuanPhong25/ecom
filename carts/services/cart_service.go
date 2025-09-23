package services

import (
	"github.com/TranXuanPhong25/ecom/carts/dtos"
	"github.com/TranXuanPhong25/ecom/carts/models"
	"github.com/TranXuanPhong25/ecom/carts/repositories"
)

func GetCart(userID string) (*dtos.Cart, error) {

	items, err := repositories.GetCart(userID)
	if err != nil {
		return nil, err
	}
	cartItems := make([]dtos.CartItem, len(items))
	for i, item := range items {
		cartItems[i].Quantity = item.Quantity
		cartItems[i].ProductVariant = item.ProductVariantID
	}
	cart := dtos.Cart{
		Items: cartItems,
	}
	return &cart, nil
}

func AddItemToCart(userID string, item *dtos.AddItemRequest) error {
	cartItem := models.CartItem{
		UserID:           userID,
		ProductVariantID: item.ProductVariantID,
		ShopID:           item.ShopID,
		Quantity:         item.Quantity,
	}
	existingQuantity, err := repositories.GetItemQuantity(userID, item.ProductVariantID)
	if (existingQuantity > 0) || (err == nil) {
		item.Quantity += existingQuantity
		return repositories.UpdateItemQuantity(cartItem)
	}
	return repositories.AddItemToCart(cartItem)
}

func UpdateItemInCart(userID string, item *dtos.AddItemRequest) error {
	cartItem := models.CartItem{
		UserID:           userID,
		ProductVariantID: item.ProductVariantID,
		ShopID:           item.ShopID,
		Quantity:         item.Quantity,
	}
	return repositories.UpdateItemQuantity(cartItem)
}

func DeleteItemInCart(userID string, item *dtos.AddItemRequest) error {
	cartItem := models.CartItem{
		UserID:           userID,
		ProductVariantID: item.ProductVariantID,
		ShopID:           item.ShopID,
		Quantity:         item.Quantity,
	}
	return repositories.DeleteItemInCart(cartItem)
}

func ClearCart(userID string) error {
	return repositories.ClearCart(userID)
}
