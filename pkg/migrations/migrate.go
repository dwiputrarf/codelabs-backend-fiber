package database

import (
    "codelabs-backend-fiber/internal/model"
    "log"
)

func Migrate() {
    err := DB.AutoMigrate(&model.User{})
    if err != nil {
        log.Fatalf("AutoMigrate failed: %v", err)
    }
}