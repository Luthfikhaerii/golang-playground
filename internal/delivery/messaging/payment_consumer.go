package messaging

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"go.uber.org/zap"

	"golang-playground/internal/dto"
	"golang-playground/internal/usecase"
	"golang-playground/pkg/logger"
)

type PaymentResultConsumer struct {
	orderUsecase *usecase.OrderUsecase
}

func NewPaymentResultConsumer(orderUsecase *usecase.OrderUsecase) *PaymentResultConsumer {
	return &PaymentResultConsumer{orderUsecase: orderUsecase}
}

func (c *PaymentResultConsumer) Consume(message *sarama.ConsumerMessage) error {
	event := new(dto.PaymentResultEvent)
	if err := json.Unmarshal(message.Value, event); err != nil {
		logger.Log.Error("error unmarshalling payment result event", zap.Error(err))
		return err
	}

	// TODO: process event
	logger.Log.Info("received payment result",
		zap.String("order_id", event.OrderID),
		zap.String("status", event.Status),
		zap.Int32("partition", message.Partition),
	)

	return c.orderUsecase.HandlePaymentResult(*event)
}
