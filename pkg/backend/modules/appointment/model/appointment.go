package appointment_model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidTimezone = errors.New("invalid timezone")
)

type AppointmentEntity struct {
	UserID primitive.ObjectID
	Date   time.Time
}

// It expects the raw date to be converted to UTC
func NewEntity(userID primitive.ObjectID, localDate time.Time, timezone string) (*AppointmentEntity, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return &AppointmentEntity{}, ErrInvalidTimezone
	}

	// Overrides the timezone to prevent any double conversion
	tInLocation := time.Date(
		localDate.Year(), localDate.Month(), localDate.Day(),
		localDate.Hour(), localDate.Minute(),
		localDate.Second(), localDate.Nanosecond(),
		location,
	)

	return &AppointmentEntity{
		UserID: userID,
		Date:   tInLocation.UTC(),
	}, nil
}
