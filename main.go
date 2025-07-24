package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Constants
const (
	defaultMongoURI  = "mongodb://localhost:27017"
	defaultRedisAddr = "localhost:6379"
)

// Global variables
var (
	MongoClient *mongo.Client
	MongoDB     *mongo.Database
	RedisClient *redis.Client
	ctx         = context.Background()
)

// MongoDB setup
func initMongo() {
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

	MongoClient = client
	MongoDB = client.Database("queue_api")
	log.Println("✅ MongoDB connected")
}

// Redis setup
func initRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     defaultRedisAddr,
		Password: "",
		DB:       0,
	})

	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("❌ Redis ping error: %v", err)
	}

	log.Println("✅ Redis connected")
}

// Entry point
func main() {
	initMongo()
	initRedis()

	app := fiber.New()
	SetupRoutes(app)

	log.Println("🚀 Server running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
