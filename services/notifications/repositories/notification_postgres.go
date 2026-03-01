package repositories

import (
	"fmt"
	"net/http"

	"github.com/TranXuanPhong25/ecom/services/notifications/configs"
	"github.com/TranXuanPhong25/ecom/services/notifications/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type postgresRepository struct {
	DB *gorm.DB
}

func ConnectPostgresDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configs.AppConfig.DBHost,
		configs.AppConfig.DBUser,
		configs.AppConfig.DBPassword,
		configs.AppConfig.DBName,
		configs.AppConfig.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db
}

func (r *postgresRepository) GetByUserID(userID string) ([]models.Notification, error) {
	var notifications []models.Notification
	tx := r.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&notifications)
	if tx.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, tx.Error.Error())
	}
	return notifications, nil
}

func (r *postgresRepository) GetUnreadCount(userID string) (int, error) {
	var count int64
	if err := r.DB.Model(&models.Notification{}).
		Where("user_id = ? AND is_read = false", userID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *postgresRepository) Create(n models.Notification) (*models.Notification, error) {
	if err := r.DB.Create(&n).Error; err != nil {
		log.Error("Failed to create notification:", err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return &n, nil
}

func (r *postgresRepository) MarkAsRead(id string, userID string) error {
	tx := r.DB.Model(&models.Notification{}).
		Where("id = ? AND user_id = ?", id, userID).
		Update("is_read", true)
	if tx.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, tx.Error.Error())
	}
	if tx.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "notification not found")
	}
	return nil
}

func (r *postgresRepository) MarkAllAsRead(userID string) error {
	if err := r.DB.Model(&models.Notification{}).
		Where("user_id = ? AND is_read = false", userID).
		Update("is_read", true).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (r *postgresRepository) Delete(id string, userID string) error {
	tx := r.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Notification{})
	if tx.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, tx.Error.Error())
	}
	if tx.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "notification not found")
	}
	return nil
}
