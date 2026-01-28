package database

import (
	"chat-service/internal/models"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=myuser password=mypassword dbname=social_app port=5433 sslmode=disable"
	}

	var db *gorm.DB
	var err error

	// Retry connection loop
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Đang chờ Database... (%s)\n", err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Không thể kết nối Database:", err)
	}

	// Auto Migrate
	db.AutoMigrate(&models.Message{})
	fmt.Println("✅ Kết nối DB và Migrate thành công!")

	return db
}
