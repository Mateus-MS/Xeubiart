package appointment_model

import (
	"errors"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidTimezone = errors.New("invalid timezone")
)

// An appointment is a meeting to plan the tattoo with the customer,
// covering design, location, and size.
type AppointmentEntity struct {
	ID     primitive.ObjectID
	UserID primitive.ObjectID
	Date   internal_datetime.UTCTime
}

// It expects a local date to be converted to UTC
func NewEntity(userID primitive.ObjectID, localDate *internal_datetime.LocalTime) (*AppointmentEntity, error) {
	// Converts the received local time to UTC
	utcTime := localDate.ToUTCTime()

	return &AppointmentEntity{
		ID:     primitive.NewObjectID(),
		UserID: userID,
		Date:   utcTime,
	}, nil
}
