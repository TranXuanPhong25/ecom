package services

import (
	"github.com/TranXuanPhong25/ecom/carts/dtos"
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
		cartItems[i].ProductVariant = item.ProductID
	}
	cart := dtos.Cart{
		Items: cartItems,
	}
	return &cart, nil
}
