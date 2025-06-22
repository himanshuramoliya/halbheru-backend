package database

import (
	"fmt"
	"log"
	"os"

	"halbheru-backend/models"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect initializes the database connection
func Connect() {
	var err error
	
	// Build database connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Connect to database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection established")
}

// Migrate runs database migrations
func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Ride{},
		&models.RidePassenger{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed")
}

// Close closes the database connection
func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("Error getting underlying sql.DB:", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Println("Error closing database connection:", err)
		return
	}

	log.Println("Database connection closed")
}
