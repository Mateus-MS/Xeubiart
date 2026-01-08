package booking_service

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidBookingDate = errors.New("the given date is invalid")
)

func (s *service) Create(ctx context.Context, appointment *BookingEntity) error {
	t := time.Now().UTC()

	// Any date from now to one hour later, and from more than 1 year is invalid
	if appointment.Date.Before(t.Add(time.Hour)) || appointment.Date.After(t.AddDate(1, 0, 0)) {
		return ErrInvalidBookingDate
	}

	// TODO: prevent booking conflicts
	// - A booking cannot be booked in the same day as another booking by a user, in this case only can be booked by a ADM
	// - A booking cannot be booked out from working hours (9-18) (9am-6pm)UTC

	return s.repository.Create(ctx, appointment)
}

func (s *service) ReadByUserID(ctx context.Context, userID primitive.ObjectID) (*BookingEntity, error) {
	return s.repository.ReadByUserID(ctx, userID)
}
