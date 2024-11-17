package database

import (
	"log"
	"lostandfounditemmanagment/models"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file!")
	}
	dsn := os.Getenv("DSN")

	database, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to open a new connection to database" + err.Error())
	}

	DB = database

	database.AutoMigrate(
		models.Item{},
		models.User{},
		models.Report{},
	)
}
