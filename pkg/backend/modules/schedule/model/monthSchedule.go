package schedule_model

import (
	"time"

	appointment_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/model"
	booking_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/booking/model"
)

// Represents all month's callendar
type Schedule struct {
	Appointments []appointment_model.AppointmentEntity `json:"appointments"`
	Bookings     []booking_model.BookEntity            `json:"bookings"`
}

type MonthScheduleDTO struct {
	Date struct {
		Year         int          `json:"year"`
		Month        time.Month   `json:"month"`
		FirstWeekday time.Weekday `json:"firstWeekday"`
		DaysInMonth  int          `json:"daysInMonth"`
	} `json:"date"`

	Schedule Schedule `json:"schedule"`
}

func NewMonthScheduleDTO(appointments []appointment_model.AppointmentEntity, bookings []booking_model.BookEntity, year int, month time.Month, timezone string) *MonthScheduleDTO {
	return &MonthScheduleDTO{
		Date: struct {
			Year         int          `json:"year"`
			Month        time.Month   `json:"month"`
			FirstWeekday time.Weekday `json:"firstWeekday"`
			DaysInMonth  int          `json:"daysInMonth"`
		}{
			Year:         year,
			Month:        month,
			FirstWeekday: 0,  // Temp
			DaysInMonth:  20, // Temp
		},
		Schedule: Schedule{
			Appointments: appointments,
			Bookings:     bookings,
		},
	}
}
