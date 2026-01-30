package postgres

import (
	"fmt"
	"log"

	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/config"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/dto"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/entity"
	"github.com/TranXuanPhong25/ecom/services/order-placement/internal/core/port"
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

// CreateOrderWithItems creates an order and order items in a transaction
func (r *OrderRepository) CreateOrderWithItems(order *entity.Order, items []dto.OrderItemInput) (*entity.Order, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Create order
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Create order items
	var totalAmount float64
	orderItems := make([]entity.OrderItem, 0)

	for _, item := range items {
		orderItem := entity.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     0, // TODO: Fetch from product service
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		totalAmount += orderItem.Price * float64(orderItem.Quantity)
		orderItems = append(orderItems, orderItem)
	}

	// Update total amount
	order.TotalAmount = totalAmount
	if err := tx.Save(order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Load order items into order
	order.OrderItems = orderItems

	return order, nil
}
