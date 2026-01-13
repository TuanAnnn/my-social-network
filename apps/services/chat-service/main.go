package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Go Chat Service!",
			"service": "chat-service",
			"status":  "fast & furious",
		})
	})

	r.GET("/chat/history", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"room_id":  "123",
			"messages": []string{"Hello", "Hi there!", "How are you?"},
		})
	})

	// Chạy ở port 8080
	r.Run(":8080")
}
