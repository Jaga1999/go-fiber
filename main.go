package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/all", func(c *fiber.Ctx) error {
		return c.SendString("Hello, All!")
	})

	app.Listen(":3000")
}
