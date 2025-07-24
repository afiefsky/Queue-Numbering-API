package controller

import (
	"log"
	"queue-numbering-api/service"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx) error {
	var req service.CreateOrderRequest

	if err := c.BodyParser(&req); err != nil {
		log.Printf("❌ Body parse error: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	order, err := service.CreateOrder(c.Context(), req)
	if err != nil {
		log.Printf("❌ CreateOrder error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create order")
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}
