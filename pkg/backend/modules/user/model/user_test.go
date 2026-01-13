package user_model_test

import (
	"testing"

	user_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/user/model"
	"github.com/stretchr/testify/require"
)

func HelperGetValidUserParams(t *testing.T) user_model.UserParams {
	t.Helper()

	name := "John Doe"
	passwordHash := "hashed_password"
	email, err := user_model.NewEmail("test@outlook.com")
	require.NoError(t, err)
	phone, err := user_model.NewPhone("+1234567890")
	require.NoError(t, err)
	emergencyNumber, err := user_model.NewPhone("+0987654321")
	require.NoError(t, err)
	age := 30
	pronouns := "he/him"
	allergies := []string{"peanuts", "shellfish"}
	additionalInfo := "Some additional info"

	return user_model.UserParams{
		Name:            name,
		PasswordHash:    passwordHash,
		Email:           *email,
		PhoneNumber:     *phone,
		EmergencyNumber: *emergencyNumber,
		Age:             age,
		Pronouns:        pronouns,
		Allergies:       allergies,
		AdditionalInfo:  additionalInfo,
	}
}

func TestUserNewEntity_Success(t *testing.T) {
	t.Parallel()

	params := HelperGetValidUserParams(t)

	_, err := user_model.NewUserEntity(params)
	require.NoError(t, err)
}
