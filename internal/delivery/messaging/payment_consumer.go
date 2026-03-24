package messaging

import (
	"encoding/json"

	"golang-playground/internal/dto"
	"golang-playground/internal/usecase"

	"github.com/IBM/sarama"
)

type PaymentConsumer struct {
	usecase *usecase.OrderUsecase
}

func NewPaymentConsumer(u *usecase.OrderUsecase) *PaymentConsumer {
	return &PaymentConsumer{usecase: u}
}

func (c *PaymentConsumer) Consume(msg *sarama.ConsumerMessage) error {
	var event dto.OrderCreatedEvent
	if err := json.Unmarshal(msg.Value, &event); err != nil {
		return err
	}

	return c.usecase.ProcessPayment(event.OrderID)
}
