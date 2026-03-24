package messaging

import (
	"encoding/json"
	"golang-playground/internal/dto"
	"log"

	"github.com/IBM/sarama"
)

type NotificationConsumer struct{}

func NewNotificationConsumer() *NotificationConsumer {
	return &NotificationConsumer{}
}

func (c *NotificationConsumer) Consume(msg *sarama.ConsumerMessage) error {
	var event dto.OrderCreatedEvent
	if err := json.Unmarshal(msg.Value, &event); err != nil {
		return err
	}
	log.Println("send notification:", string(msg.Value))
	return nil
}
