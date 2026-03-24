package messaging

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type EventPublisher interface {
	Publish(topic string, key string, payload any) error
}

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer(p sarama.SyncProducer) *KafkaProducer {
	return &KafkaProducer{producer: p}
}

func (p *KafkaProducer) Send(topic string, key string, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(data),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed send: %w", err)
	}

	log.Printf("[Kafka] topic=%s partition=%d offset=%d", topic, partition, offset)
	return nil
}
