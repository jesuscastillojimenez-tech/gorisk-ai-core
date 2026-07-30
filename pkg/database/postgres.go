package database
package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jesuscastillojimenez-tech/gorisk-ai-core/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB guarda la instancia global de la conexión a la base de datos
var DB *gorm.DB

// ConnectDatabase inicializa la conexión con PostgreSQL y ejecuta las migraciones automáticas
func ConnectDatabase() *gorm.DB {
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "gorisk_db")
	port := getEnv("DB_PORT", "5432")
	sslmode := getEnv("DB_SSLMODE", "disable")

	// DSN: Texto formateado con la información necesaria para conectarse
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos PostgreSQL: %v", err)
	}

	log.Println("Conexión con PostgreSQL establecida con éxito.")

	// Crea o actualiza automáticamente la tabla CreditApplication
	err = db.AutoMigrate(&models.CreditApplication{})
	if err != nil {
		log.Fatalf("Error al migrar la tabla en la base de datos: %v", err)
	}

	log.Println("Migración de la base de datos completada con éxito.")

	DB = db
	return db
}

// getEnv busca variables de entorno y si no existen, usa un valor por defecto
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}