package config

import (
	"fmt"
	"log"
	"os"

	"codelabs-backend-fiber/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgreDB() {
    err := godotenv.Load()
    if err != nil {
        log.Println("⚠️  No .env file found, using system environment variables")
    }

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASS"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("❌ Failed to connect to postgre database: %v", err)
    }

    log.Println("✅ Connected to the postgre database")
    DB = db

    db.AutoMigrate(&model.User{})
}
