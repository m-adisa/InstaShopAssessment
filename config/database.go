package config

import (
	"instashop/models"
	"path/filepath"

	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Determine environment
	environment := os.Getenv("TESTING")

	var envFile string

	if environment == "true" {
		cwd, _ := os.Getwd()
		envFile = filepath.Join(cwd, "../.env.test")
	} else {
		envFile = ".env"
	}

	// Load appropriate environment variables
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file. Error: %v", envFile, err)
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

	if environment != "true" {
		// Migrate the database
		err = DB.AutoMigrate(&models.Product{}, &models.User{}, &models.Order{})
		if err != nil {
			log.Fatalf("Failed to migrate database. Error: %v", err)
		}

		fmt.Println("Database migrated successfully!")
	}
}
