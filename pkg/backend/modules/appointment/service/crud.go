package appointment_service

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidAppointmentDate  = errors.New("the given date is invalid")
	ErrAppointmentTimeConflict = errors.New("another appointment exists within 3 hours of the requested time")
)

func (s *service) Create(ctx context.Context, appointment *AppointmentEntity) error {
	t := s.clock.Now()

	// Any date from now to one hour later, and from more than 1 year is invalid
	if appointment.Date.Before(t.Add(time.Hour)) || appointment.Date.After(t.AddDate(1, 0, 0)) {
		return ErrInvalidAppointmentDate
	}

	// An appointment must not exist within ±3 hours of another appointment
	appointmentsInRange, err := s.repository.ReadInRange(ctx, appointment.Date.Add(time.Hour*-3), appointment.Date.Add(time.Hour*3))
	if err != nil {
		return err
	}
	if len(appointmentsInRange) > 0 {
		return ErrAppointmentTimeConflict
	}

	// TODO: prevent appointment conflicts
	// - An appointment cannot be appointed in the same day as a booking
	// - An appointment cannot be appointed out from working hours (9-18) (9am-6pm)UTC

	return s.repository.Create(ctx, appointment)
}

func (s *service) ReadByUserID(ctx context.Context, userID primitive.ObjectID) (*AppointmentEntity, error) {
	return s.repository.ReadByUserID(ctx, userID)
}

func (s *service) ReadAllByMonth(ctx context.Context, year int, month time.Month) ([]AppointmentEntity, error) {
	return s.repository.ReadAllByMonth(ctx, year, month)
}
