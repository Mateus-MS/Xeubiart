package user_service

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrEmailTaken = errors.New("email is already taken")
)

func (s *service) Register(ctx context.Context, user *UserEntity) error {
	err := s.repository.Create(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return ErrEmailTaken
		}
		return err
	}

	return nil
}
