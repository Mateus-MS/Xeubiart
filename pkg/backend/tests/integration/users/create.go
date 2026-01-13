package integration_user_test

import (
	"testing"

	integration_setup "github.com/Mateus-MS/Xeubiart.git/backend/tests/setup"
)

func TestScheduleRead_MultipleAppointments(t *testing.T) {
	t.Parallel()
	_ = integration_setup.NewHarness(t)
}
