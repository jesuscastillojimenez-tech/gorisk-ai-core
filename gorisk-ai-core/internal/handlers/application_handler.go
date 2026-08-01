package handlers

import (
	"encoding/json"
	"net/http"
	"time"

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

	// We consulted the AI ​​engine before saving
	aiDecision := services.EvaluateRisk(input.MonthlyIncome, input.RequestedAmount)
	input.Status = aiDecision.Status

	// 2. Insert the record into PostgreSQL using GORM
	result := database.DB.Create(&input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create application"})
		return
	}

	// 3. Return the created application with HTTP 201 Created
	services.SaveAuditLog(input.ID, "CREATE_APPLICATION", "Credit application created and stored in PostgreSQL")
	services.PublishRiskEvaluation(input.ID, input.Status)
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

// GetApplicationByID retrieves a single application using Redis cache
func GetApplicationByID(c *gin.Context) {
	id := c.Param("id")
	cacheKey := "application:" + id

	// 1. Intenta buscar primero en la memoria RAM de Redis
	cachedData, err := database.RedisClient.Get(c, cacheKey).Result()
	if err == nil {
		// CACHE HIT: ¡Se encontró en Redis! Se responde al instante
		c.Header("X-Cache-Status", "HIT")
		c.Data(http.StatusOK, "application/json", []byte(cachedData))
		return
	}

	// 2. CACHE MISS: No estaba en Redis, así que se busca en PostgreSQL
	var app models.CreditApplication
	if err := database.DB.First(&app, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	// 3. Guarda el resultado en Redis por 5 minutos (TTL)
	jsonBytes, _ := json.Marshal(app)
	database.RedisClient.Set(c, cacheKey, jsonBytes, 5*time.Minute)

	// 4. Responde la solicitud desde PostgreSQL
	c.Header("X-Cache-Status", "MISS")
	c.JSON(http.StatusOK, app)
}
