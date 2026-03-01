package repositories

import (
	"github.com/TranXuanPhong25/ecom/services/notifications/models"
)

type INotificationRepository interface {
	GetByUserID(userID string) ([]models.Notification, error)
	GetUnreadCount(userID string) (int, error)
	Create(n models.Notification) (*models.Notification, error)
	MarkAsRead(id string, userID string) error
	MarkAllAsRead(userID string) error
	Delete(id string, userID string) error
}

func InitDBConnection() {
	ConnectPostgresDB()
	if err := DB.AutoMigrate(&models.Notification{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}

func NewNotificationRepository() INotificationRepository {
	return &postgresRepository{DB: DB}
}
