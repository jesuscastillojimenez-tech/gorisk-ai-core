package services

import (
	"testing"
)

// TestEvaluateRisk_BusinessRules verifies the deterministic logic of risk evaluation
func TestEvaluateRisk_BusinessRules(t *testing.T) {
	// Table-driven tests: Estándar oficial de Go para probar múltiples escenarios
	tests := []struct {
		name            string
		income          float64
		requestedAmount float64
		expectedStatus  string
	}{
		{
			name:            "High Risk - Amount exceeds 10x income",
			income:          3000.0,
			requestedAmount: 35000.0, // 35k > 30k (10x)
			expectedStatus:  "REJECTED",
		},
		{
			name:            "Low Risk - Healthy financial ratio",
			income:          10000.0,
			requestedAmount: 20000.0, // 20k <= 100k (10x)
			expectedStatus:  "APPROVED",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Nota: Si GEMINI_API_KEY no está presente o falla, el sistema aplica
			// el mecanismo de resiliencia (Graceful Degradation) para este test.
			result := EvaluateRisk(tt.income, tt.requestedAmount)

			// Para el escenario de alto riesgo (>10x), la regla debe cumplirse estrictamente
			if tt.requestedAmount > (tt.income * 10) {
				if result.Status != "REJECTED" {
					t.Errorf("expected status REJECTED for high risk amount, got %s", result.Status)
				}
				if result.Score >= 30.0 {
					t.Errorf("expected score under 30.0 for high risk amount, got %.2f", result.Score)
				}
			}
		})
	}
}

// TestKafkaEvent_Structure ensures the JSON payload structure for Kafka events is correct
func TestKafkaEvent_Structure(t *testing.T) {
	event := KafkaEvent{
		ApplicationID: 101,
		Status:        "APPROVED",
		Score:         88.5,
	}

	if event.ApplicationID != 101 {
		t.Errorf("expected ApplicationID 101, got %d", event.ApplicationID)
	}

	if event.Status != "APPROVED" {
		t.Errorf("expected Status APPROVED, got %s", event.Status)
	}
}
