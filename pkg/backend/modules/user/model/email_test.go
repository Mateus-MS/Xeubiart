package user_model_test

import (
	"testing"

	user_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/user/model"
	"github.com/stretchr/testify/require"
)

func TestEmailCreate_Success(t *testing.T) {
	t.Parallel()
	email, err := user_model.NewEmail("test@gmail.com")
	require.NoError(t, err)
	require.Equal(t, "test@gmail.com", email.GetAddress())
}

func TestEmailCreate_Invalid(t *testing.T) {
	t.Parallel()
	_, err := user_model.NewEmail("")
	require.ErrorIs(t, err, user_model.ErrInvalidEmail)
}
