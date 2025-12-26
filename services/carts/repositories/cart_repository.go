package repositories

import (
	"github.com/TranXuanPhong25/ecom/services/carts/models"
)

type ICartRepository interface {
	GetCart(userID string) ([]models.CartItem, error)
	GetItemQuantity(userID string, productVariantID int, shopID string) (int, error)
	AddItemToCart(item models.CartItem) error
	UpdateItemQuantity(item models.CartItem) error
	DeleteItemInCart(userID string, itemIDs []int) error
	ClearCart(userID string) error
}

func InitDBConnection() {
	ConnectPostgresDB()
	err := DB.AutoMigrate(&models.CartItem{})
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
		return
	}
}
func NewCartRepository() ICartRepository {
	return &postgresRepository{
		DB: DB,
	}
}
