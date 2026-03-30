package chat

import (
	"encoding/json"
	"log"
	"time"

	"github.com/TranXuanPhong25/ecom/services/chat/internal/app/dto"
	"github.com/TranXuanPhong25/ecom/services/chat/internal/domain/entity"
	"github.com/TranXuanPhong25/ecom/services/chat/internal/infras/configs"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID     string
	UserID string
	ConvID uuid.UUID
	Send   chan []byte
	Hub    *Hub
	Conn   *websocket.Conn
}

func NewClient(hub *Hub, conn *websocket.Conn, userID string, convID uuid.UUID) *Client {
	return &Client{
		ID:     uuid.New().String(),
		UserID: userID,
		ConvID: convID,
		Send:   make(chan []byte, configs.SendBufferSize),
		Hub:    hub,
		Conn:   conn,
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(configs.MaxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(configs.PongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(configs.PongWait))
		return nil
	})
	for {
		_, rawMsg, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err,
				websocket.CloseGoingAway,
				websocket.CloseNormalClosure) {
				log.Printf("[Client] read error client=%s: %v", c.ID, err)
			}
			return
		}
		var incomingMsg dto.IncomingMessageWS
		if err = json.Unmarshal(rawMsg, &incomingMsg); err != nil {
			log.Printf("[Client] invalid message from client=%s: %v", c.ID, err)
			return
		}
		if incomingMsg.MessageType != "text" && incomingMsg.Content != "image" {
			c.Conn.WriteMessage(websocket.PongMessage, nil)
			continue
		}
		persistMsg := &dto.SendMessagePayload{
			SenderID:    c.UserID,
			SenderType:  "customer",
			Content:     incomingMsg.Content,
			MessageType: entity.MessageType(incomingMsg.MessageType),
		}
		savedMsg, err := c.Hub.chatService.SendMessage(c.ConvID, persistMsg)
		//feedback sent/ error
		if err != nil {
			log.Printf("[Client] error sending message from client=%s: %v", c.ID, err)
			c.Conn.WriteMessage(websocket.TextMessage, []byte("Messaging error"))
			return
		}
		msgSendToHub := &dto.Message{
			ConvID:      c.ConvID.String(),
			SenderID:    c.UserID,
			Content:     incomingMsg.Content,
			MessageType: incomingMsg.MessageType,
			CreatedAt:   savedMsg.CreatedAt.UnixMilli(),
		}
		c.Hub.broadcast <- msgSendToHub
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(configs.PingPeriod)

	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(configs.WriteWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			}
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write([]byte("1.\n"))
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(configs.WriteWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
