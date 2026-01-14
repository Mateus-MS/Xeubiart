package user_model

import (
	"errors"

	internal_security "github.com/Mateus-MS/Xeubiart.git/backend/internal/security"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidName = errors.New("invalid username")
)

type UserEntity struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	// Initially mandatory client fields
	Name         string                         `bson:"name"`
	PasswordHash internal_security.PasswordHash `bson:"passwordHash"`
	Email        Email                          `bson:"email"`

	// Initially optional client fields
	Phone *Phone `bson:"phone,omitempty"`
}

func (ue *UserEntity) IsRegistrationComplete() bool {
	return ue.Phone != nil
}

func NewUserEntity(name string, hashedPassword internal_security.PasswordHash, email Email, phone *Phone) (*UserEntity, error) {
	if name == "" {
		return nil, ErrInvalidName
	}

	userEntity := &UserEntity{
		Name:         name,
		PasswordHash: hashedPassword,
		Email:        email,
		Phone:        phone,
	}

	return userEntity, nil
}
