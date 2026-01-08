package booking_repository

import (
	internal_repository "github.com/Mateus-MS/Xeubiart.git/backend/internal/repository"
	booking_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/booking/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingEntity = booking_model.BookEntity

type Repository struct {
	internal_repository.BaseRepository
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{
		BaseRepository: internal_repository.BaseRepository{Collection: coll},
	}
}
