package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// MongoClient holds the global MongoDB connection
var MongoClient *mongo.Client

// ConnectMongoDB initializes the connection to MongoDB Atlas
func ConnectMongoDB() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI is not set in .env file")
	}

	// 1. Configure the client options with your URI
	clientOptions := options.Client().ApplyURI(uri)

	// 2. Open the connection
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// 3. Ping the database to verify the connection is active
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("MongoDB connection successfully established.")
	MongoClient = client
}
