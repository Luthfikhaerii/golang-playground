package model

import (
	"time"
)

type Order struct {
	OrderID   string    `json:"order_id"   gorm:"primaryKey;column:order_id"`
	UserID    string    `json:"user_id"    gorm:"column:user_id"`
	Items     []string  `json:"items"      gorm:"column:items;serializer:json"`
	Amount    float64   `json:"amount"     gorm:"column:amount"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}
