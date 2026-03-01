package services

import (
	"github.com/TranXuanPhong25/ecom/services/notifications/dtos"
	"github.com/TranXuanPhong25/ecom/services/notifications/models"
	"github.com/TranXuanPhong25/ecom/services/notifications/repositories"
)

type INotificationService interface {
	GetMyNotifications(userID string) ([]dtos.NotificationResponse, error)
	GetUnreadCount(userID string) (int, error)
	CreateNotification(payload *dtos.CreateNotificationPayload) (*dtos.NotificationResponse, error)
	MarkAsRead(id string, userID string) error
	MarkAllAsRead(userID string) error
	DeleteNotification(id string, userID string) error
}

type NotificationService struct {
	repo repositories.INotificationRepository
}

func NewNotificationService(repo repositories.INotificationRepository) INotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) GetMyNotifications(userID string) ([]dtos.NotificationResponse, error) {
	notifications, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	result := make([]dtos.NotificationResponse, len(notifications))
	for i, n := range notifications {
		result[i] = toResponse(n)
	}
	return result, nil
}

func (s *NotificationService) GetUnreadCount(userID string) (int, error) {
	return s.repo.GetUnreadCount(userID)
}

func (s *NotificationService) CreateNotification(payload *dtos.CreateNotificationPayload) (*dtos.NotificationResponse, error) {
	n := models.Notification{
		UserID:  payload.UserID,
		Title:   payload.Title,
		Message: payload.Message,
		Type:    payload.Type,
	}
	created, err := s.repo.Create(n)
	if err != nil {
		return nil, err
	}
	resp := toResponse(*created)
	return &resp, nil
}

func (s *NotificationService) MarkAsRead(id string, userID string) error {
	return s.repo.MarkAsRead(id, userID)
}

func (s *NotificationService) MarkAllAsRead(userID string) error {
	return s.repo.MarkAllAsRead(userID)
}

func (s *NotificationService) DeleteNotification(id string, userID string) error {
	return s.repo.Delete(id, userID)
}

func toResponse(n models.Notification) dtos.NotificationResponse {
	return dtos.NotificationResponse{
		ID:        n.ID,
		UserID:    n.UserID,
		Title:     n.Title,
		Message:   n.Message,
		Type:      n.Type,
		IsRead:    n.IsRead,
		CreatedAt: n.CreatedAt,
	}
}
