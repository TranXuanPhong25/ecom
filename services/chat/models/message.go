package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageType string

const (
	MessageTypeText        MessageType = "text"
	MessageTypeImage       MessageType = "image"
	MessageTypeFile        MessageType = "file"
	MessageTypeSystemEvent MessageType = "system_event"
)

type Message struct {
	ID             uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ConversationID uuid.UUID      `gorm:"type:uuid;not null;index" json:"conversation_id"`
	SenderID       string         `gorm:"type:varchar(255);not null" json:"sender_id"`
	SenderType     string         `gorm:"type:varchar(50);not null" json:"sender_type"`
	Content        string         `gorm:"type:text;not null" json:"content"`
	MessageType    MessageType    `gorm:"type:varchar(20);not null;default:'text'" json:"message_type"`
	IsBotMessage   bool           `gorm:"not null;default:false" json:"is_bot_message"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
