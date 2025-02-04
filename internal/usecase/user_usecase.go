package usecase

import (
	"errors"

	"golang-user-auth/internal/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type UserUsecase interface {
	RegisterUser(email string, password string) error
}

type userUsecase struct {
	userRepo UserRepository
}

func NewUserUsecase(repo UserRepository) UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (u *userUsecase) RegisterUser(email string, password string) error {
	if exiting, err := u.userRepo.FindByEmail(email); err == nil && exiting != nil {
		return errors.New("ユーザーは既に存在します")
	}

	user := &entity.User{
		Email:    email,
		Password: password,
	}

	return u.userRepo.Create(user)
}
