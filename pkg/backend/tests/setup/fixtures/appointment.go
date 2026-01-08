package integration_fixtures

import (
	"context"
	"testing"
	"time"

	appointment_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertAppointment(t *testing.T, ctx context.Context, db *mongo.Database, date time.Time, timezone string) {
	t.Helper()

	// Create the uuid
	uuid := primitive.NewObjectIDFromTimestamp(time.Now())

	// Create the appointment
	appointment, _ := appointment_model.NewEntity(uuid, date, timezone)

	// Store into DB
	_, err := db.Collection("appointment").InsertOne(ctx, appointment)
	if err != nil {
		panic(err)
	}
}
