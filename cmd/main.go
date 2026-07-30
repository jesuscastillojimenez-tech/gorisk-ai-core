package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa el motor de Gin (tu servidor web)
	r := gin.Default()

	// Crea un endpoint de prueba
	r.GET("/health", func(c *gin.Context) {
		// Devuelve una respuesta en formato JSON
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	// Enciende el servidor en el puerto 8080
	r.Run(":8080")
}
