package usecase

import (
	"errors"
	"golang-playground/internal/dto"
	"golang-playground/internal/model"
	"golang-playground/internal/repository"
	"golang-playground/pkg/jwt"
)

type UserUsecase struct {
	repo *repository.UserRepository
}

func NewUserUsecase(repo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) Create(user dto.CreateUserRequest) (*dto.Auth, error) {
	if existing, _ := u.repo.FindOneByEmail(user.Email); existing != nil {
		return nil, errors.New("email already exists")
	}
	newUser := &model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	if err := u.repo.Create(newUser); err != nil {
		return nil, err
	}
	token, err := jwt.GenerateToken(newUser)
	if err != nil {
		return nil, err
	}

	return &dto.Auth{
		Token: token,
	}, nil
}
