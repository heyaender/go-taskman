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

	var tasksResponse []models.TasksResponse
	for _, task := range tasks {

		tasksResponse = append(tasksResponse, models.TasksResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			Deadline:    task.Deadline,
		})
	}

	if tasksResponse == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Cannot find any tasks!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "This is all your tasks!",
		"data":    tasksResponse,
	})
}
