package dtos

type CartItem struct {
	ProductVariant string `json:"productVariant"`
	Quantity       int    `json:"quantity"`
}

type Cart struct {
	Items []CartItem `json:"items"`
}
