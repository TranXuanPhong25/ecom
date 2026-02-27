package routes

import (
	"github.com/TranXuanPhong25/ecom/services/chat/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterChatRoutes(e *echo.Echo, ctrl controllers.IChatController) {
	// Conversations
	e.POST("/api/chats/conversations", ctrl.CreateConversation)
	e.GET("/api/chats/conversations", ctrl.ListConversations)
	e.GET("/api/chats/conversations/:id", ctrl.GetConversation)
	e.PATCH("/api/chats/conversations/:id", ctrl.UpdateConversationStatus)
	e.DELETE("/api/chats/conversations/:id", ctrl.DeleteConversation)

	// Messages
	e.POST("/api/chats/conversations/:id/messages", ctrl.SendMessage)
	e.GET("/api/chats/conversations/:id/messages", ctrl.ListMessages)
	e.GET("/api/chats/conversations/:id/messages/:msgId", ctrl.GetMessage)
	e.DELETE("/api/chats/conversations/:id/messages/:msgId", ctrl.DeleteMessage)

	// LastRead
	e.PUT("/api/chats/conversations/:id/read", ctrl.UpdateLastRead)
	e.GET("/api/chats/conversations/:id/read", ctrl.ListLastReads)
	e.GET("/api/chats/conversations/:id/read/me", ctrl.GetLastRead)
}
