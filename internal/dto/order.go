package dto

import "time"

// Event yang di-publish oleh order service
type OrderCreatedEvent struct {
	OrderID   string    `json:"order_id"`
	UserID    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Items     []string  `json:"items"`
	CreatedAt time.Time `json:"created_at"`
}

// Event yang di-consume dari payment service
type PaymentResultEvent struct {
	OrderID string    `json:"order_id"`
	Status  string    `json:"status"` // "success" | "failed"
	PaidAt  time.Time `json:"paid_at"`
}

// HTTP request masuk
type CreateOrderRequest struct {
	UserID string   `json:"user_id"`
	Items  []string `json:"items"`
	Amount float64  `json:"amount"`
}

type CreateOrderResponse struct {
	OrderID string `json:"order_id"`
	Message string `json:"message"`
}
