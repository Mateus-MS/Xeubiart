package schedule_service

import (
	"context"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
	appointment_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/model"
	schedule_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/model"
)

func (s *service) ReadByMonth(ctx context.Context, utcTime internal_datetime.UTCTime) (*MonthScheduleDTO, error) {
	dto := &schedule_model.MonthScheduleDTO{}

	appointments, err := s.DepsServices.Appointment.ReadAllByMonth(ctx, utcTime)
	if err != nil {
		return dto, err
	}

	appointmentsDTO := make([]appointment_model.AppointmentDTO, len(appointments))
	for i, appointment := range appointments {
		// Convert to DTO
		appointmentsDTO[i] = *appointment.ToDTO()
	}

	// booking, err := s.DepsServices.Booking.ReadAllByMonth(ctx, year, month)
	// if err != nil {
	// 	return dto, err
	// }

	// TODO: The received month cannot be previous from the actual month, neighter year

	return schedule_model.NewMonthScheduleDTO(appointmentsDTO, nil, utcTime.Year(), utcTime.Month()), nil
}
