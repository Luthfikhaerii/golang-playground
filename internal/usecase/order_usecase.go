package usecase

import (
	"fmt"
	"golang-playground/internal/dto"
	"golang-playground/internal/messaging"
	"log"
	"time"

	"github.com/google/uuid"
)

type OrderUsecase struct {
	publisher *messaging.KafkaPublisher
	orderRepo any
}

func NewOrderUsecase(publisher *messaging.KafkaPublisher, orderRepo any) *OrderUsecase {
	return &OrderUsecase{
		publisher: publisher,
		orderRepo: orderRepo,
	}
}

func (u *OrderUsecase) CreateOrder(req dto.CreateOrderRequest) error {
	orderID := uuid.New().String()

	//save database

	//init value
	event := &dto.OrderCreatedEvent{
		OrderID:   orderID,
		UserID:    req.UserID,
		Amount:    req.Amount,
		Items:     req.Items,
		CreatedAt: time.Now(),
	}

	//publish topic & value
	if err := u.publisher.Publish(messaging.TopicOrderCreated, event); err != nil {
		return fmt.Errorf("failed to publish order event: %w", err)
	}

	return nil
}

func (u *OrderUsecase) HandlePaymentResult(event dto.PaymentResultEvent) error {
	log.Printf("[OrderUsecase] payment result for order %s: %s", event.OrderID, event.Status)

	// iupdate status

	return nil
}
