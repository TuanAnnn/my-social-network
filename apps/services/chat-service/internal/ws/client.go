package ws

import (
	"chat-service/internal/auth"
	"chat-service/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Client struct {
	ID   uint
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
	DB   *gorm.DB // Inject DB vào Client để lưu tin nhắn
}

func (c *Client) readPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, messageBytes, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		var msgInput models.MessageInput
		if err := json.Unmarshal(messageBytes, &msgInput); err != nil {
			log.Println("Lỗi JSON:", err)
			continue
		}

		msgInput.SenderID = c.ID

		newMessage := models.Message{
			Content:    msgInput.Content,
			SenderID:   msgInput.SenderID,
			ReceiverID: msgInput.ReceiverID,
			CreatedAt:  time.Now(),
		}

		// Dùng DB của Client để lưu
		if err := c.DB.Create(&newMessage).Error; err != nil {
			log.Println("Lỗi lưu DB:", err)
			continue
		}

		c.Hub.Broadcast <- newMessage
	}
}

func (c *Client) writePump() {
	defer c.Conn.Close()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

// ServeWs nhận DB từ bên ngoài vào
func ServeWs(hub *Hub, db *gorm.DB, c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	userId, err := auth.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		ID:   userId,
		Hub:  hub,
		Conn: conn,
		Send: make(chan []byte, 256),
		DB:   db, // Gán DB cho client dùng
	}
	client.Hub.Register <- client

	go client.writePump()
	go client.readPump()
}
