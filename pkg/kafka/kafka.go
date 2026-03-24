package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

func NewConsumerGroup(brokers []string, groupID string) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	return sarama.NewConsumerGroup(brokers, groupID, config)
}

func NewProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 3

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	return producer, nil
}
