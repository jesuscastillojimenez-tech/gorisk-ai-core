package models

import (
	"time"

	"gorm.io/gorm"
)

// CreditApplication representa la estructura de una solicitud de crédito en la Base de Datos
type CreditApplication struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	ApplicantName   string         `gorm:"not null" json:"applicant_name"`
	MonthlyIncome   float64        `gorm:"not null" json:"monthly_income"`
	RequestedAmount float64        `gorm:"not null" json:"requested_amount"`
	Status          string         `gorm:"default:'PENDING'" json:"status"` // PENDING, APPROVED, REJECTED
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
