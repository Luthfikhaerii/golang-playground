package usecase

import (
	"errors"
	"golang-playground/internal/dto"
	"golang-playground/internal/model"
	"golang-playground/internal/repository"
	"golang-playground/pkg/bcrypt"
	"golang-playground/pkg/jwt"
	"golang-playground/pkg/logger"

	"go.uber.org/zap"
)

type UserUsecase struct {
	repo *repository.UserRepository
}

func NewUserUsecase(repo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) Create(user dto.CreateUserRequest) (*dto.Auth, error) {
	existing, err := u.repo.FindOneByEmail(user.Email)
	if err != nil {
		logger.Log.Error("find_user_failed",
			zap.Error(err),
		)
		return nil, err
	}

	if existing != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.HashPassword(user.Password)
	if err != nil {
		logger.Log.Error("hash_password_failed",
			zap.Error(err),
		)

		return nil, err
	}

	newUser := &model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}
	if err := u.repo.Create(newUser); err != nil {
		logger.Log.Error("create_user_failed",
			zap.Error(err),
		)
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
