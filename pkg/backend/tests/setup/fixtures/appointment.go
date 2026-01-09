package integration_fixtures

import (
	"context"
	"testing"
	"time"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
	appointment_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertAppointment(t *testing.T, ctx context.Context, db *mongo.Database, date *internal_datetime.LocalTime) {
	t.Helper()

	// Create the uuid
	uuid := primitive.NewObjectIDFromTimestamp(time.Now())

	// Create the appointment
	appointment, _ := appointment_model.NewEntity(uuid, date)

	// Store into DB
	_, err := db.Collection("appointment").InsertOne(ctx, appointment)
	if err != nil {
		panic(err)
	}
}
