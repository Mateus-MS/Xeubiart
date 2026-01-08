package appointment_service

import (
	"errors"
)

var (
	ErrInvalidDate = errors.New("the given date is invalid")
)
