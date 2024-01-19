package routes

import (
	"go-tugasku/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// ----- API Routes -----
	api := app.Group("/api")

	// ----- Tools -----
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	app.Get("/env", func(c *fiber.Ctx) error {
		return c.SendString("Environment: " + os.Getenv("TEST_ENV"))
	})

	// ----- Versioning -----
	v1 := api.Group("/v1")

	// ----- API Routes -----
	v1.Post("/tasks/create", handlers.CreateTask)

}
