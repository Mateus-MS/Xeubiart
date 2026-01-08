package appointment_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) ReadByUserID(ctx context.Context, userID primitive.ObjectID) (*AppointmentEntity, error) {
	appointment := AppointmentEntity{}

	err := r.Read(ctx, bson.M{"_id": userID}, appointment)
	if err != nil {
		return &appointment, err
	}

	return &appointment, nil
}
