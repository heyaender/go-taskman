package models

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

const DateTimeFormat = "2006-01-02 15:04:05"

type LocalTime time.Time

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}

	now, err := time.Parse(`"`+DateTimeFormat+`"`, string(data))
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(DateTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, DateTimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t LocalTime) Value() (driver.Value, error) {
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(DateTimeFormat)), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = LocalTime(tTime)
	return nil
}

func (t LocalTime) String() string {
	return time.Time(t).Format(DateTimeFormat)
}

type Task struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey;autoIncrement;not null"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"size:255;not null"`
	Status      string `gorm:"size:50;not null;default:'pending'"`
	Deadline    *LocalTime
}

type CreateTaskRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Deadline    *LocalTime `json:"deadline"`
}
