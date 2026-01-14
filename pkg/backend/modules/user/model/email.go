package user_model

import "errors"

var (
	ErrInvalidEmail = errors.New("invalid email")
)

type Email struct {
	address string
}

func (e Email) GetAddress() string {
	return e.address
}

func NewEmail(address string) (Email, error) {
	if address == "" {
		return Email{}, ErrInvalidEmail
	}

	return Email{address: address}, nil
}
