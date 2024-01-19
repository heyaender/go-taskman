package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          uint64     `gorm:"primaryKey;autoIncrement;not null"`
	Title       string     `gorm:"size:255;not null"`
	Description string     `gorm:"size:255;not null"`
	Deadline    *time.Time `gorm:"type:datetime;not null"`
}
