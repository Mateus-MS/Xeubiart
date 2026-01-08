package appointment_service

import (
	"errors"
)

var (
	ErrInvalidAppointmentDate = errors.New("the given date is invalid")
)
