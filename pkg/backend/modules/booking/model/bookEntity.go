package booking_model

import (
	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookEntity represents an appointment with a customer for a specific service.
type BookEntity struct {
	UserID primitive.ObjectID
	Date   internal_datetime.UTCTime
	Book   BookType
}

// It expects the raw date to be converted to UTC
func NewEntity(userID primitive.ObjectID, localDate *internal_datetime.LocalTime, bookType BookType) (*BookEntity, error) {
	// Converts the received local time to UTC
	utcTime := localDate.ToUTCTime()

	// Garants the received bookType is valid
	err := bookType.IsValid()
	if err != nil {
		return &BookEntity{}, err
	}

	return &BookEntity{
		UserID: userID,
		Date:   utcTime,
		Book:   bookType,
	}, nil
}
