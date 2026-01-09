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

// NewLocalWithLocationOverride creates a new LocaTime from a time ignoring any setted timezone conversion overriding it
// Example: 2026-01-01 12:00 +8 Asia/Shanghai -> 2026-01-01 12:00 -5 America/New_York. -- Ignoring the 13 hours difference
// Keep this func usage only to tests...
func NewLocalWithLocationOverride(t time.Time, loc *time.Location) (*LocalTime, error) {
	if loc == time.UTC {
		return nil, ErrUTCTime
	}

	temp := time.Date(
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
		t.Nanosecond(),
		loc,
	)

	return &LocalTime{
		Time: temp,
	}, nil
}
