package controllers

import (
	"net/http"

	"github.com/TranXuanPhong25/ecom/services/chat/dtos"
	"github.com/TranXuanPhong25/ecom/services/chat/services"
	"github.com/TranXuanPhong25/ecom/services/chat/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type IChatController interface {
	// Conversation
	CreateConversation(c echo.Context) error
	GetConversation(c echo.Context) error
	ListConversations(c echo.Context) error
	UpdateConversationStatus(c echo.Context) error
	DeleteConversation(c echo.Context) error

	// Message
	SendMessage(c echo.Context) error
	GetMessage(c echo.Context) error
	ListMessages(c echo.Context) error
	DeleteMessage(c echo.Context) error

	// LastRead
	UpdateLastRead(c echo.Context) error
	GetLastRead(c echo.Context) error
	ListLastReads(c echo.Context) error
}

type chatController struct {
	chatService services.IChatService
}

func NewChatController(chatService services.IChatService) IChatController {
	return &chatController{chatService: chatService}
}

// ---- Conversations ----

func (ctrl *chatController) CreateConversation(c echo.Context) error {
	req := new(dtos.CreateConversationPayload)
	if err := utils.ValidateRequestStructure(c, req); err != nil {
		return err
	}
	conv, err := ctrl.chatService.CreateConversation(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, conv)
}

func (ctrl *chatController) GetConversation(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid conversation id"})
	}
	conv, err := ctrl.chatService.GetConversation(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, conv)
}

func (ctrl *chatController) ListConversations(c echo.Context) error {
	limit := utils.ParseIntQuery(c, "limit", 20)
	offset := utils.ParseIntQuery(c, "offset", 0)
	convs, err := ctrl.chatService.ListConversations(limit, offset)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, convs)
}

func (ctrl *chatController) UpdateConversationStatus(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid conversation id"})
	}
	req := new(dtos.UpdateConversationStatusPayload)
	if err := utils.ValidateRequestStructure(c, req); err != nil {
		return err
	}
	if err := ctrl.chatService.UpdateConversationStatus(id, req); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "conversation status updated"})
}

func (ctrl *chatController) DeleteConversation(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid conversation id"})
	}
	if err := ctrl.chatService.DeleteConversation(id); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

// ---- Messages ----

func (ctrl *chatController) SendMessage(c echo.Context) error {
	convID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid conversation id"})
	}
	req := new(dtos.SendMessagePayload)
	if err := utils.ValidateRequestStructure(c, req); err != nil {
		return err
	}
	msg, err := ctrl.chatService.SendMessage(convID, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, msg)
}

func (ctrl *chatController) GetMessage(c echo.Context) error {
	msgID, err := uuid.Parse(c.Param("msgId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid message id"})
	}
	msg, err := ctrl.chatService.GetMessage(msgID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, msg)
}

func (ctrl *chatController) ListMessages(c echo.Context) error {
	convID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid conversation id"})
	}
	limit := utils.ParseIntQuery(c, "limit", 50)
	offset := utils.ParseIntQuery(c, "offset", 0)
	msgs, err := ctrl.chatService.ListMessages(convID, limit, offset)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, msgs)
}

func (ctrl *chatController) DeleteMessage(c echo.Context) error {
	msgID, err := uuid.Parse(c.Param("msgId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid message id"})
	}
	if err := ctrl.chatService.DeleteMessage(msgID); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

// ---- LastRead ----

func (ctrl *chatController) UpdateLastRead(c echo.Context) error {
	convID := c.Param("id")
	req := new(dtos.UpdateLastReadPayload)
	if err := utils.ValidateRequestStructure(c, req); err != nil {
		return err
	}
	if err := ctrl.chatService.UpdateLastRead(convID, req); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "last read updated"})
}

func (ctrl *chatController) GetLastRead(c echo.Context) error {
	convID := c.Param("id")
	participantID := c.QueryParam("participant_id")
	if participantID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "participant_id is required"})
	}
	lr, err := ctrl.chatService.GetLastRead(participantID, convID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, lr)
}

func (ctrl *chatController) ListLastReads(c echo.Context) error {
	convID := c.Param("id")
	lrs, err := ctrl.chatService.ListLastReads(convID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, lrs)
}
