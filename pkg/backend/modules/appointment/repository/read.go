package appointment_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *Repository) read(ctx context.Context, filter bson.M) (*AppointmentEntity, error) {
	appointment := AppointmentEntity{}

	err := repo.Collection.FindOne(ctx, filter).Decode(&appointment)
	if err != nil {
		return &AppointmentEntity{}, nil
	}

	return &appointment, nil
}

func (repo *Repository) ReadByUserID(ctx context.Context, userID primitive.ObjectID) (*AppointmentEntity, error) {
	return repo.read(ctx, bson.M{"_id": userID})
}
