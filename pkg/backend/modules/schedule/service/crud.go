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
func (s *service) ReadByOffsetMonth(ctx context.Context, offsetMonth int) (*MonthScheduleDTO, error) {
	dto := &schedule_model.MonthScheduleDTO{}

	// Ensures only positive values
	if offsetMonth < 0 {
		return dto, ErrNegativeMonthOffset
	}

	// Get the actual server time
	t := s.clock.Now()
	utcTime, err := internal_datetime.NewUTCTimeFromTime(t)
	if err != nil {
		return dto, err
	}

	// Add the offset to the query
	utcTime.AddDate(0, offsetMonth, 0)

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

	return schedule_model.NewMonthScheduleDTO(appointments, nil, utcTime.Year(), utcTime.Month(), utcTime.Day()), nil
}
