package utils_models

import "time"

type Clock interface {
	Now() time.Time
}

type AppClock struct {
}

func (AppClock) Now() time.Time {
	return time.Now().UTC()
}
