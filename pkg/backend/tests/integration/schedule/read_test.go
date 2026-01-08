package integration_schedule_test

import (
	"testing"
	"time"

	integration_setup "github.com/Mateus-MS/Xeubiart.git/backend/tests/setup"
	integration_fixtures "github.com/Mateus-MS/Xeubiart.git/backend/tests/setup/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestScheduleRead_Success(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	actualTime := time.Now()
	integration_fixtures.InsertAppointment(t, h.Ctx, h.DB.Database, actualTime.Add(time.Hour*5), "America/New_York")
	integration_fixtures.InsertAppointment(t, h.Ctx, h.DB.Database, actualTime.Add(time.Hour*5), "America/New_York")
	integration_fixtures.InsertAppointment(t, h.Ctx, h.DB.Database, actualTime.Add(time.Hour*5), "America/New_York")
	integration_fixtures.InsertAppointment(t, h.Ctx, h.DB.Database, actualTime.Add(time.Hour*5), "America/New_York")

	dto, err := h.Services.Schedule.ReadAllByMonth(h.Ctx, actualTime.Year(), actualTime.Month(), "America/New_York")
	assert.NoError(t, err)
	assert.Equal(t, 4, len(dto.Schedule.Appointments))
}
