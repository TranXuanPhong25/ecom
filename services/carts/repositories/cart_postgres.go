package repositories

import (
	"fmt"
	"net/http"

	"github.com/TranXuanPhong25/ecom/services/carts/configs"
	"github.com/TranXuanPhong25/ecom/services/carts/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type postgresRepository struct {
	DB *gorm.DB
}

func (r *postgresRepository) GetCart(userID string) ([]models.CartItem, error) {
	cartItems := &[]models.CartItem{}
	tx := r.DB.Where("user_id = ?", userID).Find(cartItems)

	if tx.Error != nil {
		// Các lỗi khác (database connection, SQL syntax, etc.)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, tx.Error.Error())
	}

	return *cartItems, nil
}

func (r *postgresRepository) GetItemQuantity(userID string, productVariantID int, shopID string) (int, error) {
	cartItem := &[]models.CartItem{}
	tx := r.DB.Where("user_id = ? AND product_variant_id = ? AND shop_id = ?",
		userID, productVariantID, shopID).First(cartItem)
	if tx.Error != nil {
		// Các lỗi khác (database connection, SQL syntax, etc.)
		return 0, echo.NewHTTPError(http.StatusInternalServerError, tx.Error.Error())
	}

	if len(*cartItem) == 0 {
		return 0, nil
	}

	return (*cartItem)[0].Quantity, nil
}

func (r *postgresRepository) AddItemToCart(item models.CartItem) error {
	tx := r.DB.Create(&item)
	if tx.Error != nil {
		log.Error("Failed to add item to cart:", tx.Error)
		return echo.NewHTTPError(http.StatusInternalServerError, tx.Error.Error())
	}
	return nil
}

func (r *postgresRepository) UpdateItemQuantity(item models.CartItem) error {
	r.DB.Model(&models.CartItem{}).
		Where("user_id = ? AND product_variant_id = ? AND shop_id = ?",
			item.UserID, item.ProductVariantID, item.ShopID).
		Update("quantity", item.Quantity)

	return nil
}

func (r *postgresRepository) DeleteItemInCart(userID string, itemIDs []int) error {
	tx := r.DB.Where("user_id = ? AND product_variant_id IN ?", userID, itemIDs).Delete(&models.CartItem{})
	if tx.Error != nil {
		log.Error("Failed to delete item from cart:", tx.Error)
		return echo.NewHTTPError(http.StatusInternalServerError, tx.Error.Error())
	}
	return nil
}

func (r *postgresRepository) ClearCart(userID string) error {
	tx := r.DB.Where("user_id = ?", userID).Delete(&models.CartItem{})
	if tx.Error != nil {
		log.Error("Failed to clear cart:", tx.Error)
		return echo.NewHTTPError(http.StatusInternalServerError, tx.Error.Error())
	}
	return nil
}
func (r *postgresRepository) GetTotalItemsInCart(userID string) (int, error) {
	var total int64

	if err := r.DB.
		Model(&models.CartItem{}).
		Where("user_id = ?", userID).
		Count(&total).Error; err != nil {
		log.Error("GetTotalItemsInCart failed:", err)
		return 0, err
	}

	return int(total), nil
}

func ConnectPostgresDB() {
	dbHost := configs.AppConfig.DBHost
	dbUser := configs.AppConfig.DBUser
	dbPassword := configs.AppConfig.DBPassword
	dbName := configs.AppConfig.DBName
	dbPort := configs.AppConfig.DBPort
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db
}
