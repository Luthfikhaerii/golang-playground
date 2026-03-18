package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// migrate create -ext sql -dir migrations create_users_table
// migrate -path migrations -database "mysql://root@tcp(localhost:3306)/golang" up
// migrate -path migrations -database "mysql://root:password@tcp(localhost:3306)/dbname" down 1

func SettupDatabase() *gorm.DB {
	host := "localhost"
	port := "3306"
	user := "root"
	password := ""
	dbname := "golang"

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

	return db
}
