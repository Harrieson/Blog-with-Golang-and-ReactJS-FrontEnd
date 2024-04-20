package database

import (
	"log"
	"os"

	"github.com/Harrieson/golangbackend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Fetching Envirement settings.")
	}

	connection := os.Getenv("ConnString")
	database, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("Could Not Connect to the database")
	} else {
		log.Println("Successful Connection")
	}
	DB = database

	database.AutoMigrate(
		&models.User{},
	)
}
