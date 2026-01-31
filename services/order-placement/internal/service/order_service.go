package service

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/dto"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/entity"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/port"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// OrderService - Implementation of order business logic
type OrderService struct {
	repo port.OrderRepository
}

// NewOrderService creates a new OrderService
func NewOrderService(repo port.OrderRepository) port.OrderService {
	return &OrderService{
		repo: repo,
	}
}

// CreateOrder creates a new order with items
func (s *OrderService) CreateOrder(request *dto.CreateOrderRequest, userID string) (*dto.CreateOrderResponse, *echo.HTTPError) {
	// Generate order number
	orderNumber := generateOrderNumber()

	// Convert discount
	discount := entity.DiscountInfo{}
	if request.Discount != nil {
		discount = request.Discount
	}

	// Create order entity
	order := &entity.Order{
		OrderNumber:     orderNumber,
		UserID:          userID,
		ShopID:          request.ShopID,
		RecipientName:   request.RecipientName,
		RecipientPhone:  request.RecipientPhone,
		DeliveryAddress: request.DeliveryAddress,
		Status:          "CREATED",
		PaymentMethod:   request.PaymentMethod,
		PaymentStatus:   "UNPAID",
		ShippingMethod:  request.ShippingMethod,
		ShippingFee:     request.ShippingFee,
		Discount:        discount,
		CustomerNote:    request.CustomerNote,
	}

	// Call repository to create order with items
	createdOrder, err := s.repo.CreateOrderWithItems(order, request.Items)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to create order: "+err.Error())
	}

	response := &dto.CreateOrderResponse{
		Order:   *toOrderDTO(createdOrder),
		Message: "Order created successfully",
	}

	return response, nil
}

// generateOrderNumber generates a unique order number
func generateOrderNumber() string {
	now := time.Now()

	date := now.Format("020106") // ddmmyy

	// combine time (ns) + uuid, then hex, take 6 chars
	raw := fmt.Sprintf("%d%s", now.UnixNano(), uuid.New().String())
	hex6 := hex.EncodeToString([]byte(raw))[:6]
	return fmt.Sprintf("ORD-%s-%s", date, hex6)
}

// toOrderDTO converts Order entity to OrderDTO
func toOrderDTO(order *entity.Order) *dto.OrderDTO {
	orderItems := make([]dto.OrderItemDTO, 0, len(order.Items))

	for _, item := range order.Items {
		orderItems = append(orderItems, dto.OrderItemDTO{
			ProductID:     item.ProductID,
			ProductName:   item.ProductName,
			ProductSku:    item.ProductSku,
			ImageUrl:      item.ImageUrl,
			VariantID:     item.VariantID,
			VariantName:   item.VariantName,
			OriginalPrice: item.OriginalPrice,
			SalePrice:     item.SalePrice,
			Quantity:      item.Quantity,
			Subtotal:      item.SalePrice * int64(item.Quantity),
		})
	}

	discount := map[string]interface{}{}
	if order.Discount != nil {
		discount = order.Discount
	}

	return &dto.OrderDTO{
		ID:                order.ID,
		OrderNumber:       order.OrderNumber,
		UserID:            order.UserID,
		ShopID:            order.ShopID,
		RecipientName:     order.RecipientName,
		RecipientPhone:    order.RecipientPhone,
		DeliveryAddress:   order.DeliveryAddress,
		Status:            order.Status,
		PaymentMethod:     order.PaymentMethod,
		PaymentStatus:     order.PaymentStatus,
		PaidAt:            order.PaidAt,
		Subtotal:          order.Subtotal,
		ShippingFee:       order.ShippingFee,
		Discount:          discount,
		TotalAmount:       order.TotalAmount,
		ShippingMethod:    order.ShippingMethod,
		ShippingProvider:  order.ShippingProvider,
		TrackingNumber:    order.TrackingNumber,
		EstimatedDelivery: order.EstimatedDelivery,
		ActualDelivery:    order.ActualDelivery,
		CustomerNote:      order.CustomerNote,
		SellerNote:        order.SellerNote,
		CancelReason:      order.CancelReason,
		ConfirmedAt:       order.ConfirmedAt,
		ProcessingAt:      order.ProcessingAt,
		ShippedAt:         order.ShippedAt,
		DeliveredAt:       order.DeliveredAt,
		CompletedAt:       order.CompletedAt,
		CancelledAt:       order.CancelledAt,
		CreatedAt:         order.CreatedAt,
		UpdatedAt:         order.UpdatedAt,
		OrderItems:        orderItems,
	}
}
