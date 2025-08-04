package domain

import (
	"time"
)

type User struct {
	ID        uint
	FullName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
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