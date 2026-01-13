package integration_schedule_test

import (
	"testing"
	"time"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
	integration_setup "github.com/Mateus-MS/Xeubiart.git/backend/tests/setup"
	integration_fixtures "github.com/Mateus-MS/Xeubiart.git/backend/tests/setup/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestScheduleRead_MultipleAppointments(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	location, _ := time.LoadLocation("America/New_York")

	baseDate, err := internal_datetime.NewLocalFromTime(time.Date(2026, time.January, 1, 12, 0, 0, 0, location))
	if err != nil {
		t.Fatalf("failed to build base date: %v", err)
	}

	a1, err := internal_datetime.NewLocalFromTime(time.Date(2026, time.January, 15, 12, 0, 0, 0, location))
	if err != nil {
		t.Fatalf("failed to build a1: %v", err)
	}
	a2, err := internal_datetime.NewLocalFromTime(time.Date(2026, time.January, 20, 12, 0, 0, 0, location))
	if err != nil {
		t.Fatalf("failed to build a2: %v", err)
	}

	integration_fixtures.InsertAppointment(t, h.Ctx, h.DB.Database, a1)
	integration_fixtures.InsertAppointment(t, h.Ctx, h.DB.Database, a2)

	result, err := h.Services.Schedule.ReadByOffsetMonth(h.Ctx, baseDate, 0)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Schedule.Days[15]))
	assert.Equal(t, 1, len(result.Schedule.Days[20]))
}

func TestScheduleRead_OffsetNextMonth(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	location, _ := time.LoadLocation("America/New_York")

	baseDate, err := internal_datetime.NewLocalFromTime(time.Date(2026, time.January, 1, 12, 0, 0, 0, location))
	if err != nil {
		t.Fatalf("failed to build base date: %v", err)
	}

	febDate, err := internal_datetime.NewLocalFromTime(time.Date(2026, time.February, 5, 12, 0, 0, 0, location))
	if err != nil {
		t.Fatalf("failed to build febDate: %v", err)
	}

	integration_fixtures.InsertAppointment(t, h.Ctx, h.DB.Database, febDate)

	result, err := h.Services.Schedule.ReadByOffsetMonth(h.Ctx, baseDate, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Schedule.Days[5]))
}

func TestScheduleRead_NoAppointments(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	location, _ := time.LoadLocation("America/New_York")
	date, err := internal_datetime.NewLocalFromTime(time.Date(2026, time.March, 10, 12, 0, 0, 0, location))
	if err != nil {
		t.Fatalf("failed to build date: %v", err)
	}

	result, err := h.Services.Schedule.ReadByOffsetMonth(h.Ctx, date, 0)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(result.Schedule.Days[10]))
}
