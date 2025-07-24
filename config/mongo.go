package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoDB *mongo.Database
)

const defaultMongoURI = "mongodb://localhost:27017"

func InitMongo(ctx context.Context) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = defaultMongoURI
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("❌ Failed to create Mongo client: %v", err)
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := client.Connect(timeoutCtx); err != nil {
		log.Fatalf("❌ Mongo connection failed: %v", err)
	}

	MongoDB = client.Database("queue_api")
	log.Println("✅ MongoDB connected")
}
