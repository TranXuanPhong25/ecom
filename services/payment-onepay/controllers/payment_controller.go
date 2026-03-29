package controllers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/TranXuanPhong25/ecom/services/payment-onepay/dtos"
	"github.com/TranXuanPhong25/ecom/services/payment-onepay/services"
	"github.com/TranXuanPhong25/ecom/services/payment-onepay/utils"
	"github.com/labstack/echo/v4"
)

type IPaymentController interface {
	CreatePayment(c echo.Context) error
	HandleReturn(c echo.Context) error
	HandleIPN(c echo.Context) error
	GetPaymentStatus(c echo.Context) error
	QueryTransaction(c echo.Context) error
}

type PaymentController struct {
	service services.IPaymentService
}

func NewPaymentController(service services.IPaymentService) IPaymentController {
	return &PaymentController{service: service}
}

// CreatePayment - POST /api/payments/create
func (ctrl *PaymentController) CreatePayment(c echo.Context) error {
	userID := c.Request().Header.Get("X-User-Id")
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	req := new(dtos.CreatePaymentRequest)
	if err := utils.ValidateRequestStructure(c, req); err != nil {
		return err
	}

	clientIP := c.RealIP()
	if clientIP == "" {
		clientIP = "127.0.0.1" // fallback 
	}

	resp, err := ctrl.service.CreatePayment(req, clientIP)
	if err != nil {
		log.Println("Error creating payment:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

// HandleReturn - GET /api/payments/return
func (ctrl *PaymentController) HandleReturn(c echo.Context) error {
	params := make(url.Values)
	for k, v := range c.QueryParams() {
		params[k] = v
	}

	result, err := ctrl.service.HandleReturn(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

// HandleIPN - GET/POST /api/payments/ipn
func (ctrl *PaymentController) HandleIPN(c echo.Context) error {
	params := make(url.Values)

	if c.Request().Method == http.MethodPost {
		if err := c.Request().ParseForm(); err != nil {
			return c.String(http.StatusBadRequest, "responsecode=99&desc=invalid-request")
		}
		for k, v := range c.Request().Form {
			params[k] = v
		}
	} else {
		for k, v := range c.QueryParams() {
			params[k] = v
		}
	}

	if err := ctrl.service.HandleIPN(params); err != nil {
		return c.String(http.StatusOK, "responsecode=99&desc=failed")
	}
	return c.String(http.StatusOK, "responsecode=1&desc=confirm-success")
}

// GetPaymentStatus - GET /api/payments/:txn_ref/status
func (ctrl *PaymentController) GetPaymentStatus(c echo.Context) error {
	txnRef := c.Param("txn_ref")
	if txnRef == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "txn_ref is required"})
	}

	result, err := ctrl.service.GetPaymentStatus(txnRef)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "payment not found"})
	}
	return c.JSON(http.StatusOK, result)
}

// QueryTransaction - GET /api/payments/:txn_ref/query
func (ctrl *PaymentController) QueryTransaction(c echo.Context) error {
	txnRef := c.Param("txn_ref")
	if txnRef == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "txn_ref is required"})
	}

	result, err := ctrl.service.QueryTransaction(txnRef)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
