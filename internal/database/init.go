package database

import (
	"fmt"
	"golang-playground/pkg/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// migrate create -ext sql -dir migrations create_users_table
// migrate -path migrations -database "mysql://root@tcp(localhost:3306)/golang" up
// migrate -path migrations -database "mysql://root:password@tcp(localhost:3306)/dbname" down 1

func SettupDatabase() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	log.Println("DB_HOST:", host)
	log.Println("DB_PORT:", port)
	log.Println("DB_USER:", user)
	log.Println("DB_NAME:", dbname)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Retry konek ke MySQL maksimal 10 kali
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Waiting for database... attempt %d/10", i+1)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		logger.Log.Error("failed connect database")
	}

	logger.Log.Info("Database success conected")

	return db
}
