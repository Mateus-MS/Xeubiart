package user_model

import "errors"

var (
	ErrInvalidEmail = errors.New("invalid email")
)

type Email struct {
	Address string `json:"address" bson:"address"`
}

func (e *Email) IsValid() bool {
	// Simple test validation
	return e.Address != ""
}

func NewEmail(address string) (*Email, error) {
	email := &Email{Address: address}
	if !email.IsValid() {
		return nil, ErrInvalidEmail
	}
	return email, nil
}
