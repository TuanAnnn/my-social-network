package models

import "time"

type Message struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Content    string    `json:"content"`
	SenderID   uint      `json:"sender_id"`
	ReceiverID uint      `json:"receiver_id"`
	CreatedAt  time.Time `json:"created_at"`
}

// Struct dùng để client gửi lên
type MessageInput struct {
	Content    string `json:"content"`
	SenderID   uint   `json:"sender_id"`
	ReceiverID uint   `json:"receiver_id"`
}
