package integration_setup

import (
	"context"
	"testing"
	"time"

	appointment_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/service"
	booking_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/booking/service"
	schedule_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/service"
	utils_models "github.com/Mateus-MS/Xeubiart.git/backend/utils/models"
)

type MockClock struct {
}

func (MockClock) Now() time.Time {
	return time.Date(2026, 01, 11, 15, 00, 00, 00, time.UTC)
}

type services struct {
	Appointment appointment_service.IService
	Booking     booking_service.IService
	Schedule    schedule_service.IService
}

type Harness struct {
	Ctx      context.Context
	DB       *TestDB
	Services *services
	Clock    utils_models.Clock
}

func NewHarness(t *testing.T) *Harness {
	t.Helper()

	testDB, err := NewTestDB()
	if err != nil {
		t.Fatalf("failed to setup mongo: %v", err)
	}

	mockClock := MockClock{}

	t.Cleanup(func() {
		_ = testDB.Teardown()
	})

	ctx := context.Background()

	appointment := appointment_service.New(testDB.Database.Collection("appointment"), mockClock)
	booking := booking_service.New(testDB.Database.Collection("booking"))
	schedule := schedule_service.New(appointment, booking, mockClock)

	services := services{
		Appointment: appointment,
		Booking:     booking,
		Schedule:    schedule,
	}

	return &Harness{
		Ctx:      ctx,
		DB:       testDB,
		Services: &services,
		Clock:    mockClock,
	}
}
