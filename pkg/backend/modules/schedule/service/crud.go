package schedule_service

import (
	"context"
	"time"

	schedule_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/model"
)

func (s *service) ReadAllByMonth(ctx context.Context, year int, month time.Month) (*MonthScheduleDTO, error) {
	dto := &schedule_model.MonthScheduleDTO{}

	appointments, err := s.DepsServices.Appointment.ReadAllByMonth(ctx, year, month)
	if err != nil {
		return dto, err
	}

	// TODO: do the same to bookings

	return schedule_model.NewMonthScheduleDTO(appointments, nil, year, month, "America/New_York"), nil
}
