package appointment_service

import (
	"context"
	"errors"
	"time"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
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

	// An appointment must be within working hours
	// TODO: curretly an appointment can be booked to 17:50, think how to solve this.
	// maybe blocking anything to be booked after 15:00...
	if !appointment.Date.IsValidWorkingHours() {
		return internal_datetime.ErrOutsideWorkingHours
	}

	// TODO: prevent appointment conflicts
	// - An appointment cannot be appointed in the same day as a booking

	return s.repository.Create(ctx, appointment)
}

func (s *service) ReadByUserID(ctx context.Context, userID primitive.ObjectID) (*AppointmentEntity, error) {
	return s.repository.ReadByUserID(ctx, userID)
}

func (s *service) ReadAllByMonth(ctx context.Context, utcTime internal_datetime.UTCTime) ([]AppointmentEntity, error) {
	return s.repository.ReadAllByMonth(ctx, utcTime)
}
