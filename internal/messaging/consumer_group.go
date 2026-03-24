package messaging

import (
	"log"
	"os"

	"github.com/IBM/sarama"
)

func NewKafkaConsumerGroup(groupID string) sarama.ConsumerGroup {
	brokers := []string{os.Getenv("KAFKA_BROKERS")}

	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{
		sarama.NewBalanceStrategyRoundRobin(),
	}
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	//group
	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Fatalf("failed to create consumer group: %v", err)
	}

	return consumerGroup
}
