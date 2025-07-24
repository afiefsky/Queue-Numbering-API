package routes

import (
	"queue-numbering-api/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Order endpoints
	app.Post("/orders", controller.CreateOrder)

	// Payment endpoints
	// app.Post("/payments", controller.PayOrder)
}
