package main

import (
	"golang-playground/internal/database"
	"golang-playground/internal/delivery/http"
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

	//running gin
	port := os.Getenv("PORT")
	mode := os.Getenv("MODE")

	// debug
	if mode == "debug" {
		mode = gin.DebugMode
		gin.SetMode(mode)
	}
	r := http.SettupRoute()
	r.Run(":" + port)
}
