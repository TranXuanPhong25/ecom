package utils

import (
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

type productServiceClient struct {
	baseURL    string
	httpClient *http.Client
	cb         *gobreaker.CircuitBreaker
}

func NewProductService(baseURL string, client *http.Client) *productServiceClient {
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "ProductService",
		MaxRequests: 3,
		Timeout:     10 * time.Second,
	})

	return &productServiceClient{
		baseURL:    baseURL,
		httpClient: client,
		cb:         cb,
	}
}
