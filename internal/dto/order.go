package dto

import "time"

type OrderCreatedEvent struct {
	OrderID   string    `json:"order_id"`
	UserID    string    `json:"user_id"`
	Items     []string  `json:"items"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type PaymentResultEvent struct {
	OrderID string    `json:"order_id"`
	Status  string    `json:"status"` // "success" | "failed"
	PaidAt  time.Time `json:"paid_at"`
}

type CreateOrderRequest struct {
	UserID string   `json:"user_id"`
	Items  []string `json:"items"`
	Amount float64  `json:"amount"`
}

type CreateOrderResponse struct {
	OrderID string `json:"order_id"`
	Message string `json:"message"`
}
