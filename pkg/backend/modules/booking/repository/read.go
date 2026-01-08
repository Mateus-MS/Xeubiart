package booking_repository

import (
	"context"

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
