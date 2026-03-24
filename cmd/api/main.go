package main

import (
	"golang-playground/internal/database"
	"golang-playground/internal/delivery/http"
	"golang-playground/internal/messaging"
	"golang-playground/pkg/logger"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

	// database
	db := database.SettupDatabase()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Database not connected")
	}
	log.Println("Database connected!")

	//publisher
	brokers := []string{"localhost:9092"}
	publisher, err := messaging.NewKafkaPublisher(brokers)
	if err != nil {
		log.Fatalf("failed to create publisher: %v", err)
	}
	defer publisher.Close()

	//running gin
	port := os.Getenv("PORT")
	mode := os.Getenv("MODE")

	// debug
	if mode == "debug" {
		mode = gin.DebugMode
		gin.SetMode(mode)
	}

	r := http.SettupRoute(db, publisher)
	r.Run(":" + port)
}

func getEnv(s1, s2 string) {
	panic("unimplemented")
}
