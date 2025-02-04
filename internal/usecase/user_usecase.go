package usecase

import (
	"errors"

	"golang-user-auth/internal/entity"

	"golang.org/x/crypto/bcrypt"
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

  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("パスワードのハッシュ化に失敗しました")
	}

	user := &entity.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	return u.userRepo.Create(user)
}
