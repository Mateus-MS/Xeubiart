package integration_setup

import (
	"context"
	"testing"

	appointment_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/service"
	booking_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/booking/service"
	schedule_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/service"
)

type services struct {
	Appointment appointment_service.IService
	Booking     booking_service.IService
	Schedule    schedule_service.IService
}

type Harness struct {
	Ctx      context.Context
	DB       *TestDB
	Services *services
}

func NewHarness(t *testing.T) *Harness {
	t.Helper()

	testDB, err := NewTestDB(t.Name())
	if err != nil {
		t.Fatalf("failed to setup mongo: %v", err)
	}

	t.Cleanup(func() {
		_ = testDB.Teardown()
	})

	ctx := context.Background()

	appointment := appointment_service.New(testDB.Database.Collection("appointment"))
	booking := booking_service.New(testDB.Database.Collection("booking"))
	schedule := schedule_service.New(appointment, booking)

	services := services{
		Appointment: appointment,
		Booking:     booking,
		Schedule:    schedule,
	}

	return &Harness{
		Ctx:      ctx,
		DB:       testDB,
		Services: &services,
	}
}
