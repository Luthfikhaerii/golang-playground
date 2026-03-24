package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"golang-playground/internal/database"
	deliveryMessaging "golang-playground/internal/delivery/messaging"
	"golang-playground/internal/messaging"
	"golang-playground/internal/repository"
	"golang-playground/internal/usecase"
	"golang-playground/pkg/logger"
	"log"
)

func main() {
	// env
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed load .env")
	}

	// logger
	logger.Init()
	defer logger.Log.Sync()
	logger.Log.Info("starting worker service")

	// dependencies
	brokers := []string{"localhost:9092"}
	publisher, err := messaging.NewKafkaPublisher(brokers)
	if err != nil {
		log.Fatalf("failed to create publisher: %v", err)
	}
	defer publisher.Close()

	db := database.SettupDatabase()
	orderRepo := repository.NewOrderRepository(db)
	orderUsecase := usecase.NewOrderUsecase(publisher, orderRepo)

	// context
	ctx, cancel := context.WithCancel(context.Background())

	// run consumers
	go RunPaymentResultConsumer(ctx, orderUsecase)
	go RunOrderNotificationConsumer(ctx)
	// tambah consumer baru? tinggal go RunXxxConsumer(ctx, ...) di sini

	// SYSTEM SHUTDOWN CRASH
	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGTERM)

	stop := false
	for !stop {
		select {
		case s := <-terminateSignals:
			logger.Log.Info("shutting down worker", zap.String("signal", s.String()))
			cancel()
			stop = true
		}
	}

	time.Sleep(5 * time.Second)
}

// Running consumer Payment
func RunPaymentResultConsumer(ctx context.Context, orderUsecase *usecase.OrderUsecase) {
	logger.Log.Info("setup payment result consumer")

	//group
	consumerGroup := messaging.NewKafkaConsumerGroup(os.Getenv("KAFKA_GROUP_ID"))

	//handler
	handler := deliveryMessaging.NewPaymentResultConsumer(orderUsecase)

	//group, topic, & handler
	messaging.ConsumeTopic(ctx, consumerGroup, messaging.TopicOrderPaymentResult, handler.Consume)
}

// Running consumer Order
func RunOrderNotificationConsumer(ctx context.Context) {
	logger.Log.Info("setup order notification consumer")
	consumerGroup := messaging.NewKafkaConsumerGroup(os.Getenv("KAFKA_GROUP_ID"))
	handler := deliveryMessaging.NewOrderNotificationConsumer()
	messaging.ConsumeTopic(ctx, consumerGroup, messaging.TopicOrderNotification, handler.Consume)
}
