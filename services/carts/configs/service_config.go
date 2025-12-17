// configs/service_config.go
package configs

import (
	"time"
)

type ServiceConfig struct {
	ProductServiceURL string
	Timeout           time.Duration
	MaxRetries        int
	ShopsServiceAddr  string
}

func LoadServiceConfig() *ServiceConfig {
	productServiceURL := getEnv("PRODUCT_SERVICE_URL", "http://10.109.152.194:8080/api")
	shopsServiceAddr := getEnv("SHOPS_SERVICE_ADDR", "10.102.226.128:50050")
	return &ServiceConfig{
		ProductServiceURL: productServiceURL,
		Timeout:           10 * time.Second,
		MaxRetries:        3,
		ShopsServiceAddr:  shopsServiceAddr,
	}
}
