package auction

import (
	"context"
	"fmt"
	"time"

	"github.com/zemartins81/posgo_leilao/configuration/logger"
	"github.com/zemartins81/posgo_leilao/internal/entity/auction_entity"
	"github.com/zemartins81/posgo_leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ar *AuctionRepository) FindAuctionByID(ctx context.Context, id string) (*auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{"_id": id}

	var auctionEntityMongo AuctionEntityMongo
	if err := ar.Collection.FindOne(ctx, filter).Decode(&auctionEntityMongo); err != nil {
		logger.Error(fmt.Sprintf("Error trying to find auction by ID: %s", id), err)
		return nil, internal_error.NewINternalServerError("Error trying to find auction by ID")
	}

	return &auction_entity.Auction{
		ID:          auctionEntityMongo.ID,
		ProductName: auctionEntityMongo.ProductName,
		Category:    auctionEntityMongo.Category,
		Description: auctionEntityMongo.Description,
		Conditon:    auctionEntityMongo.Conditon,
		Status:      auctionEntityMongo.Status,
		TimeStamp:   time.Unix(auctionEntityMongo.TimeStamp, 0),
	}, nil
}

func (ar *AuctionRepository) FindAuctions(ctx context.Context, status auction_entity.AuctionStatus, category, productName string) ([]auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{}

	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = category
	}

	if productName != "" {
		filter["product_name"] = primitive.Regex{Pattern: productName, Options: "i"}
	}

	cursor, err := ar.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("Error trying to find auctions in database: ", err)
		return nil, internal_error.NewINternalServerError("Error trying to find auctions in database")
	}

	defer cursor.Close(ctx)

	var auctionEntityMongo []AuctionEntityMongo
	if err := cursor.All(ctx, &auctionEntityMongo); err != nil {
		logger.Error("Error trying to decode auctions from database: ", err)
		return nil, internal_error.NewINternalServerError("Error trying to decode auctions from database")
	}

	var auctionEntity []auction_entity.Auction
	for _, a := range auctionEntityMongo {
		auctionEntity = append(auctionEntity, auction_entity.Auction{
			ID:          a.ID,
			ProductName: a.ProductName,
			Category:    a.Category,
			Description: a.Description,
			Conditon:    a.Conditon,
			Status:      a.Status,
			TimeStamp:   time.Unix(a.TimeStamp, 0),
		})
	}

	return auctionEntity, nil
}
