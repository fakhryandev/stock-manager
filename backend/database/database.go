package database

import (
	"fmt"
	"log"
	"stock-manager/config"
	"stock-manager/models"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Error parsing str to int")
	}

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Config("DB_HOST"), config.Config("DB_USERNAME"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)

	Db, err = gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Success to connect database")

	Db.AutoMigrate(&models.Item{})

	fmt.Println("Database migrated")
}
