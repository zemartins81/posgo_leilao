package mongodb

import (
	"context"
	"github.com/zemartins81/posgo_leilao/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const (
	MONGODB_URL = "MONGODB_URL"
	MONGODB_DB  = "MONGODB_DB"
)

func NewMongoDBConnections(ctx context.Context) (*mongo.Database, error) {
	mongoUrl := os.Getenv(MONGODB_URL)
	mongoDatabase := os.Getenv(MONGODB_DB)

	client, err := mongo.Connect(
		ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		logger.Error("Error trying to connect to MongoDB: ", err)
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Error trying to ping MongoDB: ", err)
		return nil, err
	}

	return client.Database(mongoDatabase), nil
}
