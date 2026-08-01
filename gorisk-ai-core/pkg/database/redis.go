package database

import (
	"context"
	"crypto/tls"
	"log"

	"github.com/redis/go-redis/v9"
)

// RedisClient guarda la conexión global a la base de datos en memoria RAM
var RedisClient *redis.Client

// ConnectRedis abre la conexión con Redis
func ConnectRedis() *redis.Client {
	// Buscamos REDIS_ADDR primero (la variable de Docker), si no existe caemos a REDIS_HOST
	redisAddr := getEnv("REDIS_ADDR", getEnv("REDIS_HOST", "localhost:6379"))
	redisPassword := getEnv("REDIS_PASSWORD", "")

	// Configuramos las opciones base
	opt := &redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	}

	// Solo habilitamos TLS si estamos usando Upstash (si la contraseña no está vacía o si hay un flag específico)
	// Para el Redis local en Docker (sin contraseña), no usamos TLS.
	if redisPassword != "" {
		opt.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	rdb := redis.NewClient(opt)

	// Enviamos un mensaje de prueba ("PING") para confirmar la conexión
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Redis connection successfully established.")

	RedisClient = rdb
	return rdb
}
