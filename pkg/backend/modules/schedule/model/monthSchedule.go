package schedule_model

import (
	"encoding/json"
	"fmt"
	"time"

	appointment_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/model"
	booking_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/booking/model"
)

// Represents all month's callendar
type Schedule struct {
	Appointments []appointment_model.AppointmentEntity `json:"appointments"`
	Bookings     []booking_model.BookEntity            `json:"bookings"`
}

type date struct {
	Year         int          `json:"year"`
	Month        time.Month   `json:"month"`
	FirstWeekday time.Weekday `json:"firstWeekday"`
	DaysInMonth  int          `json:"daysInMonth"`
}

type MonthScheduleDTO struct {
	Date     date     `json:"date"`
	Schedule Schedule `json:"schedule"`
}

func (dto MonthScheduleDTO) String() string {
	// Convert to JSON for human-readable output
	data, err := json.MarshalIndent(dto, "", "  ")
	if err != nil {
		return fmt.Sprintf("MonthScheduleDTO{error marshalling: %v}", err)
	}
	return string(data)
}

func NewMonthScheduleDTO(appointments []appointment_model.AppointmentEntity, bookings []booking_model.BookEntity, year int, month time.Month) *MonthScheduleDTO {
	return &MonthScheduleDTO{
		Date: date{
			Year:         year,
			Month:        month,
			FirstWeekday: time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Weekday(),
			DaysInMonth:  time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, -1).Day(),
		},
		Schedule: Schedule{
			Appointments: appointments,
			Bookings:     bookings,
		},
	}
}
