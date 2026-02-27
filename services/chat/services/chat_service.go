package services

import (
	"time"

	"github.com/TranXuanPhong25/ecom/services/chat/dtos"
	"github.com/TranXuanPhong25/ecom/services/chat/models"
	"github.com/TranXuanPhong25/ecom/services/chat/repositories"
	"github.com/google/uuid"
)

type IChatService interface {
	// Conversation
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

type chatService struct {
	convRepo     repositories.IConversationRepository
	msgRepo      repositories.IMessageRepository
	lastReadRepo repositories.ILastReadRepository
}

func NewChatService(
	convRepo repositories.IConversationRepository,
	msgRepo repositories.IMessageRepository,
	lastReadRepo repositories.ILastReadRepository,
) IChatService {
	return &chatService{
		convRepo:     convRepo,
		msgRepo:      msgRepo,
		lastReadRepo: lastReadRepo,
	}
}

func (s *chatService) CreateConversation(payload *dtos.CreateConversationPayload) (*models.Conversation, error) {
	conv := &models.Conversation{
		Type:   payload.Type,
		Status: models.ConversationStatusOpen,
	}
	if err := s.convRepo.Create(conv); err != nil {
		return nil, err
	}
	return conv, nil
}

func (s *chatService) GetConversation(id uuid.UUID) (*models.Conversation, error) {
	return s.convRepo.GetByID(id)
}

func (s *chatService) ListConversations(limit, offset int) ([]models.Conversation, error) {
	return s.convRepo.List(limit, offset)
}

func (s *chatService) UpdateConversationStatus(id uuid.UUID, payload *dtos.UpdateConversationStatusPayload) error {
	return s.convRepo.UpdateStatus(id, payload.Status)
}

func (s *chatService) DeleteConversation(id uuid.UUID) error {
	return s.convRepo.Delete(id)
}

func (s *chatService) SendMessage(conversationID uuid.UUID, payload *dtos.SendMessagePayload) (*models.Message, error) {
	msg := &models.Message{
		ConversationID: conversationID,
		SenderID:       payload.SenderID,
		SenderType:     payload.SenderType,
		Content:        payload.Content,
		MessageType:    payload.MessageType,
		IsBotMessage:   payload.IsBotMessage,
	}
	if err := s.msgRepo.Create(msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (s *chatService) GetMessage(id uuid.UUID) (*models.Message, error) {
	return s.msgRepo.GetByID(id)
}

func (s *chatService) ListMessages(conversationID uuid.UUID, limit, offset int) ([]models.Message, error) {
	return s.msgRepo.ListByConversation(conversationID, limit, offset)
}

func (s *chatService) DeleteMessage(id uuid.UUID) error {
	return s.msgRepo.SoftDelete(id)
}

func (s *chatService) UpdateLastRead(conversationID string, payload *dtos.UpdateLastReadPayload) error {
	lr := &models.LastRead{
		ParticipantID:  payload.ParticipantID,
		ConversationID: conversationID,
		LastReadAt:     time.Now(),
	}
	return s.lastReadRepo.Upsert(lr)
}

func (s *chatService) GetLastRead(participantID, conversationID string) (*models.LastRead, error) {
	return s.lastReadRepo.Get(participantID, conversationID)
}

func (s *chatService) ListLastReads(conversationID string) ([]models.LastRead, error) {
	return s.lastReadRepo.ListByConversation(conversationID)
}
