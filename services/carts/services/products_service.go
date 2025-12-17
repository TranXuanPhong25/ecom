// services/product_service.go
package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/TranXuanPhong25/ecom/carts/configs"
	"github.com/TranXuanPhong25/ecom/carts/dtos"
)

type productServiceClient struct {
	baseURL    string
	httpClient *http.Client
}

type IProductService interface {
	GetProductVariantByIds(productVariantIDs []string) (*dtos.GetProductVariantsResponse, error)
	// CheckStock(ctx context.Context, productID string) (int, error)
}
type ProductService struct {
	client *productServiceClient
}

func NewProductService(config *configs.ServiceConfig) IProductService {
	client := &http.Client{
		Timeout: config.Timeout,
	}
	return &ProductService{
		client: &productServiceClient{
			baseURL:    config.ProductServiceURL,
			httpClient: client,
		},
	}
}

func (s *ProductService) GetProductVariantByIds(productVariantIDs []string) (*dtos.GetProductVariantsResponse, error) {
	url := fmt.Sprintf("%s/product-variants?ids=%s", s.client.baseURL, strings.Join(productVariantIDs, ","))
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var product dtos.GetProductVariantsResponse
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &product, nil
}
