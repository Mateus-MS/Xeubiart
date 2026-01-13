package schedule_service

import (
	"context"
	"errors"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
	schedule_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/model"
)

var (
	ErrNegativeMonthOffset = errors.New("the given month offset is negative")
)

// From the actual time, offset the month by the received.
func (s *service) ReadByOffsetMonth(ctx context.Context, localTime *internal_datetime.LocalTime, offsetMonth int) (*MonthScheduleDTO, error) {
	dto := &schedule_model.MonthScheduleDTO{}

	// Converts localtime to UTC and add the offset to the query
	utcTime, err := internal_datetime.NewUTCTimeFromTime(localTime.Time.UTC().AddDate(0, offsetMonth, 0))
	if err != nil {
		return dto, err
	}

	// query
	appointments, err := s.DepsServices.Appointment.ReadAllByMonth(ctx, *utcTime)
	if err != nil {
		return dto, err
	}

	// booking, err := s.DepsServices.Booking.ReadAllByMonth(ctx, year, month)
	// if err != nil {
	// 	return dto, err
	// }

	// TODO: The received month cannot be previous from the actual month, neighter year

	// Only send today if the month offset is 0
	response := schedule_model.NewMonthScheduleDTO(appointments, nil, utcTime.Year(), utcTime.Month(), utcTime.Day())
	if offsetMonth != 0 {
		response.Date.Today = 0
	}

	return response, nil
}
