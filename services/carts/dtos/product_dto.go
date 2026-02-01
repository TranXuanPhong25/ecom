// dtos/product_dto.go
package dtos

type ProductVariant struct {
	ID            int      `json:"id"`
	OriginalPrice float64  `json:"originalPrice"`
	SalePrice     float64  `json:"salePrice"`
	Stock         int      `json:"stockQuantity"`
	Name          string   `json:"name"`
	Sku           string   `json:"sku,omitempty"`
	Images        []string `json:"images,omitempty"`
	CoverImage    string   `json:"coverImage,omitempty"`
}

type StockCheckRequest struct {
	ProductID string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

type GetProductVariantsResponse struct {
	Variants    []ProductVariant `json:"variants"`
	NotFoundIDs []int            `json:"notFoundIds,omitempty"`
}
