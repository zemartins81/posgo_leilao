package bid

import (
	"context"
	"time"

	"github.com/zemartins81/posgo_leilao/configuration/logger"
	"github.com/zemartins81/posgo_leilao/internal/entity/bid_entity"
	"github.com/zemartins81/posgo_leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (bd *BidRepository) FindBidByAuctionId(ctx context.Context, auctionID string) ([]bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionID}

	cursor, err := bd.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("Error trying to find bids in database: ", err)
		return nil, internal_error.NewINternalServerError("Error trying to find bids in database")
	}

	var bidEntitiesMongo []BidEntityMongo
	if err := cursor.All(ctx, &bidEntitiesMongo); err != nil {
		logger.Error("Error trying to decode bids from database: ", err)
		return nil, internal_error.NewINternalServerError("Error trying to decode bids from database")
	}

	var bidEntities []bid_entity.Bid
	for _, b := range bidEntitiesMongo {
		bidEntities = append(bidEntities, bid_entity.Bid{
			Id:        b.Id,
			UserID:    b.UserID,
			AuctionID: b.AuctionID,
			Amount:    b.Amount,
			TimeStamp: time.Unix(b.TimeStamp, 0),
		})
	}

	return bidEntities, nil
}

func (bd *BidRepository) FindWinningBidAuctionId(ctx context.Context, auctionID string) (*bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionID}

	var bidEntityMongo BidEntityMongo

	opts := options.FindOne().SetSort(bson.D{{"amount", -1}})
	if err := bd.Collection.FindOne(ctx, filter, opts).Decode(&bidEntityMongo); err != nil {
		logger.Error("Error trying to find winning bid by auction ID: ", err)
		return nil, internal_error.NewINternalServerError("Error trying to find winning bid by auction ID")
	}

	return &bid_entity.Bid{
		Id:        bidEntityMongo.Id,
		UserID:    bidEntityMongo.UserID,
		AuctionID: bidEntityMongo.AuctionID,
		Amount:    bidEntityMongo.Amount,
		TimeStamp: time.Unix(bidEntityMongo.TimeStamp, 0),
	}, nil
}
