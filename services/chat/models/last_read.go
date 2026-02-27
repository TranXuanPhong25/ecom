package models

import "time"

type LastRead struct {
	ParticipantID  string    `gorm:"type:varchar(255);not null;primaryKey" json:"participant_id"`
	ConversationID string    `gorm:"type:uuid;not null;primaryKey" json:"conversation_id"`
	LastReadAt     time.Time `gorm:"not null" json:"last_read_at"`
}
