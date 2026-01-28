package handlers

import (
	"chat-service/internal/auth"
	"chat-service/internal/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChatHandler struct {
	DB *gorm.DB
}

func NewChatHandler(db *gorm.DB) *ChatHandler {
	return &ChatHandler{DB: db}
}

func (h *ChatHandler) GetChatHistory(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	myId, err := auth.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token không hợp lệ"})
		return
	}

	targetIdStr := c.Query("targetId")
	targetId, _ := strconv.Atoi(targetIdStr)

	var messages []models.Message

	h.DB.Where(
		h.DB.Where("sender_id = ? AND receiver_id = ?", myId, targetId).
			Or("sender_id = ? AND receiver_id = ?", targetId, myId),
	).Order("created_at asc").Find(&messages)

	c.JSON(http.StatusOK, messages)
}
