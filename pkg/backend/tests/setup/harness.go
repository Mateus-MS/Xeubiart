package integration_setup

import (
	"context"
	"testing"

	appointment_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/service"
)

type services struct {
	Appointment appointment_service.IService
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

	services := services{
		Appointment: appointment_service.New(testDB.Database.Collection("appointment")),
	}

	return &Harness{
		Ctx:      ctx,
		DB:       testDB,
		Services: &services,
	}
}
