package schedule_service

import (
	"context"
	"time"

	schedule_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/model"
)

func (s *service) ReadAllByMonth(ctx context.Context, year int, month time.Month, timezone string) (*MonthScheduleDTO, error) {
	dto := &schedule_model.MonthScheduleDTO{}

	appointments, err := s.DepsServices.Appointment.ReadAllByMonth(ctx, year, month)
	if err != nil {
		return dto, err
	}

	booking, err := s.DepsServices.Booking.ReadAllByMonth(ctx, year, month)
	if err != nil {
		return dto, err
	}

	return schedule_model.NewMonthScheduleDTO(appointments, booking, year, month, timezone), nil
}
