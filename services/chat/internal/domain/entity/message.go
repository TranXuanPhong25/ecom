package entity

import (
	"time"

	"github.com/google/uuid"
)

type MessageType string

const (
	MessageTypeText        MessageType = "text"
	MessageTypeImage       MessageType = "image"
	MessageTypeFile        MessageType = "file"
	MessageTypeSystemEvent MessageType = "system_event"
)

type Message struct {
	ID             uuid.UUID   `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ConversationID uuid.UUID   `gorm:"type:uuid;not null;index" json:"conversationId"`
	SenderID       string      `gorm:"type:varchar(255);not null" json:"senderId"`
	SenderType     string      `gorm:"type:varchar(50);not null" json:"senderType"`
	Content        string      `gorm:"type:text;not null" json:"content"`
	MessageType    MessageType `gorm:"type:varchar(20);not null;default:'text'" json:"messageType"`
	CreatedAt      time.Time   `gorm:"autoCreateTime" json:"createdAt"`
}
