package repository

import (
	"golang-playground/internal/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}
