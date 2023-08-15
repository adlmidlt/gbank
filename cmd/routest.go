package cmd

import (
	"github.com/adlmidlt/GBank/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Post("/fact", handlers.CreateFact)
}
