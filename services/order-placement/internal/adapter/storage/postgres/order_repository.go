package postgres

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/config"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/dto"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/entity"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/port"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OrderRepository - PostgreSQL implementation of OrderRepository
type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new PostgreSQL order repository
func NewOrderRepository(db *gorm.DB) port.OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

// ConnectDB establishes database connection
func ConnectDB() *gorm.DB {
	dbHost := config.AppConfig.DBHost
	dbUser := config.AppConfig.DBUser
	dbPassword := config.AppConfig.DBPassword
	dbName := config.AppConfig.DBName
	dbPort := config.AppConfig.DBPort
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}

// CreateOrderWithItems creates an order with items stored as JSONB
func (r *OrderRepository) CreateOrderWithItems(order *entity.Order, items []dto.OrderItemInput) (*entity.Order, error) {
	// Convert input items to entity items
	var totalAmount float64
	orderItems := make(entity.OrderItems, 0, len(items))

	for _, item := range items {
		orderItem := entity.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     0, // TODO: Fetch from product service
		}

		totalAmount += orderItem.Price * float64(orderItem.Quantity)
		orderItems = append(orderItems, orderItem)
	}

	// Set order fields
	order.Items = orderItems
	order.TotalAmount = totalAmount

	// Start transaction
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create order with JSONB items
	if err := tx.Create(order).Error; err != nil {
		return nil, err
	}

	// Create outbox event
	if err := r.createOutboxEvent(tx, order); err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return order, nil
}

// createOutboxEvent creates an outbox event for the order
func (r *OrderRepository) createOutboxEvent(tx *gorm.DB, order *entity.Order) error {
	// Prepare event payload
	payload := map[string]interface{}{
		"orderId":     order.ID.String(),
		"userId":      order.UserID.String(),
		"totalAmount": order.TotalAmount,
		"status":      order.Status,
		"items":       order.Items,
		"createdAt":   order.CreatedAt,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Create outbox event
	outboxEvent := &entity.Outbox{
		ID:            uuid.New(),
		AggregateType: "Order",
		AggregateID:   order.ID.String(),
		Type:          "OrderCreated",
		Payload:       payloadJSON,
		Timestamp:     time.Now(),
		TracingSpanID: nil, // Optional: add tracing span ID if available
	}

	return tx.Create(outboxEvent).Error
}
