package dtos

import (
	"time"

	"github.com/google/uuid"
)

type NotificationResponse struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"userID"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	IsRead    bool      `json:"isRead"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateNotificationPayload struct {
	UserID  uuid.UUID `json:"userID" validate:"required"`
	Title   string    `json:"title" validate:"required,max=255"`
	Message string    `json:"message" validate:"required"`
	Type    string    `json:"type" validate:"required,oneof=order promotion system"`
}
