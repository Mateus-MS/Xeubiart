package internal_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// This base repository carryies the most basic CRUD mongo db functions
type BaseRepository struct {
	Collection *mongo.Collection
}

func (r *BaseRepository) Create(ctx context.Context, doc any) error {
	_, err := r.Collection.InsertOne(ctx, doc)
	return err
}

func (r *BaseRepository) Update(ctx context.Context, filter bson.M, update bson.M) error {
	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (r *BaseRepository) Delete(ctx context.Context, filter bson.M) error {
	result, err := r.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
