package appointment_model

import (
	"errors"
	"time"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidTimezone = errors.New("invalid timezone")
)

// An appointment is a meeting to plan the tattoo with the customer,
// covering design, location, and size.
type AppointmentEntity struct {
	UserID primitive.ObjectID
	Date   time.Time
}

// It expects a local date to be converted to UTC
func NewEntity(userID primitive.ObjectID, localDate time.Time, timezone string) (*AppointmentEntity, error) {
	// Converts the received local time to UTC
	dateUTC, err := internal_datetime.NormalizeToUTC(localDate, timezone)
	if err != nil {
		return nil, err
	}

	return &AppointmentEntity{
		UserID: userID,
		Date:   dateUTC,
	}, nil
}
