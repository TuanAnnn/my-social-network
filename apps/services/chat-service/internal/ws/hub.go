package ws

import (
	"chat-service/internal/models"
	"encoding/json"
)

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan models.Message // Sửa lại đường dẫn model
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan models.Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			msgBytes, _ := json.Marshal(message)
			for client := range h.Clients {
				if client.ID == message.ReceiverID || client.ID == message.SenderID {
					select {
					case client.Send <- msgBytes:
					default:
						close(client.Send)
						delete(h.Clients, client)
					}
				}
			}
		}
	}
}
