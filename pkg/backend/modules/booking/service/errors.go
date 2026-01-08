package booking_service

import (
	"errors"
)

var (
	ErrInvalidBookingDate = errors.New("the given date is invalid")
)
