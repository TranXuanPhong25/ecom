package repositories

import (
	"fmt"
	"net/http"

	"github.com/TranXuanPhong25/ecom/services/chat/internal/domain/entity"
	"github.com/TranXuanPhong25/ecom/services/chat/internal/domain/port"
	"github.com/TranXuanPhong25/ecom/services/chat/internal/infras/configs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgresDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configs.AppConfig.DBHost,
		configs.AppConfig.DBUser,
		configs.AppConfig.DBPassword,
		configs.AppConfig.DBName,
		configs.AppConfig.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db
}

func InitDBConnection() {
	ConnectPostgresDB()
	err := DB.AutoMigrate(
		&entity.Conversation{},
		&entity.Message{},
		&entity.LastRead{},
	)
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}

// --- Conversation ---

type conversationRepository struct{ DB *gorm.DB }

func NewConversationRepository() port.IConversationRepository {
	return &conversationRepository{DB: DB}
}

func (r *conversationRepository) Create(conv *entity.Conversation) error {
	if err := r.DB.Create(conv).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (r *conversationRepository) GetByID(id uuid.UUID) (*entity.Conversation, error) {
	var conv entity.Conversation
	if err := r.DB.First(&conv, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "conversation not found")
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return &conv, nil
}

func (r *conversationRepository) List(limit, offset int) ([]entity.Conversation, error) {
	var convs []entity.Conversation
	if err := r.DB.Order("created_at DESC").Limit(limit).Offset(offset).Find(&convs).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return convs, nil
}

func (r *conversationRepository) UpdateStatus(id uuid.UUID, status entity.ConversationStatus) error {
	tx := r.DB.Model(&entity.Conversation{}).Where("id = ?", id).Update("status", status)
	if tx.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, tx.Error.Error())
	}
	if tx.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "conversation not found")
	}
	return nil
}

func (r *conversationRepository) Delete(id uuid.UUID) error {
	tx := r.DB.Delete(&entity.Conversation{}, "id = ?", id)
	if tx.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, tx.Error.Error())
	}
	if tx.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "conversation not found")
	}
	return nil
}

// --- Message ---

type messageRepository struct{ DB *gorm.DB }

func NewMessageRepository() port.IMessageRepository {
	return &messageRepository{DB: DB}
}

func (r *messageRepository) Create(msg *entity.Message) error {
	if err := r.DB.Create(msg).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (r *messageRepository) GetByID(id uuid.UUID) (*entity.Message, error) {
	var msg entity.Message
	if err := r.DB.First(&msg, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "message not found")
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return &msg, nil
}

func (r *messageRepository) ListByConversation(conversationID uuid.UUID, limit, offset int) ([]entity.Message, error) {
	var msgs []entity.Message
	if err := r.DB.Where("conversation_id = ?", conversationID).
		Order("created_at ASC").
		Limit(limit).Offset(offset).
		Find(&msgs).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return msgs, nil
}

func (r *messageRepository) SoftDelete(id uuid.UUID) error {
	tx := r.DB.Delete(&entity.Message{}, "id = ?", id)
	if tx.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, tx.Error.Error())
	}
	if tx.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "message not found")
	}
	return nil
}

// --- LastRead ---

type lastReadRepository struct{ DB *gorm.DB }

func NewLastReadRepository() port.ILastReadRepository {
	return &lastReadRepository{DB: DB}
}

func (r *lastReadRepository) Upsert(lr *entity.LastRead) error {
	tx := r.DB.Save(lr)
	if tx.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, tx.Error.Error())
	}
	return nil
}

func (r *lastReadRepository) Get(participantID, conversationID string) (*entity.LastRead, error) {
	var lr entity.LastRead
	if err := r.DB.First(&lr, "participant_id = ? AND conversation_id = ?", participantID, conversationID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "last read not found")
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return &lr, nil
}

func (r *lastReadRepository) ListByConversation(conversationID string) ([]entity.LastRead, error) {
	var lrs []entity.LastRead
	if err := r.DB.Where("conversation_id = ?", conversationID).Find(&lrs).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return lrs, nil
}
