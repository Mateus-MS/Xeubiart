package integration_appointment_test

import (
	"testing"
	"time"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
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

	userID := primitive.NewObjectIDFromTimestamp(h.Clock.Now())
	location, _ := time.LoadLocation("America/New_York")
	date, err := internal_datetime.NewLocalFromTime(h.Clock.Now().Add(time.Hour * 2).In(location))

	appointment, err := appointment_model.NewEntity(userID, date)
	require.NoError(t, err)

	err = h.Services.Appointment.Create(h.Ctx, appointment)
	assert.NoError(t, err)
}

func TestAppointmentCreate_TooCloseDate(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	userID := primitive.NewObjectIDFromTimestamp(h.Clock.Now())

	// Threat the date as from the received location
	location, _ := time.LoadLocation("America/New_York")
	date, err := internal_datetime.NewLocalFromTime(h.Clock.Now().In(location))

	// Try to make an appointment instantly
	appointment, err := appointment_model.NewEntity(userID, date)
	require.NoError(t, err)

	err = h.Services.Appointment.Create(h.Ctx, appointment)
	assert.ErrorIs(t, err, appointment_service.ErrInvalidAppointmentDate)
}

func TestAppointmentCreate_TooFarDate(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	userID := primitive.NewObjectIDFromTimestamp(h.Clock.Now())
	location, _ := time.LoadLocation("America/New_York")
	date, err := internal_datetime.NewLocalFromTime(h.Clock.Now().AddDate(2, 0, 0).In(location))

	// Try to make an appointment 2 year from now
	appointment, err := appointment_model.NewEntity(userID, date)
	require.NoError(t, err)

	err = h.Services.Appointment.Create(h.Ctx, appointment)
	assert.ErrorIs(t, err, appointment_service.ErrInvalidAppointmentDate)
}

func TestAppointmentCreate_TooCloseAppointments(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	userID := primitive.NewObjectIDFromTimestamp(h.Clock.Now())
	location, _ := time.LoadLocation("America/New_York")
	date, err := internal_datetime.NewLocalFromTime(h.Clock.Now().Add(time.Hour * 4).In(location))

	appointmentOBJ, err := appointment_model.NewEntity(userID, date)
	require.NoError(t, err)

	// Try to make the first appointment
	err = h.Services.Appointment.Create(h.Ctx, appointmentOBJ)
	require.NoError(t, err)

	// Try to make the second appointment in the same hour
	err = h.Services.Appointment.Create(h.Ctx, appointmentOBJ)
	require.ErrorIs(t, err, appointment_service.ErrAppointmentTimeConflict)
}
