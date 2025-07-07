package main

import (
	"github.com/gofiber/fiber/v2"
)
import "jeopardy-backend/database"

func main() {
	app := fiber.New()
	database.Connect()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Jeopardy Backend is running!")
	})

	app.Listen(":5000")
}
