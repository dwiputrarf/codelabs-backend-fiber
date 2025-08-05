package domain

import (
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	gorm.Model
	FullName string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     Role   `gorm:"type:user_role;default:'user';not null"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(id uint) (*User, error)
	Create(user *User) error
}

type UserUsecase interface {
	GetAll() ([]User, error)
	GetByID(id uint) (*User, error)
	Create(user *User) error
}
