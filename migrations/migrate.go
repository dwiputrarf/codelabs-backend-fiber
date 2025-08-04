package database

import (
    "log"

    "gorm.io/gorm"
    "codelabs-backend-fiber/internal/user/domain"
)

func Migrate(db *gorm.DB) {
    err := db.AutoMigrate(&domain.User{})
    if err != nil {
        log.Fatalf("AutoMigrate failed: %v", err)
    }
}