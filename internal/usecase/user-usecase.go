package usecase

import "golang-playground/internal/repository"

type UserUsecase struct {
	repo *repository.UserRepository
}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{}
}

func (u *UserUsecase) Create(user string) {
	// err := u.repo.Create()
	// return user
}
