package messaging

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"go.uber.org/zap"

	"golang-playground/internal/dto"
	"golang-playground/pkg/logger"
)

type OrderNotificationConsumer struct{}

func NewOrderNotificationConsumer() *OrderNotificationConsumer {
	return &OrderNotificationConsumer{}
}

func (c *OrderNotificationConsumer) Consume(message *sarama.ConsumerMessage) error {
	event := new(dto.OrderCreatedEvent)
	if err := json.Unmarshal(message.Value, event); err != nil {
		logger.Log.Error("error unmarshalling order notification event", zap.Error(err))
		return err
	}

	// TODO: process event
	logger.Log.Info("received order notification",
		zap.String("order_id", event.OrderID),
		zap.String("user_id", event.UserID),
		zap.Int32("partition", message.Partition),
	)

	return nil
}
