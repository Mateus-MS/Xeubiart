package booking_service

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *service) Create(ctx context.Context, appointment *BookingEntity) error {
	t := time.Now().UTC()

	// Any date from now to one hour later, and from more than 1 year is invalid
	if appointment.Date.Before(t.Add(time.Hour)) || appointment.Date.After(t.AddDate(1, 0, 0)) {
		return ErrInvalidBookingDate
	}

	return s.repository.Create(ctx, appointment)
}

func (s *service) ReadByUserID(ctx context.Context, userID primitive.ObjectID) (*BookingEntity, error) {
	return s.repository.ReadByUserID(ctx, userID)
}
