package user_model

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidUser = errors.New("invalid user")
)

type UserEntity struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// Client fields
	Name            string   `bson:"name"`
	PasswordHash    string   `bson:"passwordHash"`
	Email           Email    `bson:"email"`
	PhoneNumber     Phone    `bson:"phoneNumber"`
	EmergencyNumber Phone    `bson:"emergencyNumber"`
	Age             int      `bson:"age"`
	Pronouns        string   `bson:"pronouns"`
	Allergies       []string `bson:"allergies"`
	AdditionalInfo  string   `bson:"additionalInfo"`

	// Internal fields
	SessionToken string `json:"sessionToken" bson:"sessionToken"`
	CSRFToken    string `json:"csrfToken" bson:"csrfToken"`
}

func NewUserEntity(userParams UserParams) (*UserEntity, error) {
	return &UserEntity{
		ID:              primitive.NewObjectID(),
		Name:            userParams.Name,
		PasswordHash:    userParams.PasswordHash,
		Email:           userParams.Email,
		PhoneNumber:     userParams.PhoneNumber,
		EmergencyNumber: userParams.EmergencyNumber,
		Age:             userParams.Age,
		Pronouns:        userParams.Pronouns,
		AdditionalInfo:  userParams.AdditionalInfo,
		Allergies:       userParams.Allergies,
	}, nil
}

type UserParams struct {
	Name            string
	PasswordHash    string
	Email           Email
	PhoneNumber     Phone
	EmergencyNumber Phone
	Age             int
	Pronouns        string
	Allergies       []string
	AdditionalInfo  string
}
