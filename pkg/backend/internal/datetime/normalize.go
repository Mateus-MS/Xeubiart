package internal_datetime

import (
	"errors"
	"time"
)

var ErrInvalidTimezone = errors.New("invalid timezone")

// NormalizeToUTC takes a local date and timezone and returns a UTC time
func NormalizeToUTC(localDate time.Time, timezone string) (time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, ErrInvalidTimezone
	}

	tInLocation := time.Date(
		localDate.Year(), localDate.Month(), localDate.Day(),
		localDate.Hour(), localDate.Minute(), localDate.Second(),
		localDate.Nanosecond(),
		location,
	)

	return tInLocation.UTC(), nil
}
