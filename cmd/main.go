package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jesuscastillojimenez-tech/gorisk-ai-core/pkg/database"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Carga las variables del archivo .env a la memoria del sistema
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: No se encontró el archivo .env, usando variables del sistema")
	}

	// 2. Conecta con PostgreSQL en Neon y ejecuta las migraciones automáticas
	database.ConnectDatabase()

	// 3. Inicializa el servidor web Gin
	r := gin.Default()

	// 4. Endpoint de comprobación básica del servidor
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	// 5. Arranca el servidor escuchando en el puerto 8080
	r.Run(":8080")
}
