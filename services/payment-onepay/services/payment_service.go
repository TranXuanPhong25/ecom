package services

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/TranXuanPhong25/ecom/services/payment-onepay/dtos"
	"github.com/TranXuanPhong25/ecom/services/payment-onepay/models"
	"github.com/TranXuanPhong25/ecom/services/payment-onepay/repositories"
	"github.com/google/uuid"
)

type IPaymentService interface {
	CreatePayment(req *dtos.CreatePaymentRequest, clientIP string) (*dtos.CreatePaymentResponse, error)
	HandleReturn(params url.Values) (*dtos.PaymentStatusResponse, error)
	HandleIPN(params url.Values) error
	GetPaymentStatus(txnRef string) (*dtos.PaymentStatusResponse, error)
	QueryTransaction(txnRef string) (*dtos.QueryDRResponse, error)
}

type paymentService struct {
	repo          repositories.IPaymentRepository
	onepayService IOnepayService
}

func NewPaymentService(
	repo repositories.IPaymentRepository,
	onepayService IOnepayService,
) IPaymentService {
	return &paymentService{
		repo:          repo,
		onepayService: onepayService,
	}
}

// CreatePayment creates a pending payment record and returns the OnePay redirect URL.
func (s *paymentService) CreatePayment(req *dtos.CreatePaymentRequest, clientIP string) (*dtos.CreatePaymentResponse, error) {
	txnRef := strings.ReplaceAll(uuid.New().String(), "-", "") // 32 chars, unique

	payment := &models.Payment{
		TxnRef:  txnRef,
		OrderID: req.OrderID,
		Amount:  req.Amount,
		Status:  "pending",
	}
	if err := s.repo.CreatePayment(payment); err != nil {
		return nil, err
	}

	paymentURL, err := s.onepayService.BuildPaymentURL(txnRef, req.OrderID, req.Amount, clientIP, req.Locale, req.CardList)
	if err != nil {
		return nil, err
	}

	return &dtos.CreatePaymentResponse{
		PaymentURL: paymentURL,
		TxnRef:     txnRef,
	}, nil
}

// HandleReturn processes the redirect from OnePay after user payment.
// Verifies hash integrity before updating the transaction record.
func (s *paymentService) HandleReturn(params url.Values) (*dtos.PaymentStatusResponse, error) {
	if !s.onepayService.VerifyCallback(params) {
		return nil, fmt.Errorf("hash verification failed – data may have been tampered")
	}

	txnRef := params.Get("vpc_MerchTxnRef")
	responseCode := params.Get("vpc_TxnResponseCode")
	transactionNo := params.Get("vpc_TransactionNo")
	message := params.Get("vpc_Message")
	cardType := params.Get("vpc_Card")
	payChannel := params.Get("vpc_PayChannel")

	status := "failed"
	if responseCode == "0" {
		status = "success"
	}

	if err := s.repo.UpdatePaymentStatus(txnRef, status, responseCode, transactionNo, message, cardType, payChannel); err != nil {
		return nil, err
	}

	payment, err := s.repo.GetByTxnRef(txnRef)
	if err != nil {
		return nil, err
	}

	return toStatusResponse(payment), nil
}

// HandleIPN processes the server-to-server Instant Payment Notification from OnePay.
func (s *paymentService) HandleIPN(params url.Values) error {
	if !s.onepayService.VerifyCallback(params) {
		return fmt.Errorf("hash verification failed – data may have been tampered")
	}

	txnRef := params.Get("vpc_MerchTxnRef")
	responseCode := params.Get("vpc_TxnResponseCode")
	transactionNo := params.Get("vpc_TransactionNo")
	message := params.Get("vpc_Message")
	cardType := params.Get("vpc_Card")
	payChannel := params.Get("vpc_PayChannel")

	status := "failed"
	if responseCode == "0" {
		status = "success"
	}

	return s.repo.UpdatePaymentStatus(txnRef, status, responseCode, transactionNo, message, cardType, payChannel)
}

// GetPaymentStatus returns the current DB record for a transaction.
func (s *paymentService) GetPaymentStatus(txnRef string) (*dtos.PaymentStatusResponse, error) {
	payment, err := s.repo.GetByTxnRef(txnRef)
	if err != nil {
		return nil, err
	}
	return toStatusResponse(payment), nil
}

// QueryTransaction calls the OnePay QueryDR API for authoritative status.
func (s *paymentService) QueryTransaction(txnRef string) (*dtos.QueryDRResponse, error) {
	return s.onepayService.QueryTransaction(txnRef)
}

func toStatusResponse(p *models.Payment) *dtos.PaymentStatusResponse {
	return &dtos.PaymentStatusResponse{
		TxnRef:        p.TxnRef,
		OrderID:       p.OrderID,
		Amount:        p.Amount,
		Status:        p.Status,
		ResponseCode:  p.ResponseCode,
		TransactionNo: p.TransactionNo,
		Message:       p.Message,
	}
}
