package repositories

import (
	"github.com/TranXuanPhong25/ecom/services/chat/models"
	"github.com/google/uuid"
)

type IConversationRepository interface {
	Create(conv *models.Conversation) error
	GetByID(id uuid.UUID) (*models.Conversation, error)
	List(limit, offset int) ([]models.Conversation, error)
	UpdateStatus(id uuid.UUID, status models.ConversationStatus) error
	Delete(id uuid.UUID) error
}

type IMessageRepository interface {
	Create(msg *models.Message) error
	GetByID(id uuid.UUID) (*models.Message, error)
	ListByConversation(conversationID uuid.UUID, limit, offset int) ([]models.Message, error)
	SoftDelete(id uuid.UUID) error
}

type ILastReadRepository interface {
	Upsert(lr *models.LastRead) error
	Get(participantID, conversationID string) (*models.LastRead, error)
	ListByConversation(conversationID string) ([]models.LastRead, error)
}
