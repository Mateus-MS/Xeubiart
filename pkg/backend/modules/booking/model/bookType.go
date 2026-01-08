package booking_model

import "errors"

// BookType describes the type of service being scheduled.
type BookType int

var (
	ErrInvalidBookType = errors.New("the given BookType is invalid")
)

const (
	Tattoo BookType = iota
	Retouch
	Coverage
)

func (b BookType) IsValid() error {
	switch b {
	case Tattoo, Retouch, Coverage:
		return nil
	}
	return ErrInvalidBookType
}
