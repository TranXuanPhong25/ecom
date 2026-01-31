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

// CreateOrderWithItems creates an order with items stored as JSONB
func (r *OrderRepository) CreateOrderWithItems(order *entity.Order, items []dto.OrderItemInput) (*entity.Order, error) {
	// Convert input items to entity items
	orderItems := make(entity.OrderItems, 0, len(items))

	for _, item := range items {
		orderItem := entity.OrderItem{
			ProductID:     item.ProductID,
			ProductName:   item.ProductName,
			ProductSku:    item.ProductSku,
			ImageUrl:      item.ImageUrl,
			VariantID:     item.VariantID,
			VariantName:   item.VariantName,
			OriginalPrice: item.OriginalPrice,
			SalePrice:     item.SalePrice,
			Quantity:      item.Quantity,
		}
		orderItems = append(orderItems, orderItem)
	}

	// Set order items and calculate total
	order.Items = orderItems
	order.CalculateTotalAmount()

	// Create order with JSONB items
	if err := r.db.Create(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}
