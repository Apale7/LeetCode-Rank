package db

import (
	config "LeetCode-Rank/config_loader"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init(ctx context.Context) {
	InitMongo(ctx)
}

func InitMongo(ctx context.Context) {
	var err error
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Get("mongo_uri")).SetMaxPoolSize(20))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
}

func GetDatabase() *mongo.Database {
	return client.Database("leetcode")
}
