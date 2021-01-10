package db

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var (
	RedisClient *redis.Client
)

func init() {
	if err := initClient(); err != nil {
		panic(err)
	}
	log.Info("连接Redis成功")
}

func initClient() (err error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = RedisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
