package bid_entity

import "time"

type Bid struct {
	Id		string `bson:"_id,omitempty"`
	UserID	string `bson:"user_id"`
	AuctionID	string `bson:"auction_id"`
	Amount	float64 `bson:"amount"`
	TimeStamp	time.Time `bson:"time_stamp"`
}