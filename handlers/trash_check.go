package handlers

import (
	"go-tugasku/database"
	"go-tugasku/models"

	"github.com/gofiber/fiber/v2"
)

func CheckTrash(c *fiber.Ctx) error {

	db := database.MySQLDB.DB
	if db == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error":  "Database not connected!",
		})
	}

	var tasks []models.Task

	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&tasks)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Trash checked!",
		"data":    tasks,
	})
}
