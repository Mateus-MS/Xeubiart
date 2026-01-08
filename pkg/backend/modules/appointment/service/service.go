package appointment_service

import (
	"context"
	"time"

	appointment_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/model"
	appointment_repository "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Type alias
type AppointmentEntity = appointment_model.AppointmentEntity

type IService interface {
	Create(context.Context, *AppointmentEntity) error
	ReadByUserID(context.Context, primitive.ObjectID) (*AppointmentEntity, error)
	ReadAllByMonth(ctx context.Context, year int, month time.Month) ([]AppointmentEntity, error)
}

type service struct {
	repository *appointment_repository.Repository
}

// Constructor
func New(coll *mongo.Collection) *service {
	return &service{
		repository: appointment_repository.New(coll),
	}
}
