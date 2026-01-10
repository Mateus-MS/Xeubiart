package appointment_repository

import (
	"context"
	"time"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) ReadByUserID(ctx context.Context, userID primitive.ObjectID) (*AppointmentEntity, error) {
	appointment := AppointmentEntity{}

	err := r.ReadOne(ctx, bson.M{"_id": userID}, appointment)
	if err != nil {
		return &appointment, err
	}

	return &appointment, nil
}

func (r *Repository) ReadInRange(ctx context.Context, from, to time.Time) ([]AppointmentEntity, error) {
	appointments := []AppointmentEntity{}

	filter := bson.M{
		"date": bson.M{
			"$gte": from,
			"$lt":  to,
		},
	}

	err := r.ReadAll(ctx, filter, &appointments)
	if err != nil {
		return appointments, err
	}

	return appointments, nil
}

func (r *Repository) ReadAllByMonth(ctx context.Context, utcTime internal_datetime.UTCTime) ([]AppointmentEntity, error) {
	start := time.Date(utcTime.Year(), utcTime.Month(), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)

	return r.ReadInRange(ctx, start, end)
}
