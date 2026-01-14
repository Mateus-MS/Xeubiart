package user_service

import (
	"context"

	user_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/user/model"
	user_repository "github.com/Mateus-MS/Xeubiart.git/backend/modules/user/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

// Type alias
type UserEntity = user_model.UserEntity

type IService interface {
	Register(context.Context, *UserEntity) error
}

type service struct {
	repository *user_repository.Repository
}

// Constructor
func New(ctx context.Context, coll *mongo.Collection) *service {
	return &service{
		repository: user_repository.New(ctx, coll),
	}
}
