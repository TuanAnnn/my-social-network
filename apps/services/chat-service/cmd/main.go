package main

import (
	"chat-service/internal/database"
	"chat-service/internal/handlers"
	"chat-service/internal/ws"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Kết nối Database
	db := database.Connect()

	// 2. Khởi tạo Hub và chạy ngầm
	hub := ws.NewHub()
	go hub.Run()

	// 3. Khởi tạo HTTP Handler
	chatHandler := handlers.NewChatHandler(db)

	r := gin.Default()

	// CORS Middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from Go Chat Service!"})
	})

	// 4. WebSocket Route (Truyền Hub và DB vào)
	r.GET("/ws", func(c *gin.Context) {
		ws.ServeWs(hub, db, c)
	})

	// 5. API Route
	r.GET("/api/chat/history", chatHandler.GetChatHistory)

	r.Run(":3001")
}
