package main

import (
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
	port := os.Getenv("PORT")
	mode := os.Getenv("MODE")

	//running gin
	if mode == "realease" {
		mode := gin.ReleaseMode
		gin.SetMode(mode)
	}

	r := http.SettupRoute()

	// database

	r.Run(":" + port)
}
