package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/TranXuanPhong25/ecom/services/payment-onepay/configs"
	"github.com/TranXuanPhong25/ecom/services/payment-onepay/dtos"
	"github.com/TranXuanPhong25/ecom/services/payment-onepay/utils"
)

type IOnepayService interface {
	BuildPaymentURL(txnRef, orderInfo string, amount int64, clientIP, locale, cardList string) (string, error)
	QueryTransaction(txnRef string) (*dtos.QueryDRResponse, error)
	VerifyCallback(params url.Values) bool
}

type onepayService struct {
	config *configs.OnepayConfig
	client *http.Client
}

func NewOnepayService(config *configs.OnepayConfig) IOnepayService {
	return &onepayService{
		config: config,
		client: &http.Client{Timeout: config.Timeout},
	}
}

// BuildPaymentURL constructs the OnePay redirect URL with all required params and HMAC-SHA256 hash.
func (s *onepayService) BuildPaymentURL(txnRef, orderInfo string, amount int64, clientIP, locale, cardList string) (string, error) {
	if locale == "" {
		locale = "vn"
	}

	params := url.Values{}
	params.Set("vpc_Version", "2")
	params.Set("vpc_Currency", "VND")
	params.Set("vpc_Command", "pay")
	params.Set("vpc_AccessCode", s.config.AccessCode)
	params.Set("vpc_Merchant", s.config.MerchantID)
	params.Set("vpc_Locale", locale)
	params.Set("vpc_ReturnURL", s.config.ReturnURL)
	params.Set("vpc_MerchTxnRef", txnRef)
	params.Set("vpc_OrderInfo", orderInfo)
	// Amount is multiplied by 100 per OnePay spec (VND 25,000 → 2500000)
	params.Set("vpc_Amount", fmt.Sprintf("%d", amount*100))
	params.Set("vpc_TicketNo", clientIP)
	// Non-vpc_ params (not included in hash)
	params.Set("AgainLink", s.config.ReturnURL)
	params.Set("Title", "Shopiew Payment")

	if cardList != "" {
		params.Set("vpc_CardList", cardList)
	}

	hash := utils.BuildSecureHash(params, s.config.HashKey)
	params.Set("vpc_SecureHash", hash)

	return s.config.GatewayURL + "?" + params.Encode(), nil
}

// VerifyCallback checks the vpc_SecureHash returned from OnePay on Return URL or IPN.
func (s *onepayService) VerifyCallback(params url.Values) bool {
	receivedHash := params.Get("vpc_SecureHash")
	if receivedHash == "" {
		return false
	}
	return utils.VerifySecureHash(params, s.config.HashKey, receivedHash)
}

// QueryTransaction calls the OnePay QueryDR API to check transaction status.
func (s *onepayService) QueryTransaction(txnRef string) (*dtos.QueryDRResponse, error) {
	params := url.Values{}
	params.Set("vpc_Command", "queryDR")
	params.Set("vpc_Version", "2")
	params.Set("vpc_MerchTxnRef", txnRef)
	params.Set("vpc_Merchant", s.config.MerchantID)
	params.Set("vpc_AccessCode", s.config.AccessCode)
	params.Set("vpc_User", s.config.User)
	params.Set("vpc_Password", s.config.Password)

	hash := utils.BuildSecureHash(params, s.config.HashKey)
	params.Set("vpc_SecureHash", hash)

	ctx, cancel := context.WithTimeout(context.Background(), s.config.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.config.QueryURL, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("QueryDR request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Try JSON first, then fall back to query-string format
	var raw map[string]interface{}
	if jsonErr := json.Unmarshal(body, &raw); jsonErr == nil {
		drExists, _ := raw["vpc_DRExists"].(string)
		responseCode, _ := raw["vpc_TxnResponseCode"].(string)
		transactionNo, _ := raw["vpc_TransactionNo"].(string)
		message, _ := raw["vpc_Message"].(string)
		amountStr, _ := raw["vpc_Amount"].(string)
		return &dtos.QueryDRResponse{
			TxnExists:     drExists == "Y",
			ResponseCode:  responseCode,
			TransactionNo: transactionNo,
			Amount:        amountStr,
			Message:       message,
		}, nil
	}

	// fall back to query-string format
	parsed, _ := url.ParseQuery(string(body))
	return &dtos.QueryDRResponse{
		TxnExists:     parsed.Get("vpc_DRExists") == "Y",
		ResponseCode:  parsed.Get("vpc_TxnResponseCode"),
		TransactionNo: parsed.Get("vpc_TransactionNo"),
		Amount:        parsed.Get("vpc_Amount"),
		Message:       parsed.Get("vpc_Message"),
	}, nil
}
