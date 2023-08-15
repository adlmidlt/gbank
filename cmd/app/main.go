package main

import (
	"github.com/adlmidlt/GBank/cmd"
	"github.com/adlmidlt/GBank/database"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	cmd.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Anatolii!!!!!")
	})

	log.Fatal(app.Listen(":3537"))
}
