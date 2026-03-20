package repository

import (
	"golang-playground/internal/model"
	"golang-playground/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		logger.Log.Error("db_insert_user_failed",
			zap.Error(err),
		)
	}
	return nil
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	if err != nil {
		logger.Log.Error("db_find_all_user_failed",
			zap.Error(err),
		)
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindOneById(id int) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Log.Error("db_find_one_by_id_user_failed",
			zap.Error(err),
		)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindOneByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		logger.Log.Error("db_find_one_by_email_user_failed",
			zap.Error(err),
		)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(id int, user model.User) error {
	return r.db.Model(&model.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *UserRepository) Delete(id int) error {
	var user model.User
	return r.db.Delete(user, id).Error
}
