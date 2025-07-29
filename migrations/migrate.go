package database

import (
	"github.com/dwiputrarf/codelabs-backend-fiber/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &model.User{},
    )
}