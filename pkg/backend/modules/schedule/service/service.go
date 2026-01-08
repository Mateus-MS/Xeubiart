package schedule_service

import (
	"context"
	"time"

	appointment_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/service"
	booking_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/booking/service"
	schedule_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/model"
)

type MonthScheduleDTO = schedule_model.MonthScheduleDTO

type IService interface {
	ReadAllByMonth(context.Context, int, time.Month) (*MonthScheduleDTO, error)
}

type service struct {
	DepsServices struct {
		Appointment appointment_service.IService
		Booking     booking_service.IService
	}
}

// Constructor
func New(appointmentService appointment_service.IService, bookingService booking_service.IService) *service {
	return &service{
		DepsServices: struct {
			Appointment appointment_service.IService
			Booking     booking_service.IService
		}{
			Appointment: appointmentService,
			Booking:     bookingService,
		},
	}
}
