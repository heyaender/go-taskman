package handlers

import (
	"go-tugasku/database"
	"go-tugasku/models"

	"github.com/gofiber/fiber/v2"
)

func ShowTasks(c *fiber.Ctx) error {

	// connect db
	db := database.MySQLDB.DB
	if db == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error":  "Database not connected!",
		})
	}

	var tasks []models.Task
	db.Find(&tasks)

	for _, task := range tasks {
		c.JSON(task)
	}

	if len(tasks) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "You don't have any tasks!",
			"data":    tasks,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "This is all your tasks!",
		"data":    tasks,
	})
}
