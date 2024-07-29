package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"messaggioService/models"
	"os"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	DB.AutoMigrate(&models.Message{})
}

func CloseDB() {
	DB.Close()
}
