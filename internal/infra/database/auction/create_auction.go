package auction

import (
	"context"

	"github.com/zemartins81/posgo_leilao/configuration/logger"
	"github.com/zemartins81/posgo_leilao/internal/internal_error"

	"github.com/zemartins81/posgo_leilao/internal/entity/auction_entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	ID          string                          `bson:"_id,omitempty"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Conditon    auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	TimeStamp   int64                           `bson:"time_stamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(ctx context.Context, auctionEntity auction_entity.Auction) *internal_error.InternalError {
	auctionEntityMongo := AuctionEntityMongo{
		ID:          auctionEntity.ID,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Conditon:    auctionEntity.Conditon,
		Status:      auctionEntity.Status,
		TimeStamp:   auctionEntity.TimeStamp.Unix(),
	}

	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert auction in database: ", err)
		return internal_error.NewINternalServerError("Error trying to insert auction in database")
	}
	return nil
}
