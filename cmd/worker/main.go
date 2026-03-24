package main

import (
	"context"
	"golang-playground/internal/database"
	deliveryMessaging "golang-playground/internal/delivery/messaging"
	"golang-playground/internal/messaging"
	"golang-playground/internal/repository"
	"golang-playground/internal/usecase"
	"golang-playground/pkg/kafka"
	"golang-playground/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	//env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed Load .env")
	}

	//loger
	logger.Init()
	defer logger.Log.Sync()

	// producer
	brokers := []string{"localhost:9092"}
	rawProducer, err := kafka.NewProducer(brokers)
	if err != nil {
		log.Fatalf("failed to create producer: %v", err)
	}
	defer rawProducer.Close()
	producer := messaging.NewKafkaProducer(rawProducer)

	// dependency
	db := database.SettupDatabase()
	orderRepo := repository.NewOrderRepository(db)
	orderUsecase := usecase.NewOrderUsecase(producer, orderRepo)

	// context (untuk graceful shutdown)
	ctx, cancel := context.WithCancel(context.Background())

	// run consumers (pakai goroutine)
	go RunPaymentConsumer(ctx, brokers, orderUsecase)
	go RunNotificationConsumer(ctx, brokers)

	// raceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("shutting down worker...")
	cancel()

	time.Sleep(3 * time.Second)
	log.Println("worker stopped")
}

// PAYMENT CONSUMER
func RunPaymentConsumer(ctx context.Context, brokers []string, orderUsecase *usecase.OrderUsecase) {
	log.Println("setup payment consumer")

	group, _ := kafka.NewConsumerGroup(brokers, "payment-service")
	handler := deliveryMessaging.NewPaymentConsumer(orderUsecase)

	messaging.ConsumeTopic(ctx, group, messaging.TopicOrderCreated, handler.Consume)
}

// NOTIFICATION CONSUMER
func RunNotificationConsumer(ctx context.Context, brokers []string) {
	log.Println("setup notification consumer")

	group, _ := kafka.NewConsumerGroup(brokers, "notification-service")
	handler := deliveryMessaging.NewNotificationConsumer()

	messaging.ConsumeTopic(ctx, group, messaging.TopicPaymentCompleted, handler.Consume)
}
