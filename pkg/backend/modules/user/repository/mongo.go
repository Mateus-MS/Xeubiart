package user_repository

import (
	"context"

	internal_repository "github.com/Mateus-MS/Xeubiart.git/backend/internal/repository"
	user_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/user/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserEntity = user_model.UserEntity

type Repository struct {
	internal_repository.BaseRepository
}

func New(ctx context.Context, coll *mongo.Collection) *Repository {
	rep := Repository{
		BaseRepository: internal_repository.BaseRepository{Collection: coll},
	}

	if err := rep.EnsureIndexes(ctx); err != nil {
		panic("failed to ensure indexes: " + err.Error())
	}

	return &rep
}

func (r *Repository) EnsureIndexes(ctx context.Context) error {
	// Email is unique and required
	emailIndex := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	// Phone is unique but optional
	phoneIndex := mongo.IndexModel{
		Keys: bson.M{"phone": 1},
		Options: options.Index().
			SetUnique(true).
			SetPartialFilterExpression(bson.M{"phone": bson.M{"$exists": true}}),
	}

	_, err := r.Collection.Indexes().CreateMany(ctx, []mongo.IndexModel{emailIndex, phoneIndex})
	return err
}
