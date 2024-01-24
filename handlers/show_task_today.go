package handlers

import (
	"go-tugasku/database"
	"go-tugasku/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ShowTasksToday(c *fiber.Ctx) error {

	db := database.MySQLDB.DB
	if db == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error":  "Database not connected!",
		})
	}

	var tasks []models.Task

	db.Not("status", "completed").Where("due_date = ?", time.Now().Format("2006-01-02")).Find(&tasks)

	if len(tasks) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "No tasks today!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Here's your tasks today!",
		"data":    tasks,
	})
}
