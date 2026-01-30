package service

import (
	"net/http"

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
func (s *OrderService) CreateOrder(request *dto.CreateOrderRequest, userID uuid.UUID) (*dto.CreateOrderResponse, *echo.HTTPError) {
	// Create order with processing status
	order := &entity.Order{
		UserID:      userID,
		Status:      "processing",
		TotalAmount: 0,
	}

	// Call repository to create order with items
	createdOrder, err := s.repo.CreateOrderWithItems(order, request.Items)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to create order: "+err.Error())
	}

	response := &dto.CreateOrderResponse{
		Order:   *toOrderDTO(createdOrder),
		Message: "Order created successfully with processing status",
	}

	return response, nil
}

// toOrderDTO converts Order entity to OrderDTO
func toOrderDTO(order *entity.Order) *dto.OrderDTO {
	orderItems := make([]dto.OrderItemDTO, 0)
	for _, item := range order.OrderItems {
		orderItems = append(orderItems, dto.OrderItemDTO{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			CreatedAt: item.CreatedAt,
		})
	}

	return &dto.OrderDTO{
		ID:          order.ID,
		UserID:      order.UserID,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
		OrderItems:  orderItems,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}
}
