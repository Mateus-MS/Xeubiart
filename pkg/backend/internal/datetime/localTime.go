package internal_datetime

import (
	"errors"
	"time"
)

var (
	ErrUTCTime = errors.New("in order to create a LocalTime, the given timezone should not be setted as UTC")
)

// LocalTime wraps time.Time to represent a non-UTC timestamp using a user's local timezone.
type LocalTime struct {
	time.Time
}

// ToUTC converts a LocalTime to UTCTime.
// The LocalTime must have a valid timezone set on its underlying time.Time.
func (lt *LocalTime) ToUTCTime() UTCTime {
	return UTCTime{
		Time: lt.Time.UTC(),
	}
}

// NewLocalFromTime creates a LocalTime from a time that must already be with the right timezone setted up.
func NewLocalFromTime(t time.Time) (*LocalTime, error) {
	if t.Location() == time.UTC {
		return nil, ErrUTCTime
	}

	return &LocalTime{
		Time: t,
	}, nil
}

// It expects a string in the following format: year-month-day hour:minute:second
func NewLocalFromString(str string, loc *time.Location) (*LocalTime, error) {
	layout := "2006-01-02 15:04:05"

	t, err := time.ParseInLocation(layout, str, loc)
	if err != nil {
		return &LocalTime{}, err
	}

	lt, err := NewLocalFromTime(t)
	if err != nil {
		return &LocalTime{}, err
	}

	return lt, nil
}
