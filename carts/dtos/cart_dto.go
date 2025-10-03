package dtos

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
	ProductVariantID string `json:"productVariantID" validate:"required"`
	Quantity         int    `json:"quantity" validate:"required,min=1"`
	ShopID           string `json:"shopID" validate:"required"`
}
