package chat

import (
	"time"

	"github.com/TranXuanPhong25/ecom/services/chat/internal/app/dto"
	"github.com/TranXuanPhong25/ecom/services/chat/internal/domain/entity"
	"github.com/TranXuanPhong25/ecom/services/chat/internal/domain/port"
	"github.com/google/uuid"
)

type chatService struct {
	convRepo     port.IConversationRepository
	msgRepo      port.IMessageRepository
	lastReadRepo port.ILastReadRepository
}

func NewChatService(
	convRepo port.IConversationRepository,
	msgRepo port.IMessageRepository,
	lastReadRepo port.ILastReadRepository,
) port.IChatService {
	return &chatService{
		convRepo:     convRepo,
		msgRepo:      msgRepo,
		lastReadRepo: lastReadRepo,
	}
}

func (s *chatService) CreateConversation(payload *dto.CreateConversationPayload) (*entity.Conversation, error) {
	conv := &entity.Conversation{
		Type:   payload.Type,
		Status: entity.ConversationStatusOpen,
	}
	if err := s.convRepo.Create(conv); err != nil {
		return nil, err
	}
	return conv, nil
}

func (s *chatService) GetConversation(id uuid.UUID) (*entity.Conversation, error) {
	return s.convRepo.GetByID(id)
}

func (s *chatService) ListConversations(limit, offset int) ([]entity.Conversation, error) {
	return s.convRepo.List(limit, offset)
}

func (s *chatService) UpdateConversationStatus(id uuid.UUID, payload *dto.UpdateConversationStatusPayload) error {
	return s.convRepo.UpdateStatus(id, payload.Status)
}

func (s *chatService) DeleteConversation(id uuid.UUID) error {
	return s.convRepo.Delete(id)
}

func (s *chatService) SendMessage(conversationID uuid.UUID, payload *dto.SendMessagePayload) (*entity.Message, error) {
	msg := &entity.Message{
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

func (s *chatService) GetMessage(id uuid.UUID) (*entity.Message, error) {
	return s.msgRepo.GetByID(id)
}

func (s *chatService) ListMessages(conversationID uuid.UUID, limit, offset int) ([]entity.Message, error) {
	return s.msgRepo.ListByConversation(conversationID, limit, offset)
}

func (s *chatService) DeleteMessage(id uuid.UUID) error {
	return s.msgRepo.SoftDelete(id)
}

func (s *chatService) UpdateLastRead(conversationID string, payload *dto.UpdateLastReadPayload) error {
	lr := &entity.LastRead{
		ParticipantID:  payload.ParticipantID,
		ConversationID: conversationID,
		LastReadAt:     time.Now(),
	}
	return s.lastReadRepo.Upsert(lr)
}

func (s *chatService) GetLastRead(participantID, conversationID string) (*entity.LastRead, error) {
	return s.lastReadRepo.Get(participantID, conversationID)
}

func (s *chatService) ListLastReads(conversationID string) ([]entity.LastRead, error) {
	return s.lastReadRepo.ListByConversation(conversationID)
}
