package bid

import (
	"context"
	"sync"

	"github.com/zemartins81/posgo_leilao/configuration/logger"
	"github.com/zemartins81/posgo_leilao/internal/entity/auction_entity"
	"github.com/zemartins81/posgo_leilao/internal/entity/bid_entity"
	"github.com/zemartins81/posgo_leilao/internal/infra/database/auction"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidEntityMongo struct {
	Id        string  `bson:"_id,omitempty"`
	UserID    string  `bson:"user_id"`
	AuctionID string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	TimeStamp int64   `bson:"time_stamp"`
}

type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *auction.AuctionRepository
}

func (bd *BidRepository) CreateBid(ctx context.Context, bidEntities []bid_entity.Bid) {
	var wg sync.WaitGroup

	for _, bid := range bidEntities {
		wg.Add(1)
		go func(bidValue bid_entity.Bid) {
			defer wg.Done()

			auctionEntity, err := bd.AuctionRepository.FindAuctionByID(ctx, bidValue.AuctionID)
			if err != nil {
				logger.Error("Error trying to find auction by ID: ", err)
				return
			}

			if auctionEntity.Status != auction_entity.Active {
				return
			}

			bidEntityMongo := BidEntityMongo{
				Id:        bidValue.Id,
				UserID:    bidValue.UserID,
				AuctionID: bidValue.AuctionID,
				Amount:    bidValue.Amount,
				TimeStamp: bidValue.TimeStamp.Unix(),
			}

			if _, err := bd.Collection.InsertOne(ctx, bidEntityMongo); err != nil {
				logger.Error("Error trying to insert bid in database: ", err)
				return
			}

		}(bid)
	}

	wg.Wait()
}
