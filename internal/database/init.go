package database

import (
	"fmt"
	"golang-playground/internal/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SettupDatabase() *gorm.DB {
	host := "localhost"
	port := "3306"
	user := "root"
	password := "password"
	dbname := "testdb"

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed connect database")
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("failed migrate")
	}

	return db
}
