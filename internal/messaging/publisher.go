package messaging

import (
	"encoding/json"
	"fmt"
	"golang-playground/internal/dto"
	"log"

	"github.com/IBM/sarama"
)

type KafkaPublisher struct {
	producer sarama.SyncProducer
}

func NewKafkaPublisher(brokers []string) (*KafkaPublisher, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 3

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka producer: %w", err)
	}

	return &KafkaPublisher{producer: producer}, nil
}

func (p *KafkaPublisher) Publish(topic string, payload *dto.OrderCreatedEvent) error {
	//convert json
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	//choose partition, topic, & event
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(payload.OrderID),
		Value: sarama.ByteEncoder(data),
	}

	//send message
	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send message to topic %s: %w", topic, err)
	}

	log.Printf("[Publisher] topic=%s partition=%d offset=%d", topic, partition, offset)
	return nil
}

func (p *KafkaPublisher) Close() error {
	return p.producer.Close()
}
