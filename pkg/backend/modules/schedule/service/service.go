package schedule_service

import (
	"context"

	appointment_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/service"
	booking_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/booking/service"
	schedule_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/model"
	utils_models "github.com/Mateus-MS/Xeubiart.git/backend/utils/models"
)

type MonthScheduleDTO = schedule_model.MonthScheduleDTO

type IService interface {
	ReadByOffsetMonth(context.Context, int) (*MonthScheduleDTO, error)
}

type service struct {
	DepsServices struct {
		Appointment appointment_service.IService
		Booking     booking_service.IService
	}
	clock utils_models.Clock
}

// Constructor
func New(appointmentService appointment_service.IService, bookingService booking_service.IService, clock utils_models.Clock) *service {
	return &service{
		DepsServices: struct {
			Appointment appointment_service.IService
			Booking     booking_service.IService
		}{
			Appointment: appointmentService,
			Booking:     bookingService,
		},
		clock: clock,
	}
}
