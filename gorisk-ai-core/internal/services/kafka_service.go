package services

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

// KafkaEvent representa la estructura del mensaje que enviaremos al Topic
type KafkaEvent struct {
	ApplicationID uint    `json:"application_id"`
	Status        string  `json:"status"`
	Score         float64 `json:"score,omitempty"`
}

// PublishRiskEvaluation envía un mensaje asíncrono a Kafka
func PublishRiskEvaluation(appID uint, status string) {
	broker := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")

	// Si no hay variables de entorno, no intentamos conectar
	if broker == "" || topic == "" {
		log.Println("Warning: Will not be published to Kafka because the KAFKA_BROKER or KAFKA_TOPIC variables are missing")
		return
	}

	// 1. Configuramos el "Escritor" (Productor) de Kafka
	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	// 2. Preparamos el mensaje en formato JSON
	event := KafkaEvent{
		ApplicationID: appID,
		Status:        status,
	}
	eventBytes, _ := json.Marshal(event)

	// 3. Enviamos el mensaje
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("risk_eval"), // Llave para agrupar los mensajes
			Value: eventBytes,          // El contenido real (JSON)
		},
	)

	if err != nil {
		log.Printf("Error publishing to Kafka: %v\n", err)
		return
	}

	log.Printf("Success: Event published to Kafka for ApplicationID: %d\n", appID)
}
