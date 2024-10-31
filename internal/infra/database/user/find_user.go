package user

import (
	"context"
	"errors"
	"github.com/zemartins81/posgo_leilao/internal/entity/user_entity"
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

func (r *UserRepository) FindUser(ctx context.Context, id string) (*user_entity.User, error) {
	filter := bson.M{"_id": id}
	
	var userEntityMOngo UserEntityMongo
	err := r.Collection.FindOne(ctx, filter).Decode(&userEntityMOngo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
		
		}
	}
	return
}
