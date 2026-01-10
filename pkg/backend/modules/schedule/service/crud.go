package schedule_service

import (
	"context"
	"time"

	appointment_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/model"
	schedule_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/model"
)

func (s *service) ReadByMonth(ctx context.Context, year int, month time.Month) (*MonthScheduleDTO, error) {
	dto := &schedule_model.MonthScheduleDTO{}

	appointments, err := s.DepsServices.Appointment.ReadAllByMonth(ctx, year, month)
	if err != nil {
		return dto, err
	}

	// Convert the appointment to safely be consumed
	appointmentsDTO := make([]appointment_model.AppointmentDTO, len(appointments))
	for i, appointment := range appointments {
		appointmentsDTO[i] = *appointment.ToDTO()
	}

	// booking, err := s.DepsServices.Booking.ReadAllByMonth(ctx, year, month)
	// if err != nil {
	// 	return dto, err
	// }

	// TODO: The received month cannot be previous from the actual month, neighter year

	return schedule_model.NewMonthScheduleDTO(appointmentsDTO, nil, year, month), nil
}
