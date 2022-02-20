package dal

import (
	"LeetCode-Rank/db"
	"LeetCode-Rank/db/model"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getUserCollection(ctx context.Context) *mongo.Collection {
	return db.GetDatabase().Collection("user")
}

type Option func(bson.D) bson.D

func Username(username string) Option {
	return func(filter bson.D) bson.D {
		filter = append(filter, bson.E{Key: "username", Value: bson.D{{"$eq", username}}})
		return filter
	}
}

func GetUser(ctx context.Context, queryFuncs ...Option) (user *model.User, err error) {
	user = &model.User{}
	filter := bson.D{}
	for _, f := range queryFuncs {
		filter = f(filter)
	}
	collection := getUserCollection(ctx)
	err = collection.FindOne(ctx, filter).Decode(user)
	return user, err
}

func GetUserLatest(ctx context.Context, queryFuncs ...Option) (*model.User, error) {
	filter := bson.D{}
	for _, f := range queryFuncs {
		filter = f(filter)
	}
	collection := getUserCollection(ctx)
	cursor, err := collection.Find(ctx, filter, options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}).SetLimit(1))
	if err != nil {
		logrus.Errorf("GetUserLatest error: %v", err)
		return nil, err
	}
	var results []*model.User
	if err = cursor.All(context.TODO(), &results); err != nil {
		logrus.Errorf("GetUserLatest error: %v", err)
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("GetUserLatest error: no user found")
	}

	return results[0], err
}

func GetUsers(ctx context.Context, queryFuncs ...Option) (users []*model.User, err error) {
	users = []*model.User{}
	filter := bson.D{}
	for _, f := range queryFuncs {
		filter = f(filter)
	}
	collection := getUserCollection(ctx)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return users, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		user := &model.User{}
		err = cur.Decode(user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func CreateUser(ctx context.Context, user *model.User) (err error) {
	collection := getUserCollection(ctx)
	_, err = collection.InsertOne(ctx, user)
	return err
}
