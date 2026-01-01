package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomBaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// EventBanner - Quản lý banner sự kiện (Black Friday, Flash Sale, v.v.)
type EventBanner struct {
	CustomBaseModel
	Title       string    `gorm:"type:varchar(200);not null"`
	Description string    `gorm:"type:text"`
	ImageURL    string    `gorm:"type:varchar(500);not null"`
	LinkURL     string    `gorm:"type:varchar(500)"`
	StartTime   time.Time `gorm:"type:timestamp;not null"`
	EndTime     time.Time `gorm:"type:timestamp;not null"`
	Priority    int       `gorm:"type:int;default:0"` // Độ ưu tiên hiển thị (số càng cao càng ưu tiên)
	IsActive    bool      `gorm:"type:boolean;default:true"`
	EventType   string    `gorm:"type:varchar(50);not null"`       // "black_friday", "flash_sale", "new_year", etc.
	Position    string    `gorm:"type:varchar(20);default:'main'"` // "main", "sidebar", "popup"
}

// PromoBar - Quản lý thanh thông báo khuyến mãi ở top
type PromoBar struct {
	CustomBaseModel
	Message         string    `gorm:"type:varchar(300);not null"`
	BackgroundColor string    `gorm:"type:varchar(20);default:'#ff0000'"` // Màu nền
	TextColor       string    `gorm:"type:varchar(20);default:'#ffffff'"` // Màu chữ
	LinkURL         string    `gorm:"type:varchar(500)"`
	StartTime       time.Time `gorm:"type:timestamp;not null"`
	EndTime         time.Time `gorm:"type:timestamp;not null"`
	IsActive        bool      `gorm:"type:boolean;default:true"`
	Priority        int       `gorm:"type:int;default:0"`
	IsCloseable     bool      `gorm:"type:boolean;default:true"` // Người dùng có thể đóng được không
}
