package dtos

import (
	"time"

	"github.com/TranXuanPhong25/ecom/services/chat/models"
	"github.com/google/uuid"
)

// Conversation DTOs
type CreateConversationPayload struct {
	Type models.ConversationType `json:"type" validate:"required,oneof=customer_shop customer_bot customer_system"`
}

type UpdateConversationStatusPayload struct {
	Status models.ConversationStatus `json:"status" validate:"required,oneof=open resolved pending"`
}

// Message DTOs
type SendMessagePayload struct {
	SenderID     string             `json:"senderId" validate:"required"`
	SenderType   string             `json:"senderType" validate:"required"`
	Content      string             `json:"content" validate:"required"`
	MessageType  models.MessageType `json:"messageType" validate:"required,oneof=text image file system_event"`
	IsBotMessage bool               `json:"isBotMessage"`
}

type MessageResponse struct {
	ID             uuid.UUID          `json:"id"`
	ConversationID uuid.UUID          `json:"conversationId"`
	SenderID       string             `json:"senderId"`
	SenderType     string             `json:"senderType"`
	Content        string             `json:"content"`
	MessageType    models.MessageType `json:"messageType"`
	IsBotMessage   bool               `json:"isBotMessage"`
	CreatedAt      time.Time          `json:"createdAt"`
}

// LastRead DTOs
type UpdateLastReadPayload struct {
	ParticipantID string `json:"participantId" validate:"required"`
}
