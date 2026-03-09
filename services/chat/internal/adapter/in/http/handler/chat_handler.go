package handler

import (
	"net/http"

	"github.com/TranXuanPhong25/ecom/services/chat/internal/adapter/in/http/utils"
	"github.com/TranXuanPhong25/ecom/services/chat/internal/app/dto"
	"github.com/TranXuanPhong25/ecom/services/chat/internal/domain/port"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	chatService port.IChatService
}

func NewChatHandler(chatService port.IChatService) *ChatHandler {

	return &ChatHandler{
		chatService: chatService,
	}
}
func (h *ChatHandler) RegisterChatRoutes(e *echo.Echo) {
	// Conversations
	e.POST("/api/chats/conversations", h.CreateConversation)
	e.GET("/api/chats/conversations", h.ListConversations)
	e.GET("/api/chats/conversations/:id", h.GetConversation)
	e.PATCH("/api/chats/conversations/:id", h.UpdateConversationStatus)
	e.DELETE("/api/chats/conversations/:id", h.DeleteConversation)

	// Messages
	e.POST("/api/chats/conversations/:id/messages", h.SendMessage)
	e.GET("/api/chats/conversations/:id/messages", h.ListMessages)
	e.GET("/api/chats/conversations/:id/messages/:msgId", h.GetMessage)
	e.DELETE("/api/chats/conversations/:id/messages/:msgId", h.DeleteMessage)

	// LastRead
	e.PUT("/api/chats/conversations/:id/read", h.UpdateLastRead)
	e.GET("/api/chats/conversations/:id/read", h.ListLastReads)
	e.GET("/api/chats/conversations/:id/read/me", h.GetLastRead)
}

// ---- Conversations ----

func (h *ChatHandler) CreateConversation(c echo.Context) error {
	req := new(dto.CreateConversationPayload)
	if err := utils.ValidateRequestStructure(c, req); err != nil {
		return err
	}
	conv, err := h.chatService.CreateConversation(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, conv)
}

func (h *ChatHandler) GetConversation(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid conversation id"})
	}
	conv, err := h.chatService.GetConversation(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, conv)
}

func (h *ChatHandler) ListConversations(c echo.Context) error {
	limit := utils.ParseIntQuery(c, "limit", 20)
	offset := utils.ParseIntQuery(c, "offset", 0)
	convs, err := h.chatService.ListConversations(limit, offset)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, convs)
}

func (h *ChatHandler) UpdateConversationStatus(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid conversation id"})
	}
	req := new(dto.UpdateConversationStatusPayload)
	if err := utils.ValidateRequestStructure(c, req); err != nil {
		return err
	}
	if err := h.chatService.UpdateConversationStatus(id, req); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "conversation status updated"})
}

func (h *ChatHandler) DeleteConversation(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid conversation id"})
	}
	if err := h.chatService.DeleteConversation(id); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

// ---- Messages ----

func (h *ChatHandler) SendMessage(c echo.Context) error {
	convID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid conversation id"})
	}
	req := new(dto.SendMessagePayload)
	if err := utils.ValidateRequestStructure(c, req); err != nil {
		return err
	}
	msg, err := h.chatService.SendMessage(convID, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, msg)
}

func (h *ChatHandler) GetMessage(c echo.Context) error {
	msgID, err := uuid.Parse(c.Param("msgId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid message id"})
	}
	msg, err := h.chatService.GetMessage(msgID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, msg)
}

func (h *ChatHandler) ListMessages(c echo.Context) error {
	convID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid conversation id"})
	}
	limit := utils.ParseIntQuery(c, "limit", 50)
	offset := utils.ParseIntQuery(c, "offset", 0)
	msgs, err := h.chatService.ListMessages(convID, limit, offset)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, msgs)
}

func (h *ChatHandler) DeleteMessage(c echo.Context) error {
	msgID, err := uuid.Parse(c.Param("msgId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid message id"})
	}
	if err := h.chatService.DeleteMessage(msgID); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

// ---- LastRead ----

func (h *ChatHandler) UpdateLastRead(c echo.Context) error {
	convID := c.Param("id")
	req := new(dto.UpdateLastReadPayload)
	if err := utils.ValidateRequestStructure(c, req); err != nil {
		return err
	}
	if err := h.chatService.UpdateLastRead(convID, req); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "last read updated"})
}

func (h *ChatHandler) GetLastRead(c echo.Context) error {
	convID := c.Param("id")
	participantID := c.QueryParam("participant_id")
	if participantID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "participant_id is required"})
	}
	lr, err := h.chatService.GetLastRead(participantID, convID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, lr)
}

func (h *ChatHandler) ListLastReads(c echo.Context) error {
	convID := c.Param("id")
	lrs, err := h.chatService.ListLastReads(convID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, lrs)
}
