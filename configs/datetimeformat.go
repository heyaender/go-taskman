package configs

import (
	"database/sql/driver"
	"errors"
	"time"
)

const DateTimeFormat = "2006-01-02 15:04:05"

type LocalTime time.Time

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}

	now, err := time.Parse(`"`+DateTimeFormat+`"`, string(data))
	if err != nil {
		return err
	}
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("null"), nil
	}
	b := make([]byte, 0, len(DateTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, DateTimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t LocalTime) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
		return nil, nil
	}
	return time.Time(t), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if v == nil {
		*t = LocalTime(time.Time{})
		return nil
	}

	tTime, ok := v.(time.Time)
	if !ok {
		return errors.New("failed to scan LocalTime")
	}

	*t = LocalTime(tTime)
	return nil
}

func (t LocalTime) String() string {
	return time.Time(t).Format(DateTimeFormat)
}
