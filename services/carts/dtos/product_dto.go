// dtos/product_dto.go
package dtos

type ProductVariant struct {
	ID    int     `json:"id"`
	Price float64 `json:"price"`
	Stock int     `json:"stockQuantity"`
	Name  string  `json:"name"`
}

type StockCheckRequest struct {
	ProductID string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

type GetProductVariantsResponse struct {
	Variants    []ProductVariant `json:"variants"`
	NotFoundIDs []string         `json:"notFoundIds,omitempty"`
}
