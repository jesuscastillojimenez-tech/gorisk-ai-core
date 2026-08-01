package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jesuscastillojimenez-tech/gorisk-ai-core/internal/handlers"
	"github.com/jesuscastillojimenez-tech/gorisk-ai-core/pkg/database"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("INFO: No .env file found or failed to load, using default environment variables.")
	}

	// Initialize databases using the exact function names from pkg/database
	database.ConnectDatabase() // PostgreSQL connection (Neon Cloud)
	database.ConnectMongoDB()  // MongoDB connection
	database.ConnectRedis()    // Redis connection

	r := gin.Default()

	// CORS Configuration to allow React frontend (port 5173)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		v1.POST("/applications", handlers.CreateApplication)
		v1.GET("/applications", handlers.GetApplications)
		v1.GET("/applications/:id", handlers.GetApplicationByID)
	}

	log.Println("Server running on port :8080")
	r.Run(":8080")
}
