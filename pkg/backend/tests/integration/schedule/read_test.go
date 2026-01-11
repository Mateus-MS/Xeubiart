package integration_schedule_test

import (
	"testing"
	"time"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
	integration_setup "github.com/Mateus-MS/Xeubiart.git/backend/tests/setup"
	integration_fixtures "github.com/Mateus-MS/Xeubiart.git/backend/tests/setup/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestScheduleRead_Success(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	location, _ := time.LoadLocation("America/New_York")
	date, err := internal_datetime.NewLocalFromTime(time.Date(2026, time.January, 15, 12, 0, 0, 0, location))

	integration_fixtures.InsertAppointment(t, h.Ctx, h.DB.Database, date)

	result, err := h.Services.Schedule.ReadByMonth(h.Ctx, date.ToUTCTime())
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Schedule.Appointments))
}
