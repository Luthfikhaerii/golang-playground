package messaging

import (
	"context"

	"github.com/IBM/sarama"
	"go.uber.org/zap"

	"golang-playground/pkg/logger"
)

type ConsumerHandler func(message *sarama.ConsumerMessage) error

type ConsumerGroupHandler struct {
	Handler ConsumerHandler
}

func (h *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			if message == nil {
				return nil
			}
			if err := h.Handler(message); err != nil {
				logger.Log.Error("failed to process message", zap.Error(err))
			} else {
				session.MarkMessage(message, "")
			}
		case <-session.Context().Done():
			return nil
		}
	}
}

func ConsumeTopic(ctx context.Context, consumerGroup sarama.ConsumerGroup, topic string, handler ConsumerHandler) {
	consumerHandler := &ConsumerGroupHandler{Handler: handler}

	go func() {
		for {
			if err := consumerGroup.Consume(ctx, []string{topic}, consumerHandler); err != nil {
				logger.Log.Error("error from consumer", zap.Error(err))
			}
			if ctx.Err() != nil {
				logger.Log.Info("context cancelled, stopping consumer")
				return
			}
		}
	}()

	go func() {
		for err := range consumerGroup.Errors() {
			logger.Log.Error("consumer group error", zap.Error(err))
		}
	}()

	<-ctx.Done()
	logger.Log.Info("closing consumer", zap.String("topic", topic))
	if err := consumerGroup.Close(); err != nil {
		logger.Log.Error("error closing consumer group", zap.Error(err))
	}
}
