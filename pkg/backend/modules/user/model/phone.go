package user_model

import "errors"

var (
	ErrInvalidPhone = errors.New("invalid phone")
)

type Phone struct {
	number string
}

func (p Phone) GetNumber() string {
	return p.number
}

func NewPhone(number string) (Phone, error) {
	if number == "" {
		return Phone{}, ErrInvalidPhone
	}
	return Phone{number: number}, nil
}
