package integration_appointment_test

import (
	"testing"
	"time"

	integration_setup "github.com/Mateus-MS/Xeubiart.git/backend/tests/setup"
	integration_fixtures "github.com/Mateus-MS/Xeubiart.git/backend/tests/setup/fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppointmentRead_Success(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	actualTime := time.Now()

	// Add a new appointment into database
	integration_fixtures.InsertAppointment(t, h.Ctx, h.DB.Database, actualTime.Add(time.Hour*5), "America/New_York")

	// Try to read it querying by month
	appointments, err := h.Services.Appointment.ReadAllByMonth(h.Ctx, actualTime.Year(), actualTime.Month())

	// Assert
	require.NoError(t, err)
	assert.Equal(t, 1, len(appointments))
}
