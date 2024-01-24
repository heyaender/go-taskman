package handlers

import (
	"go-tugasku/database"
	"go-tugasku/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteTask(c *fiber.Ctx) error {

	db := database.MySQLDB.DB
	if db == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error":  "Database not connected!",
		})
	}

	var task *models.Task
	checkTask := db.First(&task, c.Params("id"))
	if checkTask.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Task not found!",
		})
	}

	db.Delete(&task)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Task deleted!",
	})
}
