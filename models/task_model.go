package models

import (
	"go-tugasku/configs"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          uint64             `gorm:"primaryKey;autoIncrement;not null"`
	Title       string             `gorm:"size:255;not null"`
	Description string             `gorm:"size:255;not null"`
	Status      string             `gorm:"size:50;not null;default:'pending'"`
	Deadline    *configs.LocalTime `gorm:"type:datetime"`
}

type TasksResponse struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Deadline    *configs.LocalTime
}

type CreateTaskRequest struct {
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Status      string             `json:"status"`
	Deadline    *configs.LocalTime `json:"deadline"`
}
