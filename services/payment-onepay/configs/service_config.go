package configs

import "time"

type OnepayConfig struct {
	MerchantID string
	AccessCode string
	HashKey    string
	GatewayURL string
	QueryURL   string
	ReturnURL  string
	IPNURL     string
	User       string
	Password   string
	Timeout    time.Duration
}

func LoadServiceConfig() *OnepayConfig {
	return &OnepayConfig{
		MerchantID: getEnv("ONEPAY_MERCHANT_ID", "TESTONEPAY"),
		AccessCode: getEnv("ONEPAY_ACCESS_CODE", "6BEB2546"),
		HashKey:    getEnv("ONEPAY_HASH_KEY", "6D0870CDE5F24F34F3915FB0045120DB"),
		GatewayURL: getEnv("ONEPAY_GATEWAY_URL", "https://mtf.onepay.vn/paygate/vpcpay.op"),
		QueryURL:   getEnv("ONEPAY_QUERY_URL", "https://mtf.onepay.vn/msp/api/v1/vpc/invoices/queries"),
		ReturnURL:  getEnv("ONEPAY_RETURN_URL", "http://localhost:8080/api/payments/return"),
		IPNURL:     getEnv("ONEPAY_IPN_URL", "http://localhost:8080/api/payments/ipn"),
		User:       getEnv("ONEPAY_USER", "op01"),
		Password:   getEnv("ONEPAY_PASSWORD", "op123456"),
		Timeout:    30 * time.Second,
	}
}
