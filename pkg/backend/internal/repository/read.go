package internal_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *BaseRepository) ReadOne(ctx context.Context, filter bson.M, out any) error {
	return r.Collection.FindOne(ctx, filter).Decode(out)
}

func (r *BaseRepository) ReadAll(ctx context.Context, filter bson.M, out any) error {
	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, out)
	if err != nil {
		return err
	}

	return nil
}
