package schedule_model

import (
	"time"

	appointment_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/model"
	booking_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/booking/model"
)

// Represents all month's callendar
type Schedule struct {
	Days map[int][]time.Time `json:"days"`
}

type date struct {
	Year         int          `json:"year"`
	Month        time.Month   `json:"month"`
	FirstWeekday time.Weekday `json:"firstWeekday"`
	DaysInMonth  int          `json:"daysInMonth"`
	Today        int          `json:"today"`
}

type MonthScheduleDTO struct {
	Date     date     `json:"date"`
	Schedule Schedule `json:"schedule"`
}

func NewMonthScheduleDTO(appointments []appointment_model.AppointmentEntity, bookings []booking_model.BookEntity, year int, month time.Month, today int) *MonthScheduleDTO {
	schedule := Schedule{}
	schedule.Days = make(map[int][]time.Time)

	for _, appointment := range appointments {
		day := appointment.Date.Day()
		schedule.Days[day] = append(schedule.Days[day], appointment.Date.Time)
	}

	// TODO: same for bookings

	return &MonthScheduleDTO{
		Date: date{
			Year:         year,
			Month:        month,
			Today:        today,
			FirstWeekday: time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Weekday(),
			DaysInMonth:  time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, -1).Day(),
		},
		Schedule: schedule,
	}
}
