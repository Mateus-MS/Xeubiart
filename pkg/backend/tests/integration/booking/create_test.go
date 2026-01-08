package integration_booking_test

import (
	"testing"
	"time"

	booking_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/booking/model"
	booking_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/booking/service"
	integration_setup "github.com/Mateus-MS/Xeubiart.git/backend/tests/setup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBookingCreate_Success(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	userID := primitive.NewObjectIDFromTimestamp(time.Now())
	date := time.Now().Add(time.Hour * 2)

	booking, err := booking_model.NewEntity(userID, date, "America/New_York", booking_model.Tattoo)
	require.NoError(t, err)

	err = h.Services.Booking.Create(h.Ctx, booking)
	assert.NoError(t, err)
}

func TestBookingCreate_TooCloseDate(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	userID := primitive.NewObjectIDFromTimestamp(time.Now())

	// Threat the date as from the received location
	loc, _ := time.LoadLocation("America/New_York")
	date := time.Now().In(loc)

	// Try to make an booking instantly
	booking, err := booking_model.NewEntity(userID, date, "America/New_York", booking_model.Tattoo)
	require.NoError(t, err)

	err = h.Services.Booking.Create(h.Ctx, booking)
	assert.ErrorIs(t, err, booking_service.ErrInvalidBookingDate)
}

func TestBookingCreate_TooFarDate(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	userID := primitive.NewObjectIDFromTimestamp(time.Now())
	date := time.Now().AddDate(2, 0, 0)

	// Try to make an booking 2 year from now
	booking, err := booking_model.NewEntity(userID, date, "America/New_York", booking_model.Tattoo)
	require.NoError(t, err)

	err = h.Services.Booking.Create(h.Ctx, booking)
	assert.ErrorIs(t, err, booking_service.ErrInvalidBookingDate)
}
