package dtos

import "github.com/google/uuid"

type CartItem struct {
	ProductVariant ProductVariant `json:"productVariant"`
	Quantity       int            `json:"quantity"`
	ShopID         string         `json:"shopID,omitempty"`
}

type Cart struct {
	Items []CartItem `json:"items"`
	Shops []Shop     `json:"shops,omitempty"`
}

type CartItemPayload struct {
	ProductVariantID int       `json:"productVariantID" validate:"required"`
	Quantity         int       `json:"quantity" validate:"required,min=1"`
	ShopID           uuid.UUID `json:"shopID" validate:"required"`
}

type DeleteCartItemsPayload struct {
	Items []int `json:"ids" validate:"required,dive,required"`
}
