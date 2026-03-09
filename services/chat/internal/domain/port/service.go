package port

import (
	dtos "github.com/TranXuanPhong25/ecom/services/chat/internal/app/dto"
	models "github.com/TranXuanPhong25/ecom/services/chat/internal/domain/entity"
	"github.com/google/uuid"
)

type IChatService interface {
	// Conversations
	CreateConversation(payload *dtos.CreateConversationPayload) (*models.Conversation, error)
	GetConversation(id uuid.UUID) (*models.Conversation, error)
	ListConversations(limit, offset int) ([]models.Conversation, error)
	UpdateConversationStatus(id uuid.UUID, payload *dtos.UpdateConversationStatusPayload) error
	DeleteConversation(id uuid.UUID) error

	// Message
	SendMessage(conversationID uuid.UUID, payload *dtos.SendMessagePayload) (*models.Message, error)
	GetMessage(id uuid.UUID) (*models.Message, error)
	ListMessages(conversationID uuid.UUID, limit, offset int) ([]models.Message, error)
	DeleteMessage(id uuid.UUID) error

	// LastRead
	UpdateLastRead(conversationID string, payload *dtos.UpdateLastReadPayload) error
	GetLastRead(participantID, conversationID string) (*models.LastRead, error)
	ListLastReads(conversationID string) ([]models.LastRead, error)
}