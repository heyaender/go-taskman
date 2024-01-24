package handlers

import (
	"go-tugasku/database"
	"go-tugasku/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {

	// Connect to database
	db := database.MySQLDB.DB
	if db == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error":  "Database not connected!",
		})
	}

	var createTaskRequest *models.CreateTaskRequest
	err := c.BodyParser(&createTaskRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Cannot parse JSON!",
			"error":   err.Error(),
		})
	}

	newTask := models.Task{
		Title:       createTaskRequest.Title,
		Description: createTaskRequest.Description,
		Status:      createTaskRequest.Status,
		Deadline:    createTaskRequest.Deadline,
	}

	result := db.Create(&newTask)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Cannot create new task!",
			"error":   result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "New task created!",
		"data":    newTask,
	})
}
