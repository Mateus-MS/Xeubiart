package user_model_test

import (
	"testing"

	internal_security "github.com/Mateus-MS/Xeubiart.git/backend/internal/security"
	user_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/user/model"
	"github.com/stretchr/testify/require"
)

func TestUserNewEntity_SuccessEmptyPhone(t *testing.T) {
	t.Parallel()

	email, err := user_model.NewEmail("test@gmail.com")
	require.NoError(t, err)

	passHashed, err := internal_security.HashPassword("securePassword123")
	require.NoError(t, err)

	user, err := user_model.NewUserEntity("John Doe", passHashed, email, nil)
	require.NoError(t, err)

	require.Nil(t, user.Phone)
}

func TestUserNewEntity_SuccessValidPhone(t *testing.T) {
	t.Parallel()

	email, err := user_model.NewEmail("test@gmail.com")
	require.NoError(t, err)

	phone, err := user_model.NewPhone("1234567890")
	require.NoError(t, err)

	passHashed, err := internal_security.HashPassword("securePassword123")
	require.NoError(t, err)

	user, err := user_model.NewUserEntity("John Doe", passHashed, email, &phone)
	require.NoError(t, err)

	require.NotNil(t, user.Phone)
	require.True(t, user.IsRegistrationComplete())
}

func TestUserNewEntity_InvalidUsername(t *testing.T) {
	t.Parallel()

	email, err := user_model.NewEmail("test@gmail.com")
	require.NoError(t, err)

	passHashed, err := internal_security.HashPassword("securePassword123")
	require.NoError(t, err)

	_, err = user_model.NewUserEntity("", passHashed, email, nil)
	require.ErrorIs(t, err, user_model.ErrInvalidName)
}

func TestUserNewEntity_InvalidPassword(t *testing.T) {
	t.Parallel()

	_, err := internal_security.HashPassword("")
	require.ErrorIs(t, err, internal_security.ErrInvalidPassword)
}

func TestUserNewEntity_InvalidEmail(t *testing.T) {
	t.Parallel()

	_, err := user_model.NewEmail("")
	require.ErrorIs(t, err, user_model.ErrInvalidEmail)
}
