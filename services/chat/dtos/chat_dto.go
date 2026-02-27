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
	SenderID    string             `json:"sender_id" validate:"required"`
	SenderType  string             `json:"sender_type" validate:"required"`
	Content     string             `json:"content" validate:"required"`
	MessageType models.MessageType `json:"message_type" validate:"required,oneof=text image file system_event"`
	IsBotMessage bool             `json:"is_bot_message"`
}

type MessageResponse struct {
	ID             uuid.UUID          `json:"id"`
	ConversationID uuid.UUID          `json:"conversation_id"`
	SenderID       string             `json:"sender_id"`
	SenderType     string             `json:"sender_type"`
	Content        string             `json:"content"`
	MessageType    models.MessageType `json:"message_type"`
	IsBotMessage   bool               `json:"is_bot_message"`
	CreatedAt      time.Time          `json:"created_at"`
}

// LastRead DTOs
type UpdateLastReadPayload struct {
	ParticipantID string `json:"participant_id" validate:"required"`
}
