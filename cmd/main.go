package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jesuscastillojimenez-tech/gorisk-ai-core/internal/handlers"
	"github.com/jesuscastillojimenez-tech/gorisk-ai-core/pkg/database"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, relying on system environment variables")
	}

	// 2. Initialize PostgreSQL connection and auto-migrations
	database.ConnectDatabase()

	// 2.1 Initialize MongoDB connection for audit logs
	database.ConnectMongoDB()

	// 3. Initialize the Gin web framework router
	r := gin.Default()

	// 4. Basic health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	// 5. API v1 Router Group (Credit Applications)
	v1 := r.Group("/api/v1")
	{
		v1.POST("/applications", handlers.CreateApplication)
		v1.GET("/applications", handlers.GetApplications)
	}

	// 6. Start the HTTP server on port 8080
	r.Run(":8080")
}
