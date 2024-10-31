package user

import (
	"context"
	"errors"

	"github.com/zemartins81/posgo_leilao/configuration/logger"
	"github.com/zemartins81/posgo_leilao/internal/entity/user_entity"
	"github.com/zemartins81/posgo_leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUSerRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (r *UserRepository) FindUser(ctx context.Context, id string) (*user_entity.User, *internal_error.InternalError) {
	filter := bson.M{"_id": id}

	var userEntityMOngo UserEntityMongo
	err := r.Collection.FindOne(ctx, filter).Decode(&userEntityMOngo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error("user not found", err)
			return nil, internal_error.NotFoundError("user not found")
		}
		logger.Error("Error trying to find user by userId", err)
		return nil, internal_error.NewINternalServerError("Error trying to find user by userId")
	}
	userEntity := &user_entity.User{
		ID:   userEntityMOngo.ID,
		Name: userEntityMOngo.Name,
	}
	return userEntity, nil
}
