package models

import "time"

type LastRead struct {
	ParticipantID  string    `gorm:"type:varchar(255);not null;primaryKey" json:"participantId"`
	ConversationID string    `gorm:"type:uuid;not null;primaryKey" json:"conversationId"`
	LastReadAt     time.Time `gorm:"not null" json:"lastReadAt"`
}
