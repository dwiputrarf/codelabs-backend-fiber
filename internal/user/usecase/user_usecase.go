package usecase

import (
	"codelabs-backend-fiber/internal/user/domain"
	"codelabs-backend-fiber/pkg/security"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) GetAll() ([]domain.User, error) {
	return u.repo.FindAll()
}

func (u *userUsecase) GetByID(id uint) (*domain.User, error) {
	return u.repo.FindByID(id)
}

func (u *userUsecase) Create(user *domain.User) error {
	hashed, err := security.HashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashed

    if user.Role == "" {
        user.Role = "user"
    }
	
	return u.repo.Create(user)
}
