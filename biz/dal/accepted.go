package dal

import (
	"LeetCode-Rank/db"
	"LeetCode-Rank/db/model"
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getAcceptedCollection(ctx context.Context) *mongo.Collection {
	return db.GetDatabase().Collection("accepted")
}

func UserID(userID int) Option {
	return func(filter bson.D) bson.D {
		filter = append(filter, bson.E{Key: "user_id", Value: bson.D{{Key: "$eq", Value: userID}}})
		return filter
	}
}

func CreatedAtGTE(createdAt time.Time) Option {
	return func(filter bson.D) bson.D {
		filter = append(filter, bson.E{Key: "created_at", Value: bson.D{{Key: "$gte", Value: createdAt}}})
		return filter
	}
}

func CreatedAtLT(createdAt time.Time) Option {
	return func(filter bson.D) bson.D {
		filter = append(filter, bson.E{Key: "created_at", Value: bson.D{{Key: "$lt", Value: createdAt}}})
		return filter
	}
}

func CreateAccepted(ctx context.Context, accepted *model.Accepted) (err error) {
	collection := getAcceptedCollection(ctx)
	_, err = collection.InsertOne(ctx, accepted)
	return err
}

func GetAcceptedEarlist(ctx context.Context, queryFuncs ...Option) (*model.Accepted, error) {
	filter := bson.D{}
	for _, f := range queryFuncs {
		filter = f(filter)
	}
	collection := getAcceptedCollection(ctx)
	cursor, err := collection.Find(ctx, filter, options.Find().SetSort(bson.D{{Key: "created_at", Value: 1}}).SetLimit(1))
	if err != nil {
		logrus.Errorf("GetAcceptedEarlist Find error: %v", err)
		return nil, err
	}
	var results []*model.Accepted
	if err := cursor.All(ctx, &results); err != nil {
		logrus.Errorf("GetAcceptedEarlist cursor.All error: %v", err)
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("GetAcceptedEarlist error: no accepted found")
	}
	return results[0], err
}

func GetAcceptedLatest(ctx context.Context, queryFuncs ...Option) (*model.Accepted, error) {
	filter := bson.D{}
	for _, f := range queryFuncs {
		filter = f(filter)
	}
	collection := getAcceptedCollection(ctx)
	cursor, err := collection.Find(ctx, filter, options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}).SetLimit(1))
	if err != nil {
		logrus.Errorf("GetAcceptedEarlist Find error: %v", err)
		return nil, err
	}
	var results []*model.Accepted
	if err := cursor.All(ctx, &results); err != nil {
		logrus.Errorf("GetAcceptedEarlist cursor.All error: %v", err)
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("GetAcceptedEarlist error: no accepted found")
	}
	return results[0], err
}
