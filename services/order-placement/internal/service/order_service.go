package service

import (
	"net/http"

	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/entity"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/port"
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
func (s *OrderService) CreateOrder(request *entity.CreateOrderRequest) (*entity.CreateOrderResponse, *echo.HTTPError) {
	// Create order with processing status
	order := &entity.Order{
		UserID:      request.UserID,
		Status:      "processing",
		TotalAmount: 0,
	}

	// Call repository to create order with items
	createdOrder, err := s.repo.CreateOrderWithItems(order, request.Items)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to create order: "+err.Error())
	}

	response := &entity.CreateOrderResponse{
		Order:   *toOrderDTO(createdOrder),
		Message: "Order created successfully with processing status",
	}

	return response, nil
}

// toOrderDTO converts Order entity to OrderDTO
func toOrderDTO(order *entity.Order) *entity.OrderDTO {
	orderItems := make([]entity.OrderItemDTO, 0)
	for _, item := range order.OrderItems {
		orderItems = append(orderItems, entity.OrderItemDTO{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			CreatedAt: item.CreatedAt,
		})
	}

	return &entity.OrderDTO{
		ID:          order.ID,
		UserID:      order.UserID,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
		OrderItems:  orderItems,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}
}
