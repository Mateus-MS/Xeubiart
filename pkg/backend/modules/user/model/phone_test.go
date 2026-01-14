package user_model_test

import (
	"testing"

	user_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/user/model"
	"github.com/stretchr/testify/require"
)

func TestPhoneCreate_Success(t *testing.T) {
	t.Parallel()
	phone, err := user_model.NewPhone("1234567890")
	require.NoError(t, err)
	require.Equal(t, "1234567890", phone.GetNumber())
}

func TestPhoneCreate_Invalid(t *testing.T) {
	t.Parallel()
	_, err := user_model.NewPhone("")
	require.ErrorIs(t, err, user_model.ErrInvalidPhone)
}
