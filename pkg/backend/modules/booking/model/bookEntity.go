package booking_model

import (
	"time"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookEntity represents an appointment with a customer for a specific service.
type BookEntity struct {
	UserID primitive.ObjectID
	Date   time.Time
	Book   BookType
}

// It expects the raw date to be converted to UTC
func NewEntity(userID primitive.ObjectID, localDate time.Time, timezone string, bookType BookType) (*BookEntity, error) {
	// Converts the received local time to UTC
	dateUTC, err := internal_datetime.NormalizeToUTC(localDate, timezone)
	if err != nil {
		return nil, err
	}

	// Garants the received bookType is valid
	err = bookType.IsValid()
	if err != nil {
		return &BookEntity{}, err
	}

	return &BookEntity{
		UserID: userID,
		Date:   dateUTC,
		Book:   bookType,
	}, nil
}
