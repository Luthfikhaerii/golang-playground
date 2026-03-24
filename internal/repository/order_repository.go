package repository

import (
	"golang-playground/internal/model"
	"golang-playground/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(user *model.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		logger.Log.Error("db_insert_order_failed",
			zap.Error(err),
		)
	}
	return nil
}
