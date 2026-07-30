package services

import (
	"context"
	"log"
	"time"

	"github.com/jesuscastillojimenez-tech/gorisk-ai-core/internal/models"
	"github.com/jesuscastillojimenez-tech/gorisk-ai-core/pkg/database"
)

// SaveAuditLog inserta un documento de auditoría en la colección "audit_logs" de MongoDB
func SaveAuditLog(appID uint, action string, details string) {
	collection := database.MongoClient.Database("gorisk_audit").Collection("audit_logs")

	audit := models.AuditLog{
		ApplicationID: appID,
		Action:        action,
		Details:       details,
		Timestamp:     time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, audit)
	if err != nil {
		log.Printf("Error saving audit log to MongoDB: %v", err)
		return
	}

	log.Println("Audit log successfully saved to MongoDB.")
}
