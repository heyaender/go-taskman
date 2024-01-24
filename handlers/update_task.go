package handlers

import (
	"go-tugasku/database"
	"go-tugasku/models"

	"github.com/gofiber/fiber/v2"
)

func UpdateTask(c *fiber.Ctx) error {

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

	var markTaskCompleteRequest *models.MarkTaskCompleteRequest
	err := c.BodyParser(&markTaskCompleteRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Cannot parse JSON!",
			"error":   err.Error(),
		})
	}

	task.Status = markTaskCompleteRequest.Status

	result := db.Save(&task)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Cannot update task!",
			"error":   result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Task updated!",
		"data":    task,
	})
}
