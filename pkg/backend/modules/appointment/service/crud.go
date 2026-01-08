package appointment_service

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidAppointmentDate = errors.New("the given date is invalid")
)

func (s *service) Create(ctx context.Context, appointment *AppointmentEntity) error {
	t := time.Now().UTC()

	// Any date from now to one hour later, and from more than 1 year is invalid
	if appointment.Date.Before(t.Add(time.Hour)) || appointment.Date.After(t.AddDate(1, 0, 0)) {
		return ErrInvalidAppointmentDate
	}

	// TODO: prevent appointment conflicts
	// - An appointment cannot be appointed within less than 2 hours from another
	// - An appointment cannot be appointed in the same day as a booking
	// - An appointment cannot be appointed out from working hours (9-18) (9am-6pm)UTC

	return s.repository.Create(ctx, appointment)
}

func (s *service) ReadByUserID(ctx context.Context, userID primitive.ObjectID) (*AppointmentEntity, error) {
	return s.repository.ReadByUserID(ctx, userID)
}
