package booking_service

import (
	"context"
	"time"

	booking_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/booking/model"
	booking_repository "github.com/Mateus-MS/Xeubiart.git/backend/modules/booking/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Type alias
type BookingEntity = booking_model.BookEntity

type IService interface {
	Create(context.Context, *BookingEntity) error
	ReadByUserID(context.Context, primitive.ObjectID) (*BookingEntity, error)
	ReadAllByMonth(ctx context.Context, year int, month time.Month) ([]BookingEntity, error)
}

type service struct {
	repository *booking_repository.Repository
}

// Constructor
func New(coll *mongo.Collection) *service {
	return &service{
		repository: booking_repository.New(coll),
	}
}
