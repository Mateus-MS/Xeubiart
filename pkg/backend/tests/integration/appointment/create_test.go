package integration_appointment_test

import (
	"testing"
	"time"

	appointment_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/model"
	appointment_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/service"
	integration_setup "github.com/Mateus-MS/Xeubiart.git/backend/tests/setup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAppointmentCreate_Success(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	userID := primitive.NewObjectIDFromTimestamp(time.Now())
	date := time.Now().Add(time.Hour * 2)

	appointment, err := appointment_model.NewEntity(userID, date, "America/New_York")
	require.NoError(t, err)

	err = h.Services.Appointment.Create(h.Ctx, appointment)
	assert.NoError(t, err)
}

func TestAppointmentCreate_TooCloseDate(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	userID := primitive.NewObjectIDFromTimestamp(time.Now())

	// Threat the date as from the received location
	loc, _ := time.LoadLocation("America/New_York")
	date := time.Now().In(loc)

	// Try to make an appointment instantly
	appointment, err := appointment_model.NewEntity(userID, date, "America/New_York")
	require.NoError(t, err)

	err = h.Services.Appointment.Create(h.Ctx, appointment)
	assert.ErrorIs(t, err, appointment_service.ErrInvalidAppointmentDate)
}

func TestAppointmentCreate_TooFarDate(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	userID := primitive.NewObjectIDFromTimestamp(time.Now())
	date := time.Now().AddDate(2, 0, 0)

	// Try to make an appointment 2 year from now
	appointment, err := appointment_model.NewEntity(userID, date, "America/New_York")
	require.NoError(t, err)

	err = h.Services.Appointment.Create(h.Ctx, appointment)
	assert.ErrorIs(t, err, appointment_service.ErrInvalidAppointmentDate)
}
