package appointment_repository

import (
	appointment_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppointmentEntity = appointment_model.AppointmentEntity

type Repository struct {
	Collection *mongo.Collection
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{
		Collection: coll,
	}
}
