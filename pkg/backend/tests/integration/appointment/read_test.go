package integration_appointment_test

import (
	"testing"
	"time"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
	integration_setup "github.com/Mateus-MS/Xeubiart.git/backend/tests/setup"
	integration_fixtures "github.com/Mateus-MS/Xeubiart.git/backend/tests/setup/fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppointmentRead_Success(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	// Add a new appointment into database
	location, _ := time.LoadLocation("America/New_York")
	localTime, err := internal_datetime.NewLocalFromTime(time.Now().In(location))
	integration_fixtures.InsertAppointment(t, h.Ctx, h.DB.Database, localTime)

	// Try to read it querying by month
	appointments, err := h.Services.Appointment.ReadAllByMonth(h.Ctx, localTime.ToUTCTime())

	// Assert
	require.NoError(t, err)
	assert.Equal(t, 1, len(appointments))
}
