package models

import "time"

// AuditLog representa un registro no estructurado para auditar operaciones
type AuditLog struct {
	ID            string    `bson:"_id,omitempty" json:"id"`
	ApplicationID uint      `bson:"application_id" json:"application_id"`
	Action        string    `bson:"action" json:"action"`
	Details       string    `bson:"details" json:"details"`
	Timestamp     time.Time `bson:"timestamp" json:"timestamp"`
}
