package chat

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/TranXuanPhong25/ecom/services/chat/internal/app/dto"
	"github.com/TranXuanPhong25/ecom/services/chat/internal/infras/configs"
)

type Hub struct {
	convs       map[string]map[*Client]bool // convID -> []Client
	userClients map[string]map[*Client]bool // userID -> []Client

	// broadcast message
	broadcast chan *dto.Message

	// join, leave client channel receiver
	register   chan *Client
	unregister chan *Client

	mu sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		convs:       make(map[string]map[*Client]bool),
		userClients: make(map[string]map[*Client]bool),
		broadcast:   make(chan *dto.Message, configs.SendBufferSize),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
	}
}
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.AddClient(client)
		case client := <-h.unregister:
			h.removeClient(client)
		case msg := <-h.broadcast:
			h.fanOut(msg)
		}
	}
}
func (h *Hub) AddClient(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.convs[c.ConvID] == nil {
		h.convs[c.ConvID] = make(map[*Client]bool)
	}
	h.convs[c.ConvID][c] = true

	if h.userClients[c.UserID] == nil {
		h.userClients[c.UserID] = make(map[*Client]bool)
	}
	h.userClients[c.UserID][c] = true
	log.Printf("[Hub] client %s (user=%s) joined conv=%s | room_size=%d",
		c.ID, c.UserID, c.ConvID, len(h.convs[c.ConvID]))
}

func (h *Hub) removeClient(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if clients, ok := h.convs[c.ConvID]; ok {
		if _, exists := clients[c]; exists {
			// nếu tồn tại clients
			close(c.Send)
			delete(clients, c)
			if len(clients) == 0 {
				delete(h.convs, c.ConvID)
				log.Printf("[Hub] conv=%s room removed (empty)", c.ConvID)
			}
		}
	}
	if clients, ok := h.userClients[c.UserID]; ok {
		if _, exists := clients[c]; exists {
			delete(clients, c)
			if len(clients) == 0 {
				delete(h.userClients, c.ConvID)
			}
		}
	}

	log.Printf("[Hub] client %s (user=%s) left conv=%s", c.ID, c.UserID, c.ConvID)
}

func (h *Hub) fanOut(msg *dto.Message) {
	h.mu.RLock()
	clients, ok := h.convs[msg.ConvID]
	h.mu.RUnlock()

	if !ok {
		return
	}
	for client := range clients {
		if client.UserID == msg.SenderID {
			continue
		}
		select {
		case client.Send <- encodeMessage(msg):
		default:
			log.Printf("[Hub] dropping slow client %s (user=%s)", client.ID, client.UserID)

			go func(c *Client) {
				h.unregister <- c
			}(client)
		}
	}
}

func (h *Hub) CountOnlineClient(convID string) int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.convs[convID])
}

func (h *Hub) IsUserOnline(userID string) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.userClients[userID]) > 0
}

func encodeMessage(msg *dto.Message) []byte {
	encoded, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error while encoding message: %v", err)
		return nil
	}
	return encoded
}
