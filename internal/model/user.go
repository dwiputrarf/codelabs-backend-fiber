package model

import (
    "time"
)

type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Fullname  string    `gorm:"type:varchar(100);not null" json:"fullname"`
    Username  string    `gorm:"unique;not null" json:"username"`
    Email     string    `gorm:"unique;not null" json:"email"`
    Password  string    `gorm:"not null" json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}