package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jesuscastillojimenez-tech/gorisk-ai-core/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB holds the global database connection pool instance
var DB *gorm.DB

// ConnectDatabase initializes the PostgreSQL connection and runs automatic migrations
func ConnectDatabase() *gorm.DB {
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "gorisk_db")
	port := getEnv("DB_PORT", "5432")
	sslmode := getEnv("DB_SSLMODE", "disable")

	// DSN: Data Source Name formatting
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL database: %v", err)
	}

	log.Println("PostgreSQL connection successfully established.")

	// Automatically migrate the CreditApplication schema
	err = db.AutoMigrate(&models.CreditApplication{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate database models: %v", err)
	}

	log.Println("Database migration completed successfully.")

	DB = db
	return db
}

// getEnv retrieves environment variables with a default fallback mechanism
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
