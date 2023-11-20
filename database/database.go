package database

import (
	"crud_go/config"
	"crud_go/models"
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() {
	var err error

	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("Failed to parse database")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"), port, config.Config("DB__USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("fail connection to database")
	}

	fmt.Println("connection open to database")
	db.AutoMigrate(&models.User{})
	fmt.Println("Database migrated")
}
