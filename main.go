package main

import (
	"go-tugasku/database"
	"go-tugasku/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {

	database.MySQLConnect()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())
	app.Use(cors.New())
	routes.SetupRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		// return c.Status(404).SendFile("./views/404.html")
		return c.Status(404).JSON(fiber.Map{
			"status":  404,
			"message": "Not Found",
		})
	})
	app.Listen(GetPort())
	log.Println("Server is running on port", GetPort())
	log.Fatal(app.Listen(GetPort()))

}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return ":" + port
}
