package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesuscastillojimenez-tech/gorisk-ai-core/internal/models"
	"github.com/jesuscastillojimenez-tech/gorisk-ai-core/internal/services"
	"github.com/jesuscastillojimenez-tech/gorisk-ai-core/pkg/database"
)

// CreateApplication handles incoming HTTP POST requests to create a credit application
func CreateApplication(c *gin.Context) {
	var input models.CreditApplication

	// 1. Bind the incoming JSON body to the CreditApplication struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Insert the record into PostgreSQL using GORM
	result := database.DB.Create(&input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create application"})
		return
	}

	// 3. Return the created application with HTTP 201 Created
	services.SaveAuditLog(input.ID, "CREATE_APPLICATION", "Credit application created and stored in PostgreSQL")
	c.JSON(http.StatusCreated, gin.H{"data": input})
}

// GetApplications handles incoming HTTP GET requests to list all credit applications
func GetApplications(c *gin.Context) {
	var applications []models.CreditApplication

	// 1. Query all records from PostgreSQL
	result := database.DB.Find(&applications)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch applications"})
		return
	}

	// 2. Return the list of applications with HTTP 200 OK
	c.JSON(http.StatusOK, gin.H{"data": applications})
}
