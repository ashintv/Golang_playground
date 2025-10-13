package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env", err)
	}
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN not found")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Open DB", err)
	}
	DB = db
	
}
