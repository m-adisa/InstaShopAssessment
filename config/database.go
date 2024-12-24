package config

import (
	"instashop/models"

	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file. Error: %v", err)
	}

	// Build the DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	// Connect to the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	fmt.Println("Connected to database successfully!")

	// Migrate the database
	err = DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("Failed to migrate database. Error: %v", err)
	}

	fmt.Println("Database migrated successfully!")
}
