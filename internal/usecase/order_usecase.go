package usecase

import (
	"golang-playground/internal/dto"
	"golang-playground/internal/messaging"
	"golang-playground/pkg/logger"
	"log"
)

type OrderUsecase struct {
	producer  *messaging.KafkaProducer
	orderRepo any
}

func NewOrderUsecase(producer *messaging.KafkaProducer, orderRepo any) *OrderUsecase {
	return &OrderUsecase{
		producer:  producer,
		orderRepo: orderRepo,
	}
}

func (u *OrderUsecase) ProcessPayment(orderID string) error {
	// logic bisnis
	logger.Log.Info("=====================================================")
	// publish event baru
	return u.producer.Send(
		messaging.TopicPaymentCompleted,
		orderID,
		map[string]any{
			"order_id": orderID,
			"status":   "paid",
		},
	)
}

func (u *OrderUsecase) HandlePaymentResult(event dto.PaymentResultEvent) error {
	log.Printf("[OrderUsecase] payment result for order %s: %s", event.OrderID, event.Status)

	// iupdate status

	return nil
}
