package main

import (
	"context"
	"log"

	"queue-numbering-api/config"
	"queue-numbering-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx := context.Background()

	config.InitMongo(ctx)
	config.InitRedis(ctx)

	app := fiber.New()
	routes.Setup(app)

	log.Println("🚀 Server running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
