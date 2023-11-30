package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=admin password=admin123 dbname=gorm port=5432"
var DB *gorm.DB

func ConnectDB() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		panic(error)
	} else {
		log.Println("DB connected")
	}
}
