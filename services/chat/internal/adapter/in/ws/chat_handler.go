package ws

import (
	"net/http"

	"github.com/TranXuanPhong25/ecom/services/chat/internal/app/chat"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WSChatHandler struct {
	hub *chat.Hub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWSChatHandler(hub *chat.Hub) *WSChatHandler {

	return &WSChatHandler{
		hub: hub,
	}
}
func (h *WSChatHandler) RegisterWSChatRoutes(e *echo.Echo) {
	e.GET("/ws/chats", h.ServeWS)
}

func (h *WSChatHandler) ServeWS(c echo.Context) error {
	convID := c.Request().URL.Query().Get("convID")
	if convID == "" {
		// return c.JSON(http.StatusBadRequest, map[string]string{
		// 	"error": "convID is required",
		// })
		convID = "c41d92b5-937f-43b6-951f-0db260106923"
	}
	userID := c.Request().Header.Get("X-User-Id")
	if userID == "" {
		// return c.JSON(http.StatusUnauthorized, "")
		userID = c.Request().URL.Query().Get("UID")
	}
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Error upgrading websocket connection: " + err.Error(),
		})
	}
	client := chat.NewClient(h.hub, conn, userID, convID)
	h.hub.AddClient(client)
	go client.ReadPump()
	go client.WritePump()
	return nil

}
