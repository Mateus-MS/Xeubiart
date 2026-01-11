package integration_booking_test

import (
	"testing"
	"time"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
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
	location, _ := time.LoadLocation("America/New_York")
	date, err := internal_datetime.NewLocalFromTime(time.Now().Add(time.Hour * 2).In(location))

	// Try to make an booking instantly
	booking, err := booking_model.NewEntity(userID, date, booking_model.Tattoo)
	require.NoError(t, err)

	err = h.Services.Booking.Create(h.Ctx, booking)
	assert.NoError(t, err)
}

func TestBookingCreate_TooCloseDate(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	userID := primitive.NewObjectIDFromTimestamp(time.Now())

	location, _ := time.LoadLocation("America/New_York")
	date, err := internal_datetime.NewLocalFromTime(time.Now().In(location))

	// Try to make an booking instantly
	booking, err := booking_model.NewEntity(userID, date, booking_model.Tattoo)
	require.NoError(t, err)

	err = h.Services.Booking.Create(h.Ctx, booking)
	assert.ErrorIs(t, err, booking_service.ErrInvalidBookingDate)
}

func TestBookingCreate_TooFarDate(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	userID := primitive.NewObjectIDFromTimestamp(time.Now())
	location, _ := time.LoadLocation("America/New_York")
	date, err := internal_datetime.NewLocalFromTime(time.Now().AddDate(2, 0, 0).In(location))

	// Try to make an booking instantly
	booking, err := booking_model.NewEntity(userID, date, booking_model.Tattoo)
	require.NoError(t, err)

	err = h.Services.Booking.Create(h.Ctx, booking)
	assert.ErrorIs(t, err, booking_service.ErrInvalidBookingDate)
}
