package user_model

import "errors"

var (
	ErrInvalidPhone = errors.New("invalid phone")
)

type Phone struct {
	Number string `json:"number" bson:"number"`
}

func (p *Phone) IsValid() bool {
	// Simple test validation
	return p.Number != ""
}

func NewPhone(number string) (*Phone, error) {
	phone := &Phone{Number: number}
	if !phone.IsValid() {
		return nil, ErrInvalidPhone
	}
	return phone, nil
}
