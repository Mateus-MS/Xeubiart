package booking_model

import (
	"time"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingEntity struct {
	UserID primitive.ObjectID
	Date   time.Time
}

// It expects the raw date to be converted to UTC
func NewEntity(userID primitive.ObjectID, localDate time.Time, timezone string) (*BookingEntity, error) {
	dateUTC, err := internal_datetime.NormalizeToUTC(localDate, timezone)
	if err != nil {
		return nil, err
	}

	return &BookingEntity{
		UserID: userID,
		Date:   dateUTC,
	}, nil
}
