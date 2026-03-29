package dtos

// CreatePaymentRequest - payload to initiate a payment session
type CreatePaymentRequest struct {
	OrderID  string `json:"order_id" validate:"required,max=40"`
	Amount   int64  `json:"amount" validate:"required,min=1000"` // in VND
	Locale   string `json:"locale" validate:"omitempty,oneof=vn en"`
	CardList string `json:"card_list"` // e.g. "DOMESTIC", "INTERNATIONAL", "QR" or bank BIN
}

// CreatePaymentResponse - redirect URL and reference id
type CreatePaymentResponse struct {
	PaymentURL string `json:"payment_url"`
	TxnRef     string `json:"txn_ref"`
}

// PaymentStatusResponse - current record of a payment transaction
type PaymentStatusResponse struct {
	TxnRef        string `json:"txn_ref"`
	OrderID       string `json:"order_id"`
	Amount        int64  `json:"amount"`
	Status        string `json:"status"` // pending | success | failed
	ResponseCode  string `json:"response_code"`
	TransactionNo string `json:"transaction_no"`
	Message       string `json:"message"`
}

// QueryDRResponse - result from OnePay QueryDR API
type QueryDRResponse struct {
	TxnExists     bool   `json:"txn_exists"`
	ResponseCode  string `json:"response_code"`
	TransactionNo string `json:"transaction_no"`
	Amount        string `json:"amount"`
	Message       string `json:"message"`
}
