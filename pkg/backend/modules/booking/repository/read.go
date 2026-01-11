package booking_repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) ReadByUserID(ctx context.Context, userID primitive.ObjectID) (*BookingEntity, error) {
	booking := BookingEntity{}

	err := r.ReadOne(ctx, bson.M{"_id": userID}, booking)
	if err != nil {
		return &booking, err
	}

	return &booking, nil
}

func (r *Repository) ReadAllByMonth(ctx context.Context, year int, month time.Month) ([]BookingEntity, error) {
	appointments := []BookingEntity{}

	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)

	filter := bson.M{
		"date": bson.M{
			"$gte": start,
			"$lt":  end,
		},
	}

	err := r.ReadAll(ctx, filter, &appointments)
	if err != nil {
		return appointments, err
	}

	return appointments, nil
}
