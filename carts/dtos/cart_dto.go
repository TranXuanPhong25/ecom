package dtos

type CartItem struct {
	ProductVariant string `json:"productVariant"`
	Quantity       int    `json:"quantity"`
}

type Cart struct {
	Items []CartItem `json:"items"`
}

type AddItemRequest struct {
	ProductVariantID string `json:"productVariantID" validate:"required"`
	Quantity         int    `json:"quantity" validate:"required,min=1"`
	ShopID           string `json:"shopId" validate:"required"`
}
